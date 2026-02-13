package parser

import (
	"fmt"
	"strings"
	"time"
)

// Task represents a single todo item with full todo.txt format support.
type Task struct {
	// Core fields (from todo.txt spec)
	Priority       string              // A-Z uppercase letter, empty if no priority
	CreationDate   time.Time           // Creation date in YYYY-MM-DD format, zero if not specified
	Completed      bool                // True if task is completed
	CompletionDate time.Time           // Completion date in YYYY-MM-DD format, zero if not completed
	Description    string              // Task description text
	Contexts       map[string]struct{} // Set of @context tags for O(1) lookup
	Projects       map[string]struct{} // Set of +project tags for O(1) lookup
	Metadata       map[string]string   // Additional key:value pairs from description

	// Bookkeeping fields
	RawLine    string // Original line from file (for preservation)
	LineNumber int    // Line number in source file (1-indexed)
	Modified   bool   // True if task has unsaved changes
}

// IsComplete returns true if the task is marked as completed.
func (t *Task) IsComplete() bool {
	return t.Completed
}

// SetPriority updates the task priority.
// Priority must be a single uppercase letter A-Z or empty string.
// Returns an error if the priority is invalid.
func (t *Task) SetPriority(priority string) error {
	if priority == "" {
		t.Priority = ""
		t.Modified = true
		return nil
	}
	if len(priority) != 1 {
		return fmt.Errorf("priority must be single letter A-Z")
	}
	if priority[0] < 'A' || priority[0] > 'Z' {
		return fmt.Errorf("priority must be letter A-Z")
	}
	t.Priority = priority
	t.Modified = true
	return nil
}

// AddContext adds a @context tag to the task if not already present.
func (t *Task) AddContext(context string) error {
	if !strings.HasPrefix(context, "@") {
		return fmt.Errorf("context must start with @")
	}
	if t.Contexts == nil {
		t.Contexts = make(map[string]struct{})
	}
	// Check if already present (O(1) lookup)
	if _, exists := t.Contexts[context]; exists {
		return nil // Already exists, no-op
	}
	t.Contexts[context] = struct{}{}
	t.Modified = true
	return nil
}

// RemoveContext removes a @context tag from the task.
func (t *Task) RemoveContext(context string) {
	if t.Contexts == nil {
		return
	}
	if _, exists := t.Contexts[context]; exists {
		delete(t.Contexts, context)
		t.Modified = true
	}
}

// AddProject adds a +project tag to the task if not already present.
func (t *Task) AddProject(project string) error {
	if !strings.HasPrefix(project, "+") {
		return fmt.Errorf("project must start with +")
	}
	if t.Projects == nil {
		t.Projects = make(map[string]struct{})
	}
	// Check if already present (O(1) lookup)
	if _, exists := t.Projects[project]; exists {
		return nil // Already exists, no-op
	}
	t.Projects[project] = struct{}{}
	t.Modified = true
	return nil
}

// RemoveProject removes a +project tag from the task.
func (t *Task) RemoveProject(project string) {
	if t.Projects == nil {
		return
	}
	if _, exists := t.Projects[project]; exists {
		delete(t.Projects, project)
		t.Modified = true
	}
}

// HasContext checks if the task has a specific @context tag.
func (t *Task) HasContext(context string) bool {
	if t.Contexts == nil {
		return false
	}
	_, exists := t.Contexts[context]
	return exists
}

// HasProject checks if the task has a specific +project tag.
func (t *Task) HasProject(project string) bool {
	if t.Projects == nil {
		return false
	}
	_, exists := t.Projects[project]
	return exists
}

// Complete marks the task as completed with the given completion date.
// If the completion date is zero, it uses the current date.
func (t *Task) Complete(completionDate time.Time) {
	if completionDate.IsZero() {
		completionDate = time.Now()
	}
	t.Completed = true
	t.CompletionDate = completionDate
	t.Modified = true
}

// Uncomplete marks the task as incomplete (toggles back from completed).
func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletionDate = time.Time{}
	t.Modified = true
}

// String generates the todo.txt formatted line for this task.
// Format: [x] [(P)] [completion date] [creation date] description [contexts] [projects] [metadata]
func (t *Task) String() string {
	var parts []string

	// Completed marker
	if t.Completed {
		parts = append(parts, "x")
	}

	// Priority (only if not completed)
	if !t.Completed && t.Priority != "" {
		parts = append(parts, fmt.Sprintf("(%s)", t.Priority))
	}

	// Completion date (only if completed)
	if t.Completed && !t.CompletionDate.IsZero() {
		parts = append(parts, t.CompletionDate.Format("2006-01-02"))
	}

	// Creation date
	if !t.CreationDate.IsZero() {
		parts = append(parts, t.CreationDate.Format("2006-01-02"))
	}

	// Description with contexts and projects embedded
	parts = append(parts, t.Description)

	return strings.Join(parts, " ")
}
