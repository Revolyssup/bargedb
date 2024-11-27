package follower

import (
	"github.com/revolyssup/bargedb/pkg/consensus/signal"
	"github.com/revolyssup/bargedb/pkg/log"
	"github.com/revolyssup/bargedb/pkg/transport"
)

type Follower struct {
	Signal    *signal.Signaller
	Log       *log.Log
	Transport transport.Transport
}
type FollowerParams struct {
	Signal    *signal.Signaller
	Log       *log.Log
	Transport transport.Transport
}

func Init(params FollowerParams) *Follower {
	return &Follower{
		Signal:    params.Signal,
		Log:       params.Log,
		Transport: params.Transport,
	}
}

func (f *Follower) Run() {}
