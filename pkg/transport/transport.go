package transport

import (
	"context"

	"github.com/revolyssup/bargedb/pkg/model"
)

type Transport interface {
	GetOperationChannel() <-chan OperationWithResponseChan //Consensus layer will recieve operations from here
	AppendEntries(ctx context.Context, candidateID model.CandidateID, ae AppendEntries) AppendEntriesResponse
	RequestVote(ctx context.Context, candidateID model.CandidateID, rv RequestVote) RequestVoteResponse
	Start() error
}

type OperationWithResponseChan struct {
	Operation
	Response chan Operation //After processing request, transport layer will wait for the response here
}

type Operation interface {
	Type() OperationType
}
type OperationType string

const (
	AppendEntriesOp  OperationType = "AppendEntries"
	RequestVoteOp    OperationType = "RequestVote"
	AppendEntriesRes OperationType = "AppendEntriesResponse"
	RequestVoteRes   OperationType = "RequestVoteResponse"
)

type AppendEntriesResponse struct {
	term    uint
	success bool
}

func (ar AppendEntriesResponse) Type() OperationType {
	return AppendEntriesRes
}

type RequestVoteResponse struct {
	term        uint
	voteGranted bool
}

func (rv RequestVoteResponse) Type() OperationType {
	return RequestVoteRes
}

type AppendEntries struct {
	term         uint
	leaderID     model.CandidateID
	prevLogIndex uint
	prevLogTerm  uint
	entries      []model.LogEntry
	leaderCommit uint
}

func (ar AppendEntries) Term() uint {
	return ar.term
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
