package transport

type Transport interface {
	AppendEntries()
	RequestVote()
}
