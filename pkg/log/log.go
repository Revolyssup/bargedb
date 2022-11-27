package log

import (
	"encoding/gob"
	"os"
	"sync"
)

type IndexEntry struct {
	StartingByte int64
	EndingByte   int64
}
type Instance struct {
	indexFile *os.File
	dataFile  *os.File
	mx        sync.Mutex
}
type Index int
type Entry struct {
	Key   string
	Value []byte
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
