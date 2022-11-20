package transport

/*
1. Transport layer maps a peer UUID(at consensus layer) to an actual address. For RPCs this will be the address of the RPC server while at the same time accepts requests as an RPC server.
2. Internally for mocking, the implementation can just be a mapping of peer id to the appropriate go routine in which the peer's transport instance is running.
3. Different Transport implementation can have different configuration, consensus is agnostic of what type of transport is being used.
4. Whenever a new Transport instance is returned, it should not require any extra calls to start. Setting up appropriate listeners and other work has to be done in New function.
*/
type Transport interface {
	//For senders
	AppendEntries()
	RequestVote()
	//For recievers
	Listen()
}

// This will implement the Transport interface with RPCs
type Instance struct {
}
