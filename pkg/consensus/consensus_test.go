package consensus

import (
	"testing"
	"time"

	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/Revolyssup/bargedb/pkg/store"
	"github.com/Revolyssup/bargedb/pkg/transport"
)

func TestAppendEntries(t *testing.T) {
	log1, _ := log.NewInstance("./index.barge", "./data.barge")
	t1 := transport.MockInstance{}

	in1 := New(&t1, store.New(store.INMEM), log1, config{
		ID:          "./tests_config/barge.id",
		CurrentTerm: "./tests_config/barge.currentTerm",
		VotedFor:    "./tests_config/barge.votedFor",
	})

	log2, _ := log.NewInstance("./index2.barge", "./data2.barge")
	t2 := transport.MockInstance{}
	in2 := New(&t2, store.New(store.INMEM), log2, config{
		ID:          "./tests_config/barge copy.id",
		CurrentTerm: "./tests_config/barge copy.currentTerm",
		VotedFor:    "./tests_config/barge copy.votedFor",
	})
	transport.ConnectMockInstances(&t1, &t2)

	//Start both instances, carry out a bunch of operations and then perform queries onto the underlying store to test whether the final state matches.
	in1.Start(&FollowerState{
		timeout: 300 * time.Millisecond,
	})
	in2.Start(&FollowerState{
		timeout: 300 * time.Millisecond,
	})
}
