package log

import "os"

// The barge.wal file will be used to put logs in.
// If the file is not already present in the path provided then, a new file is created in the current directory
type Instance struct {
	Entries       []Entry
	LastCommitted Index
	filepath      string
	file          os.File //entries will be persisted on this file
}
type Index int
type Entry struct {
	Term      int   //denotes the term in which this log entry was added
	Index     Index //Will increment monotonically with each new entry
	Committed bool  //Denotes if the entry is committed
	Key       string
	Value     []byte
}

func (i *Instance) Write(key string, value interface{}) {}
