package storage

import "time"

// EventType defines the type of file system event.
type EventType int

const (
	// EventModified indicates the file was modified.
	EventModified EventType = iota
	// EventCreated indicates the file was created.
	EventCreated
	// EventDeleted indicates the file was deleted.
	EventDeleted
	// EventRenamed indicates the file was renamed.
	EventRenamed
)

// FileEvent represents a file system event.
type FileEvent struct {
	// Type is the type of event that occurred.
	Type EventType
	// Path is the path to the file that was affected.
	Path string
	// Time is when the event occurred.
	Time time.Time
}
