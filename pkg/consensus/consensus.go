package consensus

import (
	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/Revolyssup/bargedb/pkg/store"
	"github.com/Revolyssup/bargedb/pkg/transport"
)

type Instance struct {
	Transport transport.Transport
	Log       log.Instance
	Store     store.Storage
}
