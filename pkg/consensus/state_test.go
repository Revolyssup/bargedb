package consensus

import (
	"testing"

	"github.com/revolyssup/bargedb/pkg/model"
	"gotest.tools/assert"
)

func TestState(t *testing.T) {
	state, err := NewFileBasedPersistentState("testdata/state.txt")
	if err != nil {
		t.Fatalf("Failed to create persistent state: %v", err)
	}
	defer state.Close()
	state.SetCurrentTerm(1)
	candidateID := model.CandidateID("candidate1")
	state.SetVotedFor(candidateID)
	assert.Equal(t, state.GetCurrentTerm(), uint(1), "Expected current term to be 1")
	assert.Equal(t, state.GetVotedFor(), candidateID, "Expected voted for to be candidate1")
	state.SetCurrentTerm(2)
	state2, err := NewFileBasedPersistentState("testdata/state.txt")
	if err != nil {
		t.Fatalf("Failed to create persistent state: %v", err)
	}
	defer state2.Close()
	assert.Equal(t, state2.GetCurrentTerm(), uint(2), "Expected current term to be 1 after reloading")
	assert.Equal(t, state2.GetVotedFor(), candidateID, "Expected voted for to be candidate1 after reloading")
}
