package consensus

import (
	"context"
	"sync"
	"time"

	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/google/uuid"
)

type State interface {
	ApplyAction(consensus *Instance, act Action) bool //Applies the Action and returns true if it changes the state
	//Will be run by Main Start routine after cancelling the previous run.
	//Start on all states should immediately exit once it recieves signal on context.Done() channel
	Start(ctx context.Context, consensus *Instance)
}

var stateLock sync.Mutex = sync.Mutex{} //Any operation that potentially changes the states should first acquire a lock

// The states should not have redundant data that is stored at the consensus layer.
// The data here will only be relevant during the course of state change.
type LeaderState struct {
}
type CandidateState struct {
	CandidateID uuid.UUID
	term        int
	timeout     time.Duration
}

func (l *CandidateState) ApplyAction(consensus *Instance, act Action) bool {
	switch act.Name {
	case TIMEOUT:
		consensus.Start(&CandidateState{
			CandidateID: l.CandidateID,
			term:        l.term,
			timeout:     l.timeout,
		})
	case BecomeFollowerImmediately:
		go consensus.Start(&FollowerState{
			timeout: l.timeout,
		})
		return true
	}
	return false
}
func (l *CandidateState) Start(ctx context.Context, consensus *Instance) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			l.term++
			consensus.vote(l.CandidateID)
			go func() {
				time.Sleep(l.timeout) //TODO: Randomize the timeout
				_ = l.ApplyAction(consensus, Action{
					Name: TIMEOUT,
				})
			}()
			index, term := consensus.lastLogIndeAndTerm()
			totalServers := len(consensus.nextIndex) + 1
			voted := 0
			go func() {
				for resp := range consensus.Transport.RequestVote(l.term, l.CandidateID, log.Index(index), term) {
					go func(resp map[uuid.UUID]map[string]interface{}) {
						for _, r := range resp {
							if r["voted"].(bool) {
								voted++
							}
						}
					}(resp)
					if voted > totalServers/2 {
						consensus.Start(&LeaderState{})
						return
					}
				}

			}()

		}
	}
}
func (l *LeaderState) ApplyAction(consensus *Instance, act Action) bool {
	switch act.Name {
	case TIMEOUT:

	}
	return false
}
func (l *LeaderState) Start(ctx context.Context, consensus *Instance) {

}

type FollowerState struct {
	timeout time.Duration
}

func (l *FollowerState) ApplyAction(consensus *Instance, act Action) bool {
	switch act.Name {
	case TIMEOUT:
		consensus.Start(&CandidateState{
			timeout: l.timeout,
		})
		return true
	case BecomeFollowerImmediately:
		go consensus.Start(&FollowerState{
			timeout: l.timeout,
		})
		return true
	}
	return false
}

// When ApplyAction inside of Start returns true for a state change, then Start should call go Start() on the returned state.
func (l *FollowerState) Start(ctx context.Context, consensus *Instance) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			go func() {
				time.Sleep(l.timeout) //TODO: Randomize the timeout
				_ = l.ApplyAction(consensus, Action{
					Name: TIMEOUT,
				})
			}()
		}
	}
}
