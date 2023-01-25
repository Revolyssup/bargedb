package consensus

import (
	"sync"
	"time"
)

type State interface {
	ApplyAction(consensus *Instance, act Action) bool //Applies the Action and returns true if it changes the state
	//Will be run by Main Start routine after cancelling the previous run.
	//Start should exist once it recieves signal on stop channel
	Start(stop <-chan interface{}) //All state changes are manageed by ApplyAction and Start will be called ApplyAction brings the state from where it can START
}

var stateLock sync.Mutex = sync.Mutex{} //Any operation that potentially changes the states should first acquire a lock

// The states should not have redundant data that is stored at the consensus layer.
// The data here will only be relevant during the course of state change.
type LeaderState struct {
}
type CandidateState struct {
}

func (l *CandidateState) ApplyAction(consensus *Instance, act Action) bool {
	switch act.Name {
	case TIMEOUT:

	}
	return false
}
func (l *CandidateState) Start(stop <-chan interface{}) {
	for {
		select {
		case <-stop:
			return
		default:

		}
	}
}
func (l *LeaderState) ApplyAction(consensus *Instance, act Action) bool {
	switch act.Name {
	case TIMEOUT:

	}
	return false
}

type FollowerState struct {
	timeout time.Duration
	restart chan interface{}
}

func (l *FollowerState) ApplyAction(consensus *Instance, act Action) bool {
	switch act.Name {
	case TIMEOUT:
		consensus.State = &CandidateState{}
		return true
	case RESTART:
		l.restart <- 0
	}
	return false
}

// When ApplyAction inside of Start returns true for a state change, then Start should call go Start() on the returned state.
func (l *FollowerState) Start(stop <-chan interface{}) {
	for {
		select {
		case <-stop:
			return
		default:

		}
	}
}
