# API Contracts: todotxt-tui

**Feature**: todotxt-tui  
**Date**: 2026-02-02  
**Branch**: 001-todotxt-tui

## Overview

This document defines the internal API contracts between components of the todotxt-tui application. These contracts use Go interfaces to enable testability and separation of concerns.

## Component Contracts

### Parser

Interface for parsing todo.txt format and serializing Task objects.

```go
package parser

// Parser handles parsing of todo.txt format
type Parser interface {
    // ParseLine parses a single line from todo.txt file
    // Returns Task object and any parsing error
    ParseLine(line string, lineNumber int) (*Task, error)
    
    // ParseFile parses entire todo.txt file
    // Returns slice of Tasks and any errors (non-critical errors can be collected)
    ParseFile(content string) ([]*Task, []error)
    
    // Serialize converts Task to todo.txt formatted string
    // Returns the line string
    Serialize(task *Task) string
    
    // Validate checks if a line is valid todo.txt format
    // Returns validation error if invalid
    Validate(line string) error
    
    // IsPriorityLine checks if line starts with priority (A-Z)
    IsPriorityLine(line string) bool
    
    // IsCompletedLine checks if line is marked completed
    IsCompletedLine(line string) bool
}

// Task represents a single todo item
type Task struct {
    Priority        string
    CreationDate    time.Time
    Completed       bool
    CompletionDate  time.Time
    Description     string
    Contexts        []string
    Projects        []string
    Metadata        map[string]string
    RawLine         string
    LineNumber      int
    Modified        bool
}
```

**Error Handling**:
- `ParseLine`: Returns nil Task and error for critical parsing failures
- `ParseFile`: Returns partial results with slice of non-critical errors (lines with warnings)
- `Validate`: Returns descriptive error with line number context

**Constraints**:
- Must support full todo.txt specification
- Must handle malformed lines gracefully (return Task with RawLine, mark as invalid)
- Must preserve original line formatting in RawLine field

---

### Storage

Interface for file I/O operations and persistence.

```go
package storage

// Storage handles file operations for todo.txt files
type Storage interface {
    // Load reads the todo.txt file and returns content
    // Returns content, last modified time, and any error
    Load(path string) (string, time.Time, error)
    
    // Save writes content to file atomically
    // Creates temporary file and renames on success
    Save(path string, content string) error
    
    // Watch starts watching file for external changes
    // Returns channel that receives events when file changes
    // Caller must close the returned channel when done
    Watch(path string) (<-chan FileEvent, error)
    
    // Exists checks if file exists at path
    Exists(path string) bool
    
    // Create creates new file with empty content at path
    // Creates parent directories if needed
    Create(path string) error
    
    // GetModificationTime returns file's last modified timestamp
    GetModificationTime(path string) (time.Time, error)
}

// FileEvent represents a file system event
type FileEvent struct {
    Path    string
    Type    EventType
    Modified time.Time
}

type EventType int

const (
    EventModified EventType = iota
    EventCreated
    EventDeleted
)
```

**Error Handling**:
- `Load`: Returns os.IsNotExist error if file not found
- `Save`: Returns os.PermissionError if permission denied, os.ErrNoSpace if disk full
- `Watch`: Returns error if path is directory or not readable

**Constraints**:
- Must use atomic writes (temp file + rename)
- Must use UTF-8 encoding
- Must handle file permissions gracefully
- Must support watching for external changes (fsnotify or polling)

---

### Filter

Interface for filtering and searching tasks.

