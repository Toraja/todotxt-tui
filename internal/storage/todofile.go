package storage

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/Toraja/todotxt-tui/internal/parser"
)

// TodoFile represents a todo.txt file with all its tasks and metadata.
type TodoFile struct {
	// File metadata
	Path     string // Absolute path to todo.txt file
	Encoding string // File encoding (default: UTF-8)

	// Task data
	Tasks        []*parser.Task       // All tasks in file order
	CompletedIdx map[int]*parser.Task // Map of line number → completed task (for done.txt)

	// State tracking
	LastModified time.Time // Last known modification time
	Modified     bool      // True if unsaved changes exist
	Loaded       bool      // True if file has been loaded
	LoadError    error     // Last load error (nil if successful)

	// Dependencies
	storage Storage       // Storage implementation for file I/O
	parser  parser.Parser // Parser for todo.txt format
}

// NewTodoFile creates a new TodoFile instance.
func NewTodoFile(path string, storage Storage, p parser.Parser) *TodoFile {
	return &TodoFile{
		Path:         path,
		Encoding:     "UTF-8",
		Tasks:        []*parser.Task{},
		CompletedIdx: make(map[int]*parser.Task),
		storage:      storage,
		parser:       p,
	}
}

// Load reads and parses the todo.txt file from disk.
func (tf *TodoFile) Load() error {
	// Check if file exists
	exists, err := tf.storage.Exists(tf.Path)
	if err != nil {
		tf.LoadError = fmt.Errorf("failed to check file existence: %w", err)
		return tf.LoadError
	}

	// If file doesn't exist, create an empty file
	if !exists {
		if err := tf.storage.Create(tf.Path); err != nil {
			tf.LoadError = fmt.Errorf("failed to create file: %w", err)
			return tf.LoadError
		}
		tf.Tasks = []*parser.Task{}
		tf.Loaded = true
		tf.Modified = false
		tf.LoadError = nil
		tf.LastModified = time.Now()
		return nil
	}

	// Load file content
	content, err := tf.storage.Load(tf.Path)
	if err != nil {
		tf.LoadError = fmt.Errorf("failed to load file: %w", err)
		return tf.LoadError
	}

	// Parse file content
	tasks, err := tf.parser.ParseFile(content)
	if err != nil {
		tf.LoadError = fmt.Errorf("failed to parse file: %w", err)
		return tf.LoadError
	}

	// Update state
	tf.Tasks = tasks
	tf.Loaded = true
	tf.Modified = false
	tf.LoadError = nil

	// Get modification time
	modTime, err := tf.storage.GetModificationTime(tf.Path)
	if err != nil {
		// Non-fatal error, just use current time
		tf.LastModified = time.Now()
	} else {
		tf.LastModified = modTime
	}

	return nil
}

// Save writes all tasks to disk atomically.
func (tf *TodoFile) Save() error {
	if !tf.Loaded {
		return fmt.Errorf("cannot save: file not loaded")
	}

	// Build file content from tasks
	lines := make([]string, 0, len(tf.Tasks))
	for _, task := range tf.Tasks {
		lines = append(lines, tf.parser.Serialize(task))
	}

	// Convert to io.Reader
	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		content += "\n" // Add trailing newline
	}
	reader := bytes.NewReader([]byte(content))

	// Save to disk
	if err := tf.storage.Save(tf.Path, reader); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	// Update state
	tf.Modified = false

	// Update modification time
	modTime, err := tf.storage.GetModificationTime(tf.Path)
	if err != nil {
		// Non-fatal error, just use current time
		tf.LastModified = time.Now()
	} else {
		tf.LastModified = modTime
	}

	// Clear modified flags on all tasks
	for _, task := range tf.Tasks {
		task.Modified = false
	}

	return nil
}

// AddTask appends a new task to the list.
func (tf *TodoFile) AddTask(task *parser.Task) {
	// Set line number
	task.LineNumber = len(tf.Tasks) + 1

	// Append to task list
	tf.Tasks = append(tf.Tasks, task)

	// Mark as modified
	tf.Modified = true
	task.Modified = true
}

// UpdateTask replaces the task at the given index.
func (tf *TodoFile) UpdateTask(index int, task *parser.Task) error {
	if index < 0 || index >= len(tf.Tasks) {
		return fmt.Errorf("index out of range: %d", index)
	}

	// Preserve line number
	task.LineNumber = tf.Tasks[index].LineNumber

	// Update task
	tf.Tasks[index] = task

	// Mark as modified
	tf.Modified = true
	task.Modified = true

	return nil
}

// DeleteTask removes the task at the given index.
func (tf *TodoFile) DeleteTask(index int) error {
	if index < 0 || index >= len(tf.Tasks) {
		return fmt.Errorf("index out of range: %d", index)
	}

	// Remove task
	tf.Tasks = append(tf.Tasks[:index], tf.Tasks[index+1:]...)

	// Renumber tasks
	for i := index; i < len(tf.Tasks); i++ {
		tf.Tasks[i].LineNumber = i + 1
	}

	// Mark as modified
	tf.Modified = true

	return nil
}

// GetTask returns the task at the given index.
func (tf *TodoFile) GetTask(index int) (*parser.Task, error) {
	if index < 0 || index >= len(tf.Tasks) {
		return nil, fmt.Errorf("index out of range: %d", index)
	}
	return tf.Tasks[index], nil
}

// Count returns the total number of tasks.
func (tf *TodoFile) Count() int {
	return len(tf.Tasks)
}

// CountCompleted returns the number of completed tasks.
func (tf *TodoFile) CountCompleted() int {
	count := 0
	for _, task := range tf.Tasks {
		if task.IsComplete() {
			count++
		}
	}
	return count
}

// CountActive returns the number of active (non-completed) tasks.
func (tf *TodoFile) CountActive() int {
	return tf.Count() - tf.CountCompleted()
}

// ReloadIfChanged checks if the file has been modified externally.
// If changed, reloads from disk and returns true, else returns false.
func (tf *TodoFile) ReloadIfChanged() (bool, error) {
	// Check if file exists
	exists, err := tf.storage.Exists(tf.Path)
	if err != nil {
		return false, fmt.Errorf("failed to check file existence: %w", err)
	}

	if !exists {
		return false, fmt.Errorf("file no longer exists: %s", tf.Path)
	}

	// Get current modification time
	modTime, err := tf.storage.GetModificationTime(tf.Path)
	if err != nil {
		return false, fmt.Errorf("failed to get modification time: %w", err)
	}

	// Check if modified
	if modTime.After(tf.LastModified) {
		// File has been modified externally, reload
		if err := tf.Load(); err != nil {
			return true, fmt.Errorf("failed to reload file: %w", err)
		}
		return true, nil
	}

	return false, nil
}

// HasUnsavedChanges returns true if there are unsaved changes.
func (tf *TodoFile) HasUnsavedChanges() bool {
	return tf.Modified
}

// IsLoaded returns true if the file has been loaded.
func (tf *TodoFile) IsLoaded() bool {
	return tf.Loaded
}

// GetPath returns the file path.
func (tf *TodoFile) GetPath() string {
	return tf.Path
}

// GetLastError returns the last load error.
func (tf *TodoFile) GetLastError() error {
	return tf.LoadError
}
