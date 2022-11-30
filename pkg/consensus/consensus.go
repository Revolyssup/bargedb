package consensus

import (
	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/Revolyssup/bargedb/pkg/store"
	"github.com/Revolyssup/bargedb/pkg/transport"
	"github.com/google/uuid"
)

type Instance struct {
	Transport *transport.Transport //The transport layer
	Store     store.Storage        //The state machine for this RAFT
	Log       log.Instance         //The underlying WAL
	//persisted on non-volatile storage
	currentTerm int       //latest term server has seen (initialized to 0 on first boot, increases monotonically)
	votedFor    uuid.UUID //candidateId that received vote in current term (or null if none)

	//volatile state
	commitIndex int //index of highest log entry known to be committed (initialized to 0, increases monotonically)
	lastApplied int //index of highest log entry applied to state machine (initialized to 0, increases monotonically)

	//volatile states only for leaders
	nextIndex  map[uuid.UUID]log.Index //for each server, index of the next log entry to send to that server (initialized to leader last log index + 1)
	matchIndex map[uuid.UUID]log.Index //for each server, index of highest log entry known to be replicated on server (initialized to 0, increases monotonically)
}

// RecievedAppendEntries will be called when AppendEntries is detected from transport layer via Listen()
func (i *Instance) RecievedAppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) (int, bool) {
	return 0, false
}

// RespondVote will be called when RequestVote is detected from transport layer via Listen()
func (i *Instance) RespondVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) (int, bool) {
	return 0, false
}
