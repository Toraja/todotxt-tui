# Data Model: todotxt-tui

**Feature**: todotxt-tui  
**Date**: 2026-02-02  
**Branch**: 001-todotxt-tui

## Overview

This document defines the core data entities for the todotxt-tui application. All entities are designed to support the todo.txt format specification while enabling efficient TUI operations and filtering.

## Core Entities

### Task

Represents a single todo item with full todo.txt format support.

```go
type Task struct {
    // Core fields (from todo.txt spec)
    Priority        string        // A-Z uppercase letter, empty if no priority
    CreationDate    time.Time     // Creation date in YYYY-MM-DD format, zero if not specified
    Completed       bool          // True if task is completed
    CompletionDate  time.Time     // Completion date in YYYY-MM-DD format, zero if not completed
    Description     string        // Task description text
    Contexts        []string      // List of @context tags
    Projects        []string      // List of +project tags
    Metadata        map[string]string // Additional key:value pairs from description
    
    // Bookkeeping fields
    RawLine         string        // Original line from file (for preservation)
    LineNumber      int           // Line number in source file (1-indexed)
    Modified        bool          // True if task has unsaved changes
}
```

**Design Note**: `Task.Modified` tracks individual task changes for UI indicators, while `TodoFile.Modified` tracks file-level changes to determine if save is needed. Tasks are always saved atomically as a complete file, not individually.

**Validation Rules**:
- Priority: Single uppercase letter A-Z, or empty string
- CreationDate: Must be valid YYYY-MM-DD format or zero time
- CompletionDate: Must be valid YYYY-MM-DD format if Completed=true, otherwise zero
- Description: Non-empty string, can contain any characters
- Contexts: Each must be single word starting with @, unique within task
- Projects: Each must be single word starting with +, unique within task
- Metadata: Keys are alphanumeric + underscore, values can be any string

**State Transitions**:
- Active → Completed: Set Completed=true, CompletionDate=today
- Completed → Active: Set Completed=false, CompletionDate=zero, move to bottom of list
- Any → Deleted: Remove from TaskList and mark for deletion from file

**Methods**:
- `String()`: Generate todo.txt formatted line
- `IsComplete()`: Return true if task is completed
- `SetPriority(priority string)`: Update priority (A-Z or empty)
- `AddContext(context string)`: Add @context tag if not present
- `RemoveContext(context string)`: Remove @context tag
- `AddProject(project string)`: Add +project tag if not present
- `RemoveProject(project string)`: Remove +project tag
- `HasContext(context string) bool`: Check if context exists
- `HasProject(project string) bool`: Check if project exists
- `MatchesFilter(filter Filter) bool`: Check if task matches filter criteria

### TodoFile

Represents the todo.txt file containing all tasks with persistence logic.

```go
type TodoFile struct {
    // File metadata
    Path         string        // Absolute path to todo.txt file
    Encoding     string        // File encoding (default: UTF-8)
    
    // Task data
    Tasks        []*Task       // All tasks in file order
    CompletedIdx map[int]*Task // Map of line number → completed task (for done.txt)
    
    // State tracking
    LastModified time.Time     // Last known modification time
    Modified     bool          // True if unsaved changes exist
    Loaded       bool          // True if file has been loaded
    LoadError    error         // Last load error (nil if successful)
}
```

