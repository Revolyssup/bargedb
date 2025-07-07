package transport

import (
	"context"

	"github.com/revolyssup/bargedb/pkg/model"
)

type Transport interface {
	GetOperationChannel() <-chan Operation //Consensus layer will recieve operations from here
	AppendEntries(ctx context.Context, candidateID model.CandidateID, ae AppendEntries) (term uint, success bool)
	RequestVote(ctx context.Context, candidateID model.CandidateID, rv RequestVote) (term uint, voteGranted bool)
}

type Operation interface {
	Type() OperationType
}
type OperationType string

const (
	AppendEntriesOp OperationType = "AppendEntries"
	RequestVoteOp   OperationType = "RequestVote"
)

type AppendEntries struct {
	term         uint
	leaderID     model.CandidateID
	prevLogIndex uint
	prevLogTerm  uint
	entries      []model.LogEntry
	leaderCommit uint
}

func (ar AppendEntries) Type() OperationType {
	return AppendEntriesOp
}

type RequestVote struct {
	term         uint
	candidateID  model.CandidateID
	lastLogIndex uint
	lastLogTerm  uint
}

func (rv RequestVote) Type() OperationType {
	return RequestVoteOp
}
