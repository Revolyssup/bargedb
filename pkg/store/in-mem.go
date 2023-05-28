package store

import (
	"fmt"
	"io"
	"sync"
)

// Implements the storage interface
type InMemStore struct {
	data map[string][]byte
	mx   sync.Mutex
}

func NewInmemStore() *InMemStore {
	return &InMemStore{}
}

func (ims *InMemStore) Init() error {
	if ims.data == nil {
		ims.data = make(map[string][]byte)
	}
	return nil
}

// Get returns the value for the given key.
func (ims *InMemStore) Get(key string) ([]byte, error) {
	ims.mx.Lock()
	defer ims.mx.Unlock()
	if ims.data == nil {
		return nil, fmt.Errorf("in memory store not initialized")
	}
	return ims.data[key], nil
}

// Set sets the value for the given key.
func (ims *InMemStore) Set(key string, value []byte) error {
	ims.mx.Lock()
	defer ims.mx.Unlock()
	if ims.data == nil {
		return fmt.Errorf("in memory store not initialized")
	}
	ims.data[key] = value
	return nil
}

// Delete deletes the value for the given key.
func (ims *InMemStore) Delete(key string) error {
	ims.mx.Lock()
	defer ims.mx.Unlock()
	if ims.data[key] != nil {
		delete(ims.data, key)
	}
	return nil
}

// Exists returns true if the given key exists.
func (ims *InMemStore) Exists(key string) (bool, error) {
	ims.mx.Lock()
	defer ims.mx.Unlock()
	if ims.data[key] != nil {
		return true, nil
	}
	return false, nil
}

// Len returns the number of keys in the storage.
func (ims *InMemStore) Len() (int, error) {
	ims.mx.Lock()
	defer ims.mx.Unlock()
	count := len(ims.data)
	return count, nil
}

// PhysicalSnapshot writes snapshot of the storage data to
// the given writer.
// Todo:
func (ims *InMemStore) PhysicalSnapshot(w io.Writer) error {
	return nil
}

// Close closes the storage. (Since this is a map implementation, we don't really need to implement Close for now as we are not holding to any file descriptor)
func (ims *InMemStore) Close() error {
	return nil
}