```go
package filter

// Filter handles filtering and searching of tasks
type Filter interface {
    // Apply applies filter criteria to tasks
    // Returns filtered subset of tasks
    Apply(tasks []*Task, criteria FilterCriteria) []*Task
    
    // BuildIndex builds lookup indexes for fast filtering
    // Should be called after tasks are loaded or modified
    BuildIndex(tasks []*Task) Index
    
    // Search performs text search on tasks
    // Returns tasks matching search query
    Search(tasks []*Task, query string) []*Task
    
    // FilterByPriority filters tasks by priority letter
    FilterByPriority(tasks []*Task, priority string) []*Task
    
    // FilterByContext filters tasks by @context tag
    FilterByContext(tasks []*Task, context string) []*Task
    
    // FilterByProject filters tasks by +project tag
    FilterByProject(tasks []*Task, project string) []*Task
    
    // FilterByCompletion filters tasks by completion status
    FilterByCompletion(tasks []*Task, completed bool) []*Task
}

// FilterCriteria defines filter parameters
type FilterCriteria struct {
    Priorities  []string  // A-Z letters
    Contexts    []string  // @context tags
    Projects    []string  // +project tags
    Search      string    // Text search query
    Completed   *bool     // true/false/nil (both)
    Combination FilterLogic // AND or OR
}

type FilterLogic int

const (
    FilterAnd FilterLogic = iota
    FilterOr
)

// Index provides fast lookups for filtering
type Index struct {
    PriorityIndex map[string][]int
    ContextIndex  map[string][]int
    ProjectIndex  map[string][]int
    CompletionIndex map[bool][]int
}

// Builder provides fluent API for building filters
type Builder interface {
    WithPriority(priority string) Builder
    WithContext(context string) Builder
    WithProject(project string) Builder
    WithSearch(query string) Builder
    WithCompleted(completed bool) Builder
    WithLogic(logic FilterLogic) Builder
    Build() FilterCriteria
}
```

**Error Handling**:
- Filter operations should not return errors (invalid criteria are ignored)
- Empty criteria should return all tasks (no filtering)

**Constraints**:
- Must support AND and OR logic for combining filters
- Must be case-insensitive for text search
- Must support regex patterns in search queries
- Index must be rebuilt when tasks change

---

### Keymap

Interface for keyboard event handling and keybinding management.

```go
package keymap

// Keymap manages keyboard bindings and actions
type Keymap interface {
    // GetBinding returns the action for a key event in the current mode
    GetBinding(key tea.Key, mode Mode) Action
    
    // SetBinding maps a key to an action in a mode
    SetBinding(key tea.Key, action Action, mode Mode)
    
    // GetAvailableActions returns all available actions in a mode
    GetAvailableActions(mode Mode) []Action
    
    // GetKeysForAction returns all keys bound to an action in a mode
    GetKeysForAction(action Action, mode Mode) []tea.Key
    
    // IsActionAvailable checks if action has keybindings in mode
    IsActionAvailable(action Action, mode Mode) bool
}

// Mode represents application input mode
type Mode int

const (
    ModeNormal Mode = iota // Default navigation mode
    ModeInsert             // Input mode (adding/editing tasks)
    ModeDialog             // Dialog mode (prompts, filters)
    ModeSearch             // Search mode
)

// Action represents a user action
type Action int

const (
    ActionMoveDown Action = iota
    ActionMoveUp
    ActionMoveTop
    ActionMoveBottom
    ActionAddTask
    ActionEditTask
    ActionDeleteTask
    ActionCompleteTask
    ActionTogglePriority
    ActionIncreasePriority
    ActionDecreasePriority
    ActionFilter
    ActionSearch
    ActionClearFilters
    ActionReload
    ActionSave
    ActionQuit
    ActionHelp
    ActionCancel
    ActionConfirm
)

// Handler processes keyboard events
type Handler interface {
    // Handle processes a key event and returns a message
    Handle(key tea.Key, mode Mode) tea.Msg
    
    // SetMode changes the current input mode
    SetMode(mode Mode)
    
    // GetMode returns the current input mode
    GetMode() Mode
}
```

**Constraints**:
- Must support vim-style default bindings (j/k/g/G/h/l)
- Must be extensible for user customizations (future)
- Must provide help text describing all keybindings

---

### UI Models

Interfaces for TUI components using bubbletea.

