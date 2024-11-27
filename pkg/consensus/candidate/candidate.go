package candidate

import "github.com/revolyssup/bargedb/pkg/consensus/signal"

type Candidate struct {
	Signal *signal.Signaller
}

type CandidateParams struct {
	Signal *signal.Signaller
}

func Init(params CandidateParams) *Candidate {
	return &Candidate{
		Signal: params.Signal,
	}
}

func (c *Candidate) Run() {}
