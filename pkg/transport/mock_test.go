package transport

import (
	"testing"

	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/google/uuid"
)

func TestMockTransport(t *testing.T) {
	mt1id := uuid.New()
	mt2id := uuid.New()
	mt3id := uuid.New()
	me1 := MockExecutor{alwaysReturnTerm: 1}
	me2 := MockExecutor{alwaysReturnTerm: 2}
	me3 := MockExecutor{alwaysReturnTerm: 3}
	mt1 := MockInstance{Exec: &me1, ID: mt1id}
	mt2 := MockInstance{Exec: &me2, ID: mt2id}
	mt3 := MockInstance{Exec: &me3, ID: mt3id}
	ConnectMockInstances(&mt1, &mt2, &mt3)
	response := mt1.AppendEntries(0, uuid.UUID{}, log.Index(0), 0, nil, 0)
	for id, res := range response {
		if id == mt1id {
			if res["currentTerm"] != 1 {
				t.Errorf("Expected current term %v and recieved %v", 1, res["currentTerm"])
			}
		}
		if id == mt2id {
			if res["currentTerm"] != 2 {
				t.Errorf("Expected current term %v and recieved %v", 2, res["currentTerm"])
			}
		}
		if id == mt3id {
			if res["currentTerm"] != 3 {
				t.Errorf("Expected current term %v and recieved %v", 3, res["currentTerm"])
			}
		}
	}
}
