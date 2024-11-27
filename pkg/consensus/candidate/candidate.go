package candidate

import (
	"github.com/google/uuid"
	"github.com/revolyssup/bargedb/pkg/consensus/signal"
	"github.com/revolyssup/bargedb/pkg/log"
	"github.com/revolyssup/bargedb/pkg/transport"
)

type Candidate struct {
	Signal    *signal.Signaller
	Log       *log.Log
	Transport transport.Transport
	ID        uuid.UUID
	CurrTerm  int
	VotedFor  int

	CommitIndex int
	LastApplied int
}

type CandidateParams struct {
	Signal    *signal.Signaller
	Log       *log.Log
	Transport transport.Transport

	CurrTerm int
	VotedFor int

	CommitIndex int
	LastApplied int
}

func Init(params CandidateParams) *Candidate {
	return &Candidate{
		Signal:      params.Signal,
		Log:         params.Log,
		Transport:   params.Transport,
		CurrTerm:    params.CurrTerm,
		VotedFor:    params.VotedFor,
		CommitIndex: params.CommitIndex,
		LastApplied: params.LastApplied,
	}
}

func (c *Candidate) Run() {
	//First thing a candidate does is tries to acquire votes
	//If it gets majority votes, it becomes a leader
	//If it gets AppendEntries from a leader, it becomes a follower
	c.Transport.RequestVote(c.CurrTerm, c.ID, c.Log.LastLogIndex(), c.Log.LastLogTerm())
}
