package consensus

import (
	"fmt"
	"testing"

	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/Revolyssup/bargedb/pkg/store"
	"github.com/Revolyssup/bargedb/pkg/transport"
)

func TestAppendEntries(t *testing.T) {
	log1, _ := log.NewInstance("./index.barge", "./data.barge")
	t1 := transport.MockInstance{}
	in1 := New(&t1, store.New(store.INMEM), log1)

	log2, _ := log.NewInstance("./index2.barge", "./data2.barge")
	t2 := transport.MockInstance{}
	in2 := New(&t2, store.New(store.INMEM), log2)
	transport.ConnectMockInstances(&t1, &t2)

	//Start both instances, carry out a bunch of operations and then perform queries onto the underlying store to test whether the final state matches.
	fmt.Println(in1, in2)
	in1.Start()
	in2.Start()
}
