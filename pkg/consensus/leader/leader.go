package leader

import (
	"github.com/revolyssup/bargedb/pkg/consensus/signal"
	"github.com/revolyssup/bargedb/pkg/log"
	"github.com/revolyssup/bargedb/pkg/transport"
)

type Leader struct {
	Signal    *signal.Signaller
	Log       *log.Log
	Transport transport.Transport
}
type LeaderParams struct {
	Signal    *signal.Signaller
	Log       *log.Log
	Transport transport.Transport
}

func Init(params LeaderParams) *Leader {
	return &Leader{
		Signal:    params.Signal,
		Log:       params.Log,
		Transport: params.Transport,
	}
}

func (l *Leader) Run() {

}

func (l *Leader) becomeFollower() {
	l.Signal.SignalFollower()
}
