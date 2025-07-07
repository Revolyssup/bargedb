package transport

import (
	"context"
	"testing"

	"github.com/revolyssup/bargedb/pkg/model"
	"gotest.tools/assert"
)

func TestMock(t *testing.T) {
	candidate1 := model.CandidateID("candidate1")
	candidate2 := model.CandidateID("candidate2")
	candidate3 := model.CandidateID("candidate3")

	cidops := NewCandidateMsgChan(candidate1, candidate2, candidate3)
	var t1 *MockTransport
	var t2 *MockTransport
	var t3 *MockTransport
	// var t1recieve <-chan string
	var t2recieve <-chan string
	// var t3recieve <-chan string
	for cid, ops := range cidops {
		switch cid {
		case candidate1:
			t1 = NewMockTransport(cid, ops)
			t1.candidateMsgChan = cidops
			// t1recieve = t1.Recieve()
		case candidate2:
			t2 = NewMockTransport(cid, ops)
			t2.candidateMsgChan = cidops
			t2recieve = t2.Recieve()
		case candidate3:
			t3 = NewMockTransport(cid, ops)
			t3.candidateMsgChan = cidops
			// t3recieve = t3.Recieve()
		}
	}
	if t1 == nil || t2 == nil || t3 == nil {
		t.Fatal("MockTransport instances not created correctly")
	}

	//send AppendEntries from candidate1 to candidate2
	ae := AppendEntries{
		term:         1,
		leaderID:     candidate1,
		prevLogIndex: 0,
		prevLogTerm:  0,
		entries:      []model.LogEntry{{Term: 1, Data: []byte("entry1"), Index: 1}},
		leaderCommit: 0,
	}
	term, success := t1.AppendEntries(context.Background(), candidate2, ae)
	assert.Equal(t, term, uint(1), "Expected term to be 1")
	assert.Equal(t, success, true, "Expected AppendEntries to succeed")
	msg := <-t2recieve
	assert.Equal(t, msg, "Received AppendEntries from candidate1 with term 1", "Expected message to match")

	//send RequestVote from candidate1 to candidate2
	rv := RequestVote{
		term:         1,
		candidateID:  candidate1,
		lastLogIndex: 0,
		lastLogTerm:  0,
	}
	term, voteGranted := t1.RequestVote(context.Background(), candidate2, rv)
	assert.Equal(t, term, uint(1), "Expected term to be 1")
	assert.Equal(t, voteGranted, true, "Expected RequestVote to succeed")
	msg = <-t2recieve
	assert.Equal(t, msg, "Received RequestVote from candidate1 with term 1", "Expected message to match")
}
