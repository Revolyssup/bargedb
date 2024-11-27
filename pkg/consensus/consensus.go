package consensus

import (
	"github.com/revolyssup/bargedb/pkg/consensus/candidate"
	"github.com/revolyssup/bargedb/pkg/consensus/follower"
	"github.com/revolyssup/bargedb/pkg/consensus/leader"
	"github.com/revolyssup/bargedb/pkg/consensus/log"
	"github.com/revolyssup/bargedb/pkg/consensus/signal"
)

type State interface {
	Run(c *Consensus)
}
type Consensus struct {
	signal *signal.Signaller
	log    *log.Log
}

func (c *Consensus) Run() {
	for {
		select {
		case <-c.signal.RunCandidate:
			candidate.Init(candidate.CandidateParams{
				Signal: c.signal,
			}).Run()
		case <-c.signal.RunFollower:
			follower.Init(follower.FollowerParams{
				Signal: c.signal,
			}).Run()
		case <-c.signal.RunLeader:
			leader.Init(leader.LeaderParams{
				Signal: c.signal,
			}).Run()
		}
	}
}
