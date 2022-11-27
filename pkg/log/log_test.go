package log

import (
	"os"
	"testing"
)

func TestLogReadAndWrite(t *testing.T) {
	entry := Entry{
		Term:      1,
		Index:     1,
		Committed: false,
		Key:       "name",
		Value:     []byte("Ashish"),
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
	err = log.Write("name", []byte("Ashish"))
	if err != nil {
		t.Error(err)
		return
	}
	entryGot, err := log.Read(1)
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
