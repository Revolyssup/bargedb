package log

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// TestLogReadAndWrite tests both reads and writes together. I know this is technically not a unit test as it has two units, but, uhm, well, Bite me!
func TestLogReadAndWrite(t *testing.T) {
	type test struct {
		entries map[string]Entry
	}
	tests := test{
		entries: map[string]Entry{},
	}
	tests.entries["name=Ashish=1"] = Entry{
		Index:     1,
		Committed: 0,
		Key:       "name",
		Value:     []byte("Ashish"),
	}
	tests.entries["name=Anurag=2"] = Entry{
		Index:     1,
		Committed: 0,
		Key:       "name",
		Value:     []byte("Anurag"),
	}
	log, err := NewInstance("./mock_index.barge", "./mock_data.barge")
	if err != nil {
		t.Error(err)
		return
	}
	// defer func() {
	// 	os.Remove("./mock_index.barge")
	// 	os.Remove("./mock_data.barge")
	// }()
	for keyval, entry := range tests.entries {
		keyvals := strings.Split(keyval, "=")
		key := keyvals[0]
		val := keyvals[1]
		index := keyvals[2]
		err = log.Write(key, []byte(val))
		if err != nil {
			t.Error(err)
			return
		}
		i, _ := strconv.Atoi(index)
		entryGot, err := log.Read(Index(i))
		if err != nil {
			t.Error(err)
			fmt.Println(entry)
			return
		}

		if entryGot.Key != entry.Key {
			t.Fatalf("expected %+v got %+v", entry, entryGot)
			return
		}
		if string(entryGot.Value) != string(entry.Value) {
			t.Fatalf("expected %+v got %+v", entry, entryGot)
			return
		}
	}

}

// TestCommit adds one entry and commits a previous one
func TestCommit(t *testing.T) {
	type test struct {
		entries map[string]Entry
	}
	tests := test{
		entries: map[string]Entry{},
	}
	tests.entries["name=Ashish=1"] = Entry{
		Index:     1,
		Committed: 1,
		Key:       "name",
		Value:     []byte("Ashish"),
	}
	tests.entries["name=Anurag=2"] = Entry{
		Index:     1,
		Committed: 1,
		Key:       "name",
		Value:     []byte("Anurag"),
	}
	tests.entries["name=Utkarsh=3"] = Entry{
		Index:     1,
		Committed: 1,
		Key:       "name",
		Value:     []byte("Utkarsh"),
	}
	log, err := NewInstance("./mock_index_commit.barge", "./mock_data_commit.barge")
	if err != nil {
		t.Error(err)
		return
	}
	log.CurrentTerm = 1
	defer func() {
		os.Remove("./mock_index_commit.barge")
		os.Remove("./mock_data_commit.barge")
	}()
	for keyval, entry := range tests.entries {
		keyvals := strings.Split(keyval, "=")
		key := keyvals[0]
		val := keyvals[1]
		index := keyvals[2]
		err = log.Write(key, []byte(val))
		if err != nil {
			t.Error(err)
			return
		}

		i, _ := strconv.Atoi(index)
		if i > 1 {
			err = log.Commit(Index(i - 1)) //commit the previous entry
			if err != nil {
				t.Error(err)
				return
			}
			entryGot, err := log.Read(Index(i - 1)) //read the previous entry
			if err != nil {
				t.Error(err)
				return
			}
			if entryGot.Committed == 0 {
				t.Fatalf("expected %+v got %+v", entry, entryGot)
				return
			}
			if entryGot.Index > 1 && log.LastCommitted != entryGot.Index {
				t.Fatalf("Expected last commited %v, got %v", entryGot.Index, log.LastCommitted)
				return
			}
		}

	}

}