```go
package ui

// Model is the base interface for all TUI components
type Model interface {
    // Init initializes the model
    Init() tea.Cmd
    
    // Update handles messages and returns updated model and command
    Update(msg tea.Msg) (Model, tea.Cmd)
    
    // View returns the string representation of the model
    View() string
}

// TaskListModel manages the task list display and navigation
type TaskListModel interface {
    Model
    
    // SetTasks sets the tasks to display
    SetTasks(tasks []*Task)
    
    // GetTasks returns all tasks
    GetTasks() []*Task
    
    // GetSelectedTask returns the currently selected task
    GetSelectedTask() *Task
    
    // SetSelectedIndex sets the selected task index
    SetSelectedIndex(index int)
    
    // GetSelectedIndex returns the selected task index
    GetSelectedIndex() int
    
    // ScrollTo scrolls to the specified task
    ScrollTo(index int)
    
    // ScrollDown scrolls down one task
    ScrollDown()
    
    // ScrollUp scrolls up one task
    ScrollUp()
    
    // ScrollToTop scrolls to the top of the task list
    ScrollToTop()
    
    // ScrollToBottom scrolls to the bottom of the task list
    ScrollToBottom()
}

// DialogModel manages dialog prompts and modals
type DialogModel interface {
    Model
    
    // SetPrompt sets the dialog prompt text
    SetPrompt(prompt string)
    
    // SetValue sets the current input value
    SetValue(value string)
    
    // GetValue returns the current input value
    GetValue() string
    
    // SetPlaceholder sets placeholder text
    SetPlaceholder(text string)
    
    // Show displays the dialog
    Show()
    
    // Hide hides the dialog
    Hide()
    
    // IsVisible returns whether dialog is visible
    IsVisible() bool
    
    // Confirm confirms the dialog action
    Confirm()
    
    // Cancel cancels the dialog action
    Cancel()
}

// FilterPanelModel manages filter selection UI
type FilterPanelModel interface {
    Model
    
    // SetAvailableFilters sets available filter options
    SetAvailableFilters(filters []Filter)
    
    // GetActiveFilters returns currently active filters
    GetActiveFilters() []Filter
    
    // ToggleFilter toggles a filter on/off
    ToggleFilter(filter Filter)
    
    // ClearFilters clears all active filters
    ClearFilters()
    
    // IsVisible returns whether filter panel is visible
    IsVisible() bool
    
    // Show displays the filter panel
    Show()
    
    // Hide hides the filter panel
    Hide()
}

// HelpModel manages help screen display
type HelpModel interface {
    Model
    
    // SetKeymap sets the keybindings to display
    SetKeymap(keymap map[Action][]string)
    
    // IsVisible returns whether help screen is visible
    IsVisible() bool
    
    // Show displays the help screen
    Show()
    
    // Hide hides the help screen
    Hide()
}

// ErrorModel manages error message display
type ErrorModel interface {
    Model
    
    // SetError sets the error message to display
    SetError(err error)
    
    // GetError returns the current error
    GetError() error
    
    // Show displays the error message
    Show()
    
    // Hide hides the error message
    Hide()
    
    // IsVisible returns whether error message is visible
    IsVisible() bool
}

// Filter represents a filter option in the UI
type Filter struct {
    Type    FilterType
    Value   string
    Label   string
    Enabled bool
}

type FilterType int

const (
    FilterPriority FilterType = iota
    FilterContext
    FilterProject
    FilterCompleted
)
```

**Constraints**:
- All models must implement bubbletea.Model interface
- Models must be immutable where possible (return new Model from Update)
- Models must support terminal resize events
- Models must maintain state for redrawing without external data

---

### App

Main application interface coordinating all components.

