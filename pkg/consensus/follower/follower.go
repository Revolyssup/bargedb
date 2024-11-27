package follower

import "github.com/revolyssup/bargedb/pkg/consensus/signal"

type Follower struct {
	Signal *signal.Signaller
}
type FollowerParams struct {
	Signal *signal.Signaller
}

func Init(params FollowerParams) *Follower {
	return &Follower{
		Signal: params.Signal,
	}
}

func (f *Follower) Run() {}
