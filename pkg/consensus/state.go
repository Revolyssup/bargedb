package consensus

import (
	"sync"
)

type State interface {
	ApplyAction(consensus *Instance, act Action) State //Applies the Action and returns a new state
	//Will be run by Main Start routine after cancelling the previous run.
	//Start should exist once it recieves signal on stop channel
	Start(stop <-chan interface{}) //All state changes are manageed by ApplyAction and Start will be called ApplyAction brings the state from where it can START
}

var stateLock sync.Mutex = sync.Mutex{} //Any operation that potentially changes the states should first acquire a lock

// The states should not have redundant data that is stored at the consensus layer.
// The data here will only be relevant during the course of state change.
type LeaderState struct {
}

func (l LeaderState) ApplyAction(consensus *Instance, act Action) State {
	switch act.Name {
	case TIMEOUT:

	}
	return nil
}
