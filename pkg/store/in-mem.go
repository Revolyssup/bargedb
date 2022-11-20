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
	return nil
}

// Exists returns true if the given key exists.
func (ims *InMemStore) Exists(key string) (bool, error) {
	return false, nil
}

// Len returns the number of keys in the storage.
func (ims *InMemStore) Len() (int, error) {
	return 0, nil
}

// PhysicalSnapshot writes snapshot of the storage data to
// the given writer.
func (ims *InMemStore) PhysicalSnapshot(w io.Writer) error {
	return nil
}

// Close closes the storage.
func (ims *InMemStore) Close() error {
	return nil
}
