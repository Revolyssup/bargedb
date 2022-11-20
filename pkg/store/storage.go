package store

import (
	"io"
)

type Type string

const (
	USE   Type = "USE"
	INMEM Type = "INMEM"
)

func New(t Type) Storage {
	switch t {
	case INMEM:
		return NewInmemStore()
		// case USE:
		// 	return use.New()
	}
	return nil
}

// This storage layer is to be made compatible with https://github.com/utkarsh-pro/use/blob/master/pkg/storage/storage.go
type Storage interface {
	// Init configures the storage.
	Init() error

	// Get returns the value for the given key.
	Get(key string) ([]byte, error)

	// Set sets the value for the given key.
	Set(key string, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key string) error

	// Exists returns true if the given key exists.
	Exists(key string) (bool, error)

	// Len returns the number of keys in the storage.
	Len() (int, error)

	// PhysicalSnapshot writes snapshot of the storage data to
	// the given writer.
	PhysicalSnapshot(w io.Writer) error

	// Close closes the storage.
	Close() error
}