```go
package app

// App represents the main application
type App interface {
    // Run starts the application
    Run() error
    
    // Quit gracefully shuts down the application
    Quit()
    
    // SetConfig sets the application configuration
    SetConfig(config *Config)
    
    // GetConfig returns the current configuration
    GetConfig() *Config
    
    // GetTaskList returns the task list model
    GetTaskList() ui.TaskListModel
    
    // GetParser returns the parser instance
    GetParser() parser.Parser
    
    // GetStorage returns the storage instance
    GetStorage() storage.Storage
    
    // GetFilter returns the filter instance
    GetFilter() filter.Filter
    
    // GetKeymap returns the keymap instance
    GetKeymap() keymap.Keymap
    
    // ReloadTasks reloads tasks from disk
    ReloadTasks() error
    
    // SaveTasks saves tasks to disk
    SaveTasks() error
    
    // ShowError displays an error message
    ShowError(err error)
}

// Config represents application configuration
type Config struct {
    TodoFilePath      string
    DoneFilePath      string
    Theme             string
    ShowCompleted     bool
    ArchiveCompleted  bool
    ConfirmDelete     bool
    AutoSave          bool
    FileWatchEnabled  bool
}
```

---

## Event Flow

### Task Update Flow

```
User Input (key event)
    ↓
KeymapHandler.Handle()
    ↓
Generates tea.Msg (e.g., EditTaskMsg)
    ↓
App.Update(msg)
    ↓
Task.Edit() → Modified=true
    ↓
TaskList.UpdateTask()
    ↓
TodoFile.Modified=true
    ↓
AutoSave: Storage.Save()
    ↓
TaskList visible refresh
```

### Filter Update Flow

```
User Input (e.g., / key)
    ↓
KeymapHandler.Handle()
    ↓
Opens DialogModel (search prompt)
    ↓
User enters search text
    ↓
DialogModel.Confirm()
    ↓
Generates SearchMsg with query
    ↓
Filter.Search() → returns filtered tasks
    ↓
TaskList.SetTasks(filteredTasks)
    ↓
UI redraw with filtered view
```

## Message Types

Custom bubbletea messages for component communication.

```go
package ui

// TaskUpdatedMsg signals a task was modified
type TaskUpdatedMsg struct {
    Task *Task
    Index int
}

// TaskDeletedMsg signals a task was deleted
type TaskDeletedMsg struct {
    Index int
}

// TaskCompletedMsg signals a task completion status changed
type TaskCompletedMsg struct {
    Task     *Task
    Completed bool
}

// FilterChangedMsg signals filters were modified
type FilterChangedMsg struct {
    Filters []filter.Filter
}

// ReloadRequestedMsg signals user requested file reload
type ReloadRequestedMsg struct{}

// SaveRequestedMsg signals user requested manual save
type SaveRequestedMsg struct{}

// ErrorMsg signals an error occurred
type ErrorMsg struct {
    Err error
}

// FileChangedMsg signals external file modification
type FileChangedMsg struct {
    Path     string
    Modified time.Time
}

// HelpRequestedMsg signals user requested help screen
type HelpRequestedMsg struct{}

// HelpClosedMsg signals help screen was closed
type HelpClosedMsg struct{}
```

## Testing Contracts

All interfaces must be mockable for testing:

```go
// Example: MockParser for testing
type MockParser struct {
    ParseFunc    func(line string, lineNumber int) (*parser.Task, error)
    ParseFileFunc func(content string) ([]*parser.Task, []error)
    SerializeFunc func(task *parser.Task) string
}

func (m *MockParser) ParseLine(line string, lineNumber int) (*parser.Task, error) {
    if m.ParseFunc != nil {
        return m.ParseFunc(line, lineNumber)
    }
    return &parser.Task{}, nil
}

// ... other methods
```

Use `gomock` (`github.com/uber-go/mock`) for generating mocks automatically.

---

## Versioning

API contracts may evolve during implementation. Changes must:

1. Maintain backward compatibility where possible
2. Update this document with version tags
3. Document breaking changes
4. Update tests accordingly

Contract version: 1.0.0