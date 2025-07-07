package transport

import (
	"context"
	"fmt"

	"github.com/revolyssup/bargedb/pkg/model"
)

type MockTransport struct {
	ID               model.CandidateID // Candidate ID for the mock transport
	operationChan    chan Operation
	candidateMsgChan map[model.CandidateID]chan Operation // Channel for each candidate ID to send messages
}

func NewCandidateMsgChan(candidateID ...model.CandidateID) map[model.CandidateID]chan Operation {
	candidateMsgChan := make(map[model.CandidateID]chan Operation)
	for _, id := range candidateID {
		candidateMsgChan[id] = make(chan Operation, 10) // Buffered channel for testing
	}
	return candidateMsgChan

}
func NewMockTransport(ID model.CandidateID, opchan chan Operation) *MockTransport {
	return &MockTransport{
		ID:            ID,
		operationChan: opchan, // Buffered channel for testing
	}
}

func (m *MockTransport) GetOperationChannel() <-chan Operation {
	return m.operationChan
}

// This will be done by consensus layer. Only for testing purposes
func (m *MockTransport) Recieve() <-chan string {
	s := make(chan string, 10)
	go func() {
		// Only read from this transport's own operation channel
		for op := range m.operationChan {
			switch op.Type() {
			case AppendEntriesOp:
				ae := op.(AppendEntries)
				s <- "Received AppendEntries from " + string(ae.leaderID) + " with term " + fmt.Sprintf("%d", ae.term)
			case RequestVoteOp:
				rv := op.(RequestVote)
				s <- "Received RequestVote from " + string(rv.candidateID) + " with term " + fmt.Sprintf("%d", rv.term)
			}
		}
	}()
	return s
}

func (m *MockTransport) AppendEntries(ctx context.Context, candidateID model.CandidateID, ae AppendEntries) (term uint, success bool) {
	if msgChan, ok := m.candidateMsgChan[candidateID]; ok {
		msgChan <- ae
		return ae.term, true // Simulate success
	}
	return 0, false // Simulate failure if candidate ID not found
}

func (m *MockTransport) RequestVote(ctx context.Context, candidateID model.CandidateID, rv RequestVote) (term uint, voteGranted bool) {
	if msgChan, ok := m.candidateMsgChan[candidateID]; ok {
		msgChan <- rv
		return rv.term, true // Simulate success
	}
	return 0, false // Simulate failure if candidate ID not found
}
