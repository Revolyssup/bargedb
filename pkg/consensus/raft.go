package consensus

import (
	"github.com/revolyssup/bargedb/pkg/log"
	"github.com/revolyssup/bargedb/pkg/model"
	"github.com/revolyssup/bargedb/pkg/storage"
	"github.com/revolyssup/bargedb/pkg/transport"
)

type Barge struct {
	ID              model.CandidateID
	Peers           []model.CandidateID
	transport       transport.Transport
	store           storage.Store
	log             log.Log
	persistentState PersistentState
	commitIndex     uint
	lastApplied     uint
	//for leaders
	nextIndex     map[model.CandidateID]uint
	matchIndex    map[model.CandidateID]uint
	operationChan <-chan transport.Operation
}

type Config struct {
	MetadataPath string
	CandidateID  model.CandidateID
	Peers        []model.CandidateID
}

func NewBarge(t transport.Transport, store storage.Store, log log.Log, c Config) (*Barge, error) {
	pstate, err := NewFileBasedPersistentState(c.MetadataPath)
	if err != nil {
		return nil, err
	}

	return &Barge{
		ID:              c.CandidateID,
		transport:       t,
		store:           store,
		log:             log,
		persistentState: pstate,
		Peers:           c.Peers,
		operationChan:   t.GetOperationChannel(),
	}, nil
}

func (b *Barge) runElection() {
	b.persistentState.SetCurrentTerm(b.persistentState.GetCurrentTerm() + 1)
	b.persistentState.SetVotedFor(b.ID)
}
