package storage

import (
	"io"
	"time"
)

// Storage defines the interface for file I/O and persistence operations.
type Storage interface {
	// Load reads the file and returns its contents.
	Load(path string) (io.ReadCloser, error)

	// Save writes data to the file atomically.
	Save(path string, data io.Reader) error

	// Watch monitors the file for external changes.
	// Returns a channel that sends FileEvents when changes occur.
	Watch(path string) (<-chan FileEvent, error)

	// Exists checks if the file exists.
	Exists(path string) (bool, error)

	// Create creates a new empty file at the given path.
	Create(path string) error

	// GetModificationTime returns the last modification time of the file.
	GetModificationTime(path string) (time.Time, error)
}
