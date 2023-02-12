package consensus

import (
	"context"
	"io/ioutil"
	"strconv"

	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/Revolyssup/bargedb/pkg/store"
	"github.com/google/uuid"
)

type config struct {
	ID          filename `json:"id"`
	CurrentTerm filename `json:"currentTerm"`
	VotedFor    filename `json:"votedFor"`
}

func New(t Transport, s store.Storage, l *log.Instance, cfg config) *Instance {
	i := Instance{
		Transport:   t,
		Store:       s,
		Log:         l,
		id:          cfg.ID,
		currentTerm: cfg.CurrentTerm,
		votedFor:    cfg.CurrentTerm,
		commitIndex: 0,
		lastApplied: 0,
	}
	t.RegisterExecuter(&i)
	return &i
}

type filename string
type Instance struct {
	Transport Transport     //The transport layer
	State     State         //All instances start from FOLLOWER state
	Store     store.Storage //The state machine for this RAFT
	Log       *log.Instance //The underlying WAL
	//persisted on non-volatile storage
	//config consists of below non volatile fields
	id          filename
	currentTerm filename //latest term server has seen (initialized to 0 on first boot, increases monotonically)
	votedFor    filename //candidateId that received vote in current term (or null if none)

	//volatile state
	commitIndex int //index of highest log entry known to be committed (initialized to 0, increases monotonically)
	lastApplied int //index of highest log entry applied to state machine (initialized to 0, increases monotonically)

	//volatile states only for leaders
	nextIndex  map[uuid.UUID]log.Index //for each server, index of the next log entry to send to that server (initialized to leader last log index + 1)
	matchIndex map[uuid.UUID]log.Index //for each server, index of highest log entry known to be replicated on server (initialized to 0, increases monotonically)

	cancelState context.CancelFunc //Calling this function, cancels the run of already running state
}

// RecievedAppendEntries will be called when AppendEntries is detected from transport layer via Listen()
func (i *Instance) RecievedAppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) (int, bool) {
	return 0, false
}

// RespondVote will be called when RequestVote is detected from transport layer via Listen()
func (i *Instance) RespondVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) (int, bool) {
	return 0, false
}

func (i *Instance) setCurrentTerm(term int) {
	err := ioutil.WriteFile(string(i.currentTerm), []byte(strconv.Itoa(term)), 0644)
	if err != nil {
		panic("error writing to current term file:" + err.Error())
	}
}
func (i *Instance) getCurrentTerm() (term int) {
	data, err := ioutil.ReadFile(string(i.currentTerm))
	if err != nil {
		panic("could not read current term: " + err.Error())
	}
	term, err = strconv.Atoi(string(data))
	if err != nil {
		panic("could not read current term: " + err.Error())
	}
	return
}

var stateContext = make(chan interface{}) //should be unbuffered

// Start is the daemon running in the background creating timeouts/generating actions.
// When start is run, cancel the previous Start from another state and run Start again with current state.
// A valid state should be set in the consensus Instance before calling Start().
func (i *Instance) Start(st State) {
	if i.State != nil {
		i.cancelState() //Cancel the run of previous state
	}
	i.State = st
	ctx, cancel := context.WithCancel(context.Background())
	i.cancelState = cancel
	go i.State.Start(ctx)
}
