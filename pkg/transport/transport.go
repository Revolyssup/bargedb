package transport

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
