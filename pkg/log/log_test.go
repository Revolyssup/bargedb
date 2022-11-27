package log

import (
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
		Term:      1,
		Index:     1,
		Committed: false,
		Key:       "name",
		Value:     []byte("Ashish"),
	}
	tests.entries["name=Anurag=2"] = Entry{
		Term:      1,
		Index:     1,
		Committed: false,
		Key:       "name",
		Value:     []byte("Anurag"),
	}
	log, err := NewInstance("./mock_index.barge", "./mock_data.barge")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.Remove("./mock_index.barge")
		os.Remove("./mock_data.barge")
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
		entryGot, err := log.Read(Index(i))
		if err != nil {
			t.Error(err)
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

func TestCommit(t *testing.T) {
	type test struct {
		entries map[string]Entry
	}
	tests := test{
		entries: map[string]Entry{},
	}
	tests.entries["name=Ashish=1"] = Entry{
		Term:      1,
		Index:     1,
		Committed: true,
		Key:       "name",
		Value:     []byte("Ashish"),
	}
	tests.entries["name=Anurag=2"] = Entry{
		Term:      1,
		Index:     1,
		Committed: true,
		Key:       "name",
		Value:     []byte("Anurag"),
	}
	log, err := NewInstance("./mock_index_commit.barge", "./mock_data_commit.barge")
	if err != nil {
		t.Error(err)
		return
	}
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
		err = log.Commit(Index(i))
		if err != nil {
			t.Error(err)
			return
		}

		entryGot, err := log.Read(Index(i))
		if err != nil {
			t.Error(err)
			return
		}

		if !entryGot.Committed {
			t.Fatalf("expected %+v got %+v", entry, entryGot)
			return
		}
		if log.LastCommitted != entryGot.Index {
			t.Fatalf("Expected last commited %v, got %v", entryGot.Index, log.LastCommitted)
			return
		}

	}

}
