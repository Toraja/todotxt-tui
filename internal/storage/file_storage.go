package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

// FileStorage implements the Storage interface for local file system operations.
type FileStorage struct {
	watcher *fsnotify.Watcher
}

// NewFileStorage creates a new FileStorage instance.
func NewFileStorage() (*FileStorage, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}
	return &FileStorage{
		watcher: watcher,
	}, nil
}

// Load reads the file and returns its contents.
func (fs *FileStorage) Load(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	return file, nil
}

// Save writes data to the file atomically using a temp file + rename strategy.
func (fs *FileStorage) Save(path string, data io.Reader) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Create temporary file in the same directory
	tmpFile, err := os.CreateTemp(dir, ".todotxt-tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()

	// Ensure cleanup on error
	defer func() {
		if tmpFile != nil {
			tmpFile.Close()
			os.Remove(tmpPath)
		}
	}()

	// Write data to temp file
	if _, err := io.Copy(tmpFile, data); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	// Sync to ensure data is written to disk
	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	// Close temp file before rename
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	tmpFile = nil // Prevent cleanup from closing again

	// Atomic rename
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("failed to rename temp file to %s: %w", path, err)
	}

	return nil
}

// Watch monitors the file for external changes.
func (fs *FileStorage) Watch(path string) (<-chan FileEvent, error) {
	if err := fs.watcher.Add(path); err != nil {
		return nil, fmt.Errorf("failed to watch file %s: %w", path, err)
	}

	eventChan := make(chan FileEvent, 10)

	go func() {
		for {
			select {
			case event, ok := <-fs.watcher.Events:
				if !ok {
					close(eventChan)
					return
				}

				var eventType EventType
				switch {
				case event.Has(fsnotify.Write):
					eventType = EventModified
				case event.Has(fsnotify.Create):
					eventType = EventCreated
				case event.Has(fsnotify.Remove):
					eventType = EventDeleted
				case event.Has(fsnotify.Rename):
					eventType = EventRenamed
				default:
					continue
				}

				eventChan <- FileEvent{
					Type: eventType,
					Path: event.Name,
					Time: time.Now(),
				}

			case err, ok := <-fs.watcher.Errors:
				if !ok {
					close(eventChan)
					return
				}
				// Log error but continue watching
				_ = err // TODO: proper error handling
			}
		}
	}()

	return eventChan, nil
}

// Exists checks if the file exists.
func (fs *FileStorage) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("failed to check if file exists: %w", err)
}

// Create creates a new empty file at the given path.
func (fs *FileStorage) Create(path string) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()

	return nil
}

// GetModificationTime returns the last modification time of the file.
func (fs *FileStorage) GetModificationTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get file info: %w", err)
	}
	return info.ModTime(), nil
}

// Close closes the file watcher.
func (fs *FileStorage) Close() error {
	if fs.watcher != nil {
		return fs.watcher.Close()
	}
	return nil
}
