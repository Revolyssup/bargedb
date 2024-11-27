package log

type Log struct {
	entries []Entry
}

func (l *Log) LastLogIndex() int {
	return l.entries[len(l.entries)-1].Index
}

func (l *Log) LastLogTerm() int {
	return l.entries[len(l.entries)-1].Term
}

type Entry struct {
	Term  int
	Index int
	Data  interface{}
}
