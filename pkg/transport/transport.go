package transport

import (
	"github.com/Revolyssup/bargedb/pkg/log"
	"github.com/google/uuid"
)

/*
1. Transport layer maps a peer UUID(at consensus layer) to an actual address. For RPCs this will be the address of the RPC server while at the same time accepts requests as an RPC server.
2. Internally for mocking, the implementation can just be a mapping of peer id to the appropriate go routine in which the peer's transport instance is running.
3. Different Transport implementation can have different configuration, consensus is agnostic of what type of transport is being used.
4. Whenever a new Transport instance is returned, it should not require any extra calls to start. Setting up appropriate listeners and other work has to be done in New function.
*/
type Transport interface {
	//For senders
	/*
		Arguments:
			term: leader’s term
			leaderId: so follower can redirect clients
			prevLogIndex: index of log entry immediately preceding new ones
			prevLogTerm: term of prevLogIndex entry
			entries[]: log entries to store (empty for heartbeat; may send more than one for efficiency) leaderCommit leader’s commitIndex
		Results:
			term: currentTerm, for leader to update itself
			success: true if follower contained entry matching prevLogIndex and prevLogTerm
	*/
	AppendEntries(term int, leaderID uuid.UUID, prevLogIndex log.Index, prevLogEntry int, entries []log.Entry, leaderCommit int) (int, bool)
	/*
		Arguments:
			term: candidate’s term
			candidateId: candidate requesting vote
			lastLogIndex: index of candidate’s last log entry (§5.4)
			lastLogTerm: term of candidate’s last log entry (§5.4)
		Results:
			term: currentTerm, for candidate to update itself
			voteGranted: true means candidate received vote
	*/
	RequestVote(term int, candidateID uuid.UUID, lastLogIndex log.Index, lastLogTerm int) (int, bool)
	//For recievers
	Listen()
}

// This will implement the Transport interface with RPCs
type Instance struct {
}