**Design Note**: `LoadError` enables graceful error handling (e.g., file doesn't exist yet, permission issues, parse errors on specific lines). The TUI displays errors to the user and allows recovery actions (retry, choose different file) rather than exiting abruptly.

**Validation Rules**:
- Path: Must be valid absolute path with read/write permissions
- Encoding: Must be UTF-8 (standard for todo.txt)
- Tasks: Can be empty (no tasks in file)

**State Transitions**:
- Unloaded → Loaded: Read file from disk, parse all lines
- Loaded → Modified: Add/update/delete tasks, set Modified=true
- Modified → Saved: Write all tasks to disk atomically, set Modified=false
- External Change Detected: Reload from disk, merge or prompt user

**Methods**:
- `Load() error`: Read and parse file from disk
- `Save() error`: Write all tasks to disk atomically
- `AddTask(task *Task)`: Append task to list
- `UpdateTask(index int, task *Task)`: Replace task at index
- `DeleteTask(index int)`: Remove task from list
- `GetTask(index int) *Task`: Get task by index
- `Count() int`: Return total task count
- `CountCompleted() int`: Return count of completed tasks
- `CountActive() int`: Return count of active tasks
- `ReloadIfChanged() bool`: Check file modification time; if changed, reload from disk and return true, else return false

### Filter

Represents a filter criteria for filtering tasks.

```go
type FilterType int

const (
    FilterNone FilterType = iota
    FilterPriority
    FilterContext
    FilterProject
    FilterSearch
    FilterCompleted
)

type Filter struct {
    Type        FilterType    // Type of filter
    Value       string        // Filter value (priority letter, @context, +project, search text)
    Enabled     bool          // True if filter is active
    Combination FilterLogic   // AND/OR for combining with other filters
}

type FilterLogic int

const (
    FilterAnd FilterLogic = iota // All filters must match
    FilterOr                      // Any filter can match
)
```

**Validation Rules**:
- Type: Must be valid FilterType constant
- Value: Non-empty string when Enabled=true
- Priority: Single uppercase letter A-Z
- Context: Single word starting with @
- Project: Single word starting with +
- Search: Any non-empty string (regex supported)

**Methods**:
- `Matches(task *Task) bool`: Check if task matches filter
- `String()`: Return human-readable filter description
- `Equals(other Filter) bool`: Check if two filters are equivalent

### TaskList

Represents the filtered and sorted view of tasks displayed to users.

```go
type TaskList struct {
    // Source data
    AllTasks    []*Task       // All tasks from TodoFile (immutable reference)
    
    // Filter state
    ActiveFilters []Filter    // Currently active filters
    FilterLogic   FilterLogic // How to combine filters (AND/OR)
    
    // Sort state
    SortBy       SortCriteria // Current sort field
    SortOrder    SortOrder    // Ascending or descending
    
    // View state
    VisibleTasks  []*Task     // Filtered and sorted subset of tasks
    SelectedIndex int         // Index of currently selected task in VisibleTasks
    ScrollOffset  int         // Scroll offset for viewport
    ViewportSize  int         // Number of tasks visible in viewport
    
    // Computed
    TotalCount   int          // Total tasks in AllTasks
    VisibleCount int          // Tasks in VisibleTasks
}
```

**Design Note**: `TodoFile.Tasks` is the persistence layer (owns data, handles file I/O), while `TaskList.AllTasks` is the presentation layer (references same tasks for filtering/sorting). TaskList should not duplicate the slice; it should reference `TodoFile.Tasks` directly. Task modifications are made through TodoFile methods, then TaskList is refreshed.

**SortCriteria Enum**:
```go
type SortCriteria int

const (
    SortNone SortCriteria = iota
    SortPriority
    SortCreationDate
    SortCompletionDate
    SortDescription
)
```

**SortOrder Enum**:
```go
type SortOrder int

const (
    SortAscending SortOrder = iota
    SortDescending
)
```

**Validation Rules**:
- SelectedIndex: Must be in range [0, VisibleCount-1] or -1 (no selection)
- ScrollOffset: Must be >= 0
- ViewportSize: Must be > 0 (set from terminal height)

**State Transitions**:
- Filter Changed: Recompute VisibleTasks based on ActiveFilters
- Sort Changed: Sort VisibleTasks by SortBy and SortOrder
- Task Modified: Update Task in AllTasks, refresh VisibleTasks if affected

**Methods**:
- `ApplyFilters()`: Recompute VisibleTasks based on ActiveFilters
- `AddFilter(filter Filter)`: Add filter to ActiveFilters and recompute
- `RemoveFilter(filterType FilterType)`: Remove filter by type and recompute
- `ClearFilters()`: Clear all ActiveFilters, show all tasks
- `SetSort(by SortCriteria, order SortOrder)`: Sort VisibleTasks
- `SelectNext()`: Move selection down (cyclic)
- `SelectPrev()`: Move selection up (cyclic)
- `SelectFirst()`: Move selection to first task
- `SelectLast()`: Move selection to last task
- `GetSelectedTask() *Task`: Return currently selected task or nil
- `ScrollDown()`: Move scroll offset down
- `ScrollUp()`: Move scroll offset up
- `GetVisibleTasks() []*Task`: Return tasks in current viewport
- `GetTaskAtViewportPosition(pos int) *Task`: Get task at viewport position

### Configuration

Application configuration with defaults and user overrides.

```go
type Config struct {
    // File paths
    TodoFilePath  string        // Path to todo.txt file
    DoneFilePath  string        // Optional path to done.txt for archiving
    ConfigPath    string        // Path to config file
    
    // Display settings
    Theme         string        // Color theme (light, dark, default)
    ShowCompleted bool          // Show completed tasks in main list
    ArchiveCompleted bool        // Auto-archive completed tasks to done.txt
    
    // UI settings
    ConfirmDelete bool          // Require confirmation before deleting
    AutoSave      bool          // Automatically save after modifications
    FileWatchEnabled bool        // Watch file for external changes
    
    // Keyboard (future: customizable)
    // Keymap map[string]string // Custom keybindings
}
```

**Validation Rules**:
- TodoFilePath: Must be valid absolute path, create if doesn't exist
- DoneFilePath: Optional, valid absolute path if specified
- Theme: One of "light", "dark", "default"
- AutoSave: Must be true (auto-save is non-negotiable for data safety)

**Methods**:
- `Load() error`: Load config from file or use defaults
- `Defaults()`: Return default configuration
- `Validate() error`: Validate all configuration values

## Relationships

```
TodoFile
    └── Tasks (1:n) → Task

TaskList
    ├── AllTasks (1:n) → Task
    ├── ActiveFilters (0:n) → Filter
    └── VisibleTasks (0:n) → Task (filtered subset)

Filter
    └── FilterLogic (enum) → AND/OR

Task
    ├── Contexts (0:n) → string
    ├── Projects (0:n) → string
    └── Metadata (0:n) → key:value pairs

Configuration
    ├── TodoFilePath → TodoFile
    └── DoneFilePath → done.txt
```

## Data Flow

1. **Application Startup**:
   - Load Configuration from file
   - Load TodoFile from TodoFilePath
   - Parse all lines into Task objects
   - Initialize TaskList with all tasks
   - Apply default filters (if configured)

2. **User Navigation**:
   - TaskList maintains SelectedIndex
   - UI renders tasks in viewport based on ScrollOffset
   - Key events update SelectedIndex and ScrollOffset

3. **Filter Operations**:
   - User adds filter → AddFilter() → ApplyFilters()
   - TaskList recomputes VisibleTasks
   - SelectedIndex adjusted if out of range
   - UI re-rendered with filtered view

4. **Task Modification**:
   - User edits task → UpdateTask() in TodoFile
   - Task.Modified = true
   - TodoFile.Modified = true
   - AutoSave: TodoFile.Save()
   - TaskList refreshed if task in VisibleTasks

5. **File Persistence**:
   - TodoFile.Save() writes all tasks to temporary file
   - Atomic rename to original path
   - TodoFile.Modified = false
   - Update LastModified timestamp

## Performance Considerations

### Indexing

For efficient filtering, maintain indexes in TaskList:

```go
type TaskList struct {
    // ... existing fields ...
    
    // Indexes for fast lookups
    PriorityIndex  map[string][]int // priority letter → task indices
    ContextIndex   map[string][]int // @context → task indices
    ProjectIndex   map[string][]int // +project → task indices
}
```

- Build indexes on load/reload: O(n)
- Filter by priority/context/project: O(k) where k = matching tasks
- Rebuild indexes only when tasks change (add/delete/modify)

### Caching

Cache computed values to avoid repeated computation:

```go
type Task struct {
    // ... existing fields ...
    
    // Cached display string
    displayString string // Cached formatted string for rendering
}
```

- Generate display string on load and task modification
- Reuse when rendering (no re-computation needed)
- Invalidate cache on task update

### Lazy Loading

For very large files (10,000+ tasks):

```go
type TodoFile struct {
    // ... existing fields ...
    
    // Lazy loading state
    FullyLoaded   bool          // True if all tasks loaded
    LoadedCount   int           // Number of tasks currently loaded
}
```

- Load first 1000 tasks immediately
- Load remaining tasks in background goroutines
- Show loading indicator while loading
- Update UI incrementally as tasks load

## Validation Examples

### Priority Validation

```go
func (t *Task) SetPriority(priority string) error {
    if priority == "" {
        t.Priority = ""
        return nil
    }
    if len(priority) != 1 {
        return errors.New("priority must be single letter A-Z")
    }
    if priority[0] < 'A' || priority[0] > 'Z' {
        return errors.New("priority must be letter A-Z")
    }
    t.Priority = priority
    return nil
}
```

### Date Validation

```go
func ParseTodoDate(dateStr string) (time.Time, error) {
    if dateStr == "" {
        return time.Time{}, nil
    }
    t, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
    }
    return t, nil
}
```

### Context/Project Validation

```go
func ValidateContextTag(tag string) error {
    if !strings.HasPrefix(tag, "@") {
        return errors.New("context must start with @")
    }
    if len(tag) < 2 {
        return errors.New("context tag too short")
    }
    rest := tag[1:]
    if !isValidTagWord(rest) {
        return errors.New("context must be single word (no spaces)")
    }
    return nil
}

func isValidTagWord(s string) bool {
    // Allow alphanumeric, underscore, hyphen
    for _, r := range s {
        if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
            return false
        }
    }
    return true
}
```

## Migration Notes

This data model is designed for the initial implementation. Future enhancements may require:

- Add support for recurring tasks
- Add task dependencies or blocking
- Add task history or undo stack (beyond single-level delete undo)
- Add subtasks or hierarchical organization
- Add custom metadata fields beyond key:value pairs

All enhancements should maintain backward compatibility with todo.txt format.