package transport

import "github.com/google/uuid"

type MockInstancesBroadcaster struct {
	instances map[uuid.UUID]MockInstance
}

func NewMockInstanceBroadcaster() *MockInstancesBroadcaster {
	return &MockInstancesBroadcaster{
		instances: make(map[uuid.UUID]MockInstance),
	}
}

//The API exposed by function here will be the same as APIs exposed by RPC implemented transport layer.
//The functions here will be used for testing the consensus between consensus instances(running in seperate go routines in tests)

// This will implement the Transport interface for testing purposes.
// It will use internal go channels and during testing.
// Multiple mockinstances will talk to each other through the same instance of MockInstanceBroadcaster
type MockInstance struct {
	broadcaster *MockInstancesBroadcaster
	peerid      uuid.UUID
}
