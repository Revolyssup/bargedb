package transport

import (
	"context"
	"fmt"

	"github.com/revolyssup/bargedb/pkg/model"
)

var _ Transport = (*MockTransport)(nil)

type MockTransport struct {
	ID               model.CandidateID // Candidate ID for the mock transport
	operationChan    chan OperationWithResponseChan
	candidateMsgChan map[model.CandidateID]chan OperationWithResponseChan // Channel for each candidate ID to send messages
}

func NewCandidateMsgChan(candidateID ...model.CandidateID) map[model.CandidateID]chan OperationWithResponseChan {
	candidateMsgChan := make(map[model.CandidateID]chan OperationWithResponseChan)
	for _, id := range candidateID {
		candidateMsgChan[id] = make(chan OperationWithResponseChan, 10) // Buffered channel for testing
	}
	return candidateMsgChan

}
func NewMockTransport(ID model.CandidateID, opchan chan OperationWithResponseChan) *MockTransport {
	return &MockTransport{
		ID:            ID,
		operationChan: opchan, // Buffered channel for testing
	}
}

func (m *MockTransport) GetOperationChannel() <-chan OperationWithResponseChan {
	return m.operationChan
}

func (m *MockTransport) Start() error { return nil }

// This will be done by consensus layer. Only for testing purposes
func (m *MockTransport) Recieve() <-chan string {
	s := make(chan string, 10)
	go func() {
		// Only read from this transport's own operation channel
		for op := range m.operationChan {
			switch op.Type() {
			case AppendEntriesOp:
				ae := op.Operation.(AppendEntries)
				s <- "Received AppendEntries from " + string(ae.leaderID) + " with term " + fmt.Sprintf("%d", ae.term)
				//The below response in real will be generated in the consensus layer and sent back on Response channel
				// and after receving response, the transport layer (RPC) will send back response.
				op.Response <- AppendEntriesResponse{
					term:    ae.term + 1,
					success: true,
				}
				close(op.Response)
			case RequestVoteOp:
				rv := op.Operation.(RequestVote)
				s <- "Received RequestVote from " + string(rv.candidateID) + " with term " + fmt.Sprintf("%d", rv.term)
				op.Response <- RequestVoteResponse{
					term:        rv.term + 1,
					voteGranted: true,
				}
				close(op.Response)
			}
		}
	}()
	return s
}

func (m *MockTransport) AppendEntries(ctx context.Context, candidateID model.CandidateID, ae AppendEntries) AppendEntriesResponse {
	if msgChan, ok := m.candidateMsgChan[candidateID]; ok {
		op := OperationWithResponseChan{
			Operation: ae,
			Response:  make(chan Operation, 1),
		}
		msgChan <- op
		res := <-op.Response
		ae, _ := res.(AppendEntriesResponse)
		return ae
	}
	return AppendEntriesResponse{
		term:    ae.term,
		success: false,
	}
}

func (m *MockTransport) RequestVote(ctx context.Context, candidateID model.CandidateID, rv RequestVote) RequestVoteResponse {
	if msgChan, ok := m.candidateMsgChan[candidateID]; ok {
		op := OperationWithResponseChan{
			Operation: rv,
			Response:  make(chan Operation, 1),
		}
		msgChan <- op //In reality this willbe replaced by actual RPC call and response
		res := <-op.Response
		rv, _ := res.(RequestVoteResponse)
		return rv
	}
	return RequestVoteResponse{
		voteGranted: false,
		term:        rv.term,
	}
}
