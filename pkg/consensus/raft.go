package consensus

import (
	"time"

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
	operationChan <-chan transport.OperationWithResponseChan
	timeout       int
}

type Config struct {
	MetadataPath string
	CandidateID  model.CandidateID
	Peers        []model.CandidateID
	Timeout      int
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
		timeout:         c.Timeout,
	}, nil
}

func (b *Barge) Start() {
	err := b.transport.Start()
	if err != nil {
		panic(err)
	}
	b.runElection() //It will start with runFollower
}
func (b *Barge) runElection() {
	//Listen for all operations
	closeChan := make(chan struct{})
	closeDaemon := make(chan struct{})
	//How much knowledge does this daemon have?
	// It should know when the runElection is going to exit and it should return as well.
	//This means the logic of state transition goes here.
	// The cleanup will still be handled by the main routine???
	go func() {
		defer func() {
			close(closeChan)
		}()
		for {
			t := time.After(time.Second * time.Duration(b.timeout))
			select {
			// case ev := <-b.operationChan:

			case <-t:
				return
			case <-closeDaemon:
				return
			}
		}
	}()
	b.persistentState.SetCurrentTerm(b.persistentState.GetCurrentTerm() + 1)
	b.persistentState.SetVotedFor(b.ID)

}
