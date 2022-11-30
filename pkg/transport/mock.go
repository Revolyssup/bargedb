package transport

import (
	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/google/uuid"
)

//The API exposed by function here will be the same as APIs exposed by RPC implemented transport layer.
//The functions here will be used for testing the consensus between consensus instances(running in seperate go routines in tests)

// This will implement the Transport interface for testing purposes.
// It will use internal go channels and during testing.
// Multiple mockinstances will talk to each other through the same instance of MockInstanceBroadcaster
type MockInstance struct {
	broadcaster map[uuid.UUID]
	id          uuid.UUID
}

func (mi *MockInstance) Listen() {

}

func (mi *MockInstance) AppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) (int, bool) {

	return 0, false
}

func (mi *MockInstance) RequestVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) (int, bool) {
	return 0, false
}
