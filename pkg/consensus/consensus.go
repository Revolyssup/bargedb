package consensus

import (
	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/Revolyssup/bargedb/pkg/transport"
)

type Instance struct {
	Transport transport.Transport
	Log       log.Instance
}

func (i *Instance) AppendEntries() {
	i.Transport.AppendEntries()
}

func (i *Instance) RequestVote() {
	i.Transport.RequestVote()
}
