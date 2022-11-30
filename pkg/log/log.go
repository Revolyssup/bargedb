package log

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"os"
	"sync"
)

// The barge.wal file will be used to put logs in.
// If the file is not already present in the path provided then, a new file is created in the current directory
type Instance struct {
	LastCommitted Index
	CurrentTerm   int
	mx            sync.Mutex
	Lastindex     int
	dataFile      *os.File //entries will be persisted on this file
	//indexFile will contain the offset to a particular entry in datafile. ith entry here will contain the size and offset of ith entry in the datafile
	//indexFile will have entries of fixed size of 12 bytes. Starting 8 bytes will have offset, 4 bytes will contain size(lenbytes).
	indexFile *os.File
}

func NewInstance(indexfilepath string, datafilepath string) (*Instance, error) {
	indexFile, err := os.Create(indexfilepath)
	if err != nil {
		return nil, err
	}
	dataFile, err := os.Create(datafilepath)
	if err != nil {
		return nil, err
	}
	ins := Instance{
		indexFile: indexFile,
		dataFile:  dataFile,
		mx:        sync.Mutex{},
	}
	return &ins, nil
}

type Index int
type Entry struct {
	Term      int   //denotes the term in which this log entry was added
	Index     Index //Will increment monotonically with each new entry
	Committed int   //Denotes if the entry is committed -- boolean is not used because I am using json.Marshal to generate bytes which change in length as we go from "false" to "true"(1 byte less). 0 and 1 on the other hand both consume same amount of bytes
	Key       string
	Value     []byte
}

func (i *Instance) Commit(commitindex Index) error {
	i.mx.Lock()
	defer i.mx.Unlock()
	indexOffset := IndexSize * (int(commitindex) - 1)
	index := make([]byte, IndexSize)
	_, err := i.indexFile.ReadAt(index, int64(indexOffset))
	if err != nil {
		return err
	}
	offset := index[0:8]
	lens := index[8:]
	bytOffset := bytes.NewBuffer(offset)
	dataoffset, err := binary.ReadVarint(bytOffset)
	if err != nil {
		return err
	}

	bytLen := bytes.NewBuffer(lens)
	datalen, err := binary.ReadVarint(bytLen)
	if err != nil {
		return err
	}

	//use the offset and datalen to read data
	data := make([]byte, datalen)
	_, err = i.dataFile.ReadAt(data, dataoffset)
	if err != nil {
		return err
	}
	var entry Entry
	err = json.Unmarshal(data, &entry)
	if err != nil {
		return err
	}
	entry.Committed = 1
	eb, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	_, err = i.dataFile.WriteAt(eb, dataoffset)
	if err != nil {
		return err
	}
	i.LastCommitted = commitindex
	return nil
}
func (i *Instance) Write(key string, value []byte) error {
	i.mx.Lock()
	defer i.mx.Unlock()
	entry := Entry{
		Key:       key,
		Value:     value,
		Committed: 0,
		Index:     Index(i.Lastindex + 1),
		Term:      i.CurrentTerm,
	}
	eb, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	info, err := i.dataFile.Stat()
	if err != nil {
		return err
	}
	offset := info.Size()
	lens, err := i.dataFile.Write(eb)
	if err != nil {
		return err
	}
	offsetBytes := make([]byte, IndexOffsetSize)
	binary.PutVarint(offsetBytes, offset)
	lensBytes := make([]byte, IndexDatalenSize)
	binary.PutVarint(lensBytes, int64(lens))
	totalbytes := append(offsetBytes, lensBytes...)

	//Write to indexfile
	_, err = i.indexFile.Write(totalbytes)
	if err != nil {
		return err
	}
	i.Lastindex++
	return nil
}

const IndexSize int = 12
const IndexOffsetSize int = 8
const IndexDatalenSize int = 4

func (i *Instance) Read(entryIndex Index) (*Entry, error) {
	indexOffset := IndexSize * (int(entryIndex) - 1)
	index := make([]byte, IndexSize)
	_, err := i.indexFile.ReadAt(index, int64(indexOffset))
	if err != nil {
		return nil, err
	}
	offset := index[0:8]
	lens := index[8:]
	bytOffset := bytes.NewBuffer(offset)
	dataoffset, err := binary.ReadVarint(bytOffset)
	if err != nil {
		return nil, err
	}

	bytLen := bytes.NewBuffer(lens)
	datalen, err := binary.ReadVarint(bytLen)
	if err != nil {
		return nil, err
	}

	//use the offset and datalen to read data
	data := make([]byte, datalen)
	_, err = i.dataFile.ReadAt(data, dataoffset)
	if err != nil {
		return nil, err
	}
	entry := Entry{}
	err = json.Unmarshal(data, &entry)
	return &entry, err
}
