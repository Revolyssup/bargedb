package transport

import (
	"sync"

	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/google/uuid"
)

//The API exposed by function here will be the same as APIs exposed by RPC implemented transport layer.
//The functions here will be used for testing the consensus between consensus instances(running in seperate go routines in tests)

// This will implement the Transport interface for testing purposes.
// It will use internal go channels and during testing.
// Multiple mockinstances will talk to each other through the same instance of MockInstanceBroadcaster
type MockInstance struct {
	ID    uuid.UUID
	Exec  Executer
	Peers []MockInstance
}

func ConnectMockInstances(mis ...*MockInstance) {
	for i, mi := range mis {
		for j, mj := range mis {
			if i != j {
				mi.Peers = append(mi.Peers, *mj)
			}
		}
	}
}
func (mi *MockInstance) Listen() {

}
func (mi *MockInstance) RegisterExecuter(e Executer) {
	mi.Exec = e
}

func (mi *MockInstance) AppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) chan map[uuid.UUID]map[string]interface{} {
	var response = make(chan map[uuid.UUID]map[string]interface{})
	var wg sync.WaitGroup
	for _, peer := range mi.Peers {
		wg.Add(1)
		go func(peer MockInstance) {
			defer wg.Done()
			resint, _ := peer.Exec.RecievedAppendEntries(term, leaderID, prevLogIndex, prevLogEntry, entries, leaderCommit)
			res := make(map[uuid.UUID]map[string]interface{})
			res[peer.ID] = make(map[string]interface{})
			res[peer.ID]["currentTerm"] = resint
			response <- res
		}(peer)
	}
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(response)
	}(&wg)
	return response
}

func (mi *MockInstance) RequestVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) chan map[uuid.UUID]map[string]interface{} {
	var response = make(chan map[uuid.UUID]map[string]interface{})
	var wg sync.WaitGroup
	for _, peer := range mi.Peers {
		wg.Add(1)
		go func(peer MockInstance) {
			defer wg.Done()
			resint, voted := peer.Exec.RespondVote(term, candidateID, lastLogIndex, lastLogTerm)
			res := make(map[uuid.UUID]map[string]interface{})
			res[peer.ID] = make(map[string]interface{})
			res[peer.ID]["currentTerm"] = resint
			res[peer.ID]["voted"] = voted
			response <- res
		}(peer)
	}
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(response)
	}(&wg)
	return response
}
