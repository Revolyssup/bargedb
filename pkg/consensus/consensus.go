package consensus

import (
	"github.com/revolyssup/bargedb/pkg/consensus/candidate"
	"github.com/revolyssup/bargedb/pkg/consensus/follower"
	"github.com/revolyssup/bargedb/pkg/consensus/leader"
	"github.com/revolyssup/bargedb/pkg/consensus/signal"
	"github.com/revolyssup/bargedb/pkg/log"
	"github.com/revolyssup/bargedb/pkg/transport"
)

type State interface {
	Run(c *Consensus)
}
type Consensus struct {
	signal    *signal.Signaller
	log       *log.Log
	transport transport.Transport
}

func (c *Consensus) Run() {
	for {
		select {
		case <-c.signal.RunCandidate:
			candidate.Init(candidate.CandidateParams{
				Signal:    c.signal,
				Log:       c.log,
				Transport: c.transport,
			}).Run()
		case <-c.signal.RunFollower:
			follower.Init(follower.FollowerParams{
				Signal:    c.signal,
				Log:       c.log,
				Transport: c.transport,
			}).Run()
		case <-c.signal.RunLeader:
			leader.Init(leader.LeaderParams{
				Signal:    c.signal,
				Log:       c.log,
				Transport: c.transport,
			}).Run()
		}
	}
}
