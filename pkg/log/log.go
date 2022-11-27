package log

import (
	"encoding/gob"
	"os"
	"sync"
)

// The barge.wal file will be used to put logs in.
// If the file is not already present in the path provided then, a new file is created in the current directory
type Instance struct {
	Entries       []Entry
	LastCommitted Index
	filepath      string
	mx            sync.Mutex
	dataFile      *os.File
	indexFile     os.File //entries will be persisted on this file
}
type IndexEntry struct {
	StartingByte int64
	EndingByte   int64
}

type Index int
type Entry struct {
	Term      int   //denotes the term in which this log entry was added
	Index     Index //Will increment monotonically with each new entry
	Committed bool  //Denotes if the entry is committed
	Key       string
	Value     []byte
}

func (i *Instance) Write(key string, value []byte) error {
	entry := Entry{
		Key:   key,
		Value: value,
	}
	i.mx.Lock()
	defer i.mx.Unlock()
	enc := gob.NewEncoder(i.dataFile)
	err := enc.Encode(entry)
	if err != nil {
		return err
	}
	// var d gob.Decoder
	// d.Decode()
	return nil
}
