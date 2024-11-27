package transport

import (
	"github.com/google/uuid"
	"github.com/revolyssup/bargedb/pkg/log"
)

type Transport interface {
	AppendEntries(term int, leaderID int, prevLogIndex int, prevLogTerm int, entries []log.Entry, leaderCommit int) (updatedTerm int, success bool)
	RequestVote(term int, candidateID uuid.UUID, lastLogIndex int, lastLogTerm int) (updatedTerm int, voteGranted bool)
}

type LocalTransport struct{}

var _ Transport = &LocalTransport{}

func NewLocalTransport() *LocalTransport {
	return &LocalTransport{}
}
func (t *LocalTransport) AppendEntries(term int, leaderID int, prevLogIndex int, prevLogTerm int, entries []log.Entry, leaderCommit int) (updatedTerm int, success bool) {
	return 0, true
}
func (t *LocalTransport) RequestVote(term int, candidateID uuid.UUID, lastLogIndex int, lastLogTerm int) (updatedTerm int, voteGranted bool) {
	return 0, true
}
