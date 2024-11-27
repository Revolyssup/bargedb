package leader

import "github.com/revolyssup/bargedb/pkg/consensus/signal"

type Leader struct {
	Signal *signal.Signaller
}
type LeaderParams struct {
	Signal *signal.Signaller
}

func Init(params LeaderParams) *Leader {
	return &Leader{
		Signal: params.Signal,
	}
}

func (l *Leader) Run() {

}

func (l *Leader) becomeFollower() {
	l.Signal.SignalFollower()
}
