package transport

import (
	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/google/uuid"
)

type Executer interface {
	RecievedAppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) (int, bool)
	RespondVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) (int, bool)
}

type MockExecutor struct {
	alwaysReturnTerm int
}

func (e *MockExecutor) RecievedAppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) (int, bool) {
	return e.alwaysReturnTerm, true
}
func (e *MockExecutor) RespondVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) (int, bool) {
	return e.alwaysReturnTerm, true
}
