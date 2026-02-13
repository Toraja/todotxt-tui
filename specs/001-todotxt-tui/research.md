# Research: todotxt-tui Implementation

**Feature**: todotxt-tui  
**Date**: 2026-02-02  
**Branch**: 001-todotxt-tui

## Summary

Research findings for implementing todotxt-tui, a terminal UI application for managing todo.txt files. This document validates technology choices and establishes best practices for TUI development with large datasets.

## Technology Choices

### TUI Framework: bubbletea

**Decision**: Use bubbletea as the primary TUI framework

**Rationale**:
- Pure Go implementation with no CGO dependencies, ensuring cross-platform compatibility
- Model-View-Update (MVU) architecture provides clear separation of concerns
- Excellent performance characteristics for 60fps rendering
- Strong community support and active maintenance (Charmbracelet ecosystem)
- Built-in support for keyboard events, terminal resizing, and mouse input
- Integrates seamlessly with lipgloss for styling

**Alternatives Considered**:
- termui: Less actively maintained, more imperative style
- tcell: Lower-level, requires more boilerplate
- gocui: Less mature, fewer features

### Styling Library: lipgloss

**Decision**: Use lipgloss for styling and theming

**Rationale**:
- Designed specifically for bubbletea applications
- Declarative API for defining styles (colors, borders, padding)
- Built-in color palette with sensible defaults
- Supports complex layouts with easy-to-use primitives

**Alternatives Considered**:
- termenv: Lower-level, lipgloss provides better abstraction
- Custom ANSI escape sequences: Too much boilerplate

### Testing Framework: Ginkgo + Gomega

**Decision**: Use Ginkgo BDD framework with Gomega matchers

**Rationale**:
- Required by constitution
- BDD style produces readable, maintainable tests
- Excellent table-driven test support with DescribeTable
- Strong async/timeout support for testing TUI interactions
- Good integration with Go's testing package

**Note**: Already mandated by constitution, no alternative needed

## Implementation Patterns

### Progressive Rendering for Large Task Lists

**Decision**: Implement viewport-based rendering with virtual list

**Rationale**:
- Render only visible tasks (typically 20-30) instead of entire list
- Maintain index of all tasks for O(1) lookup by position
- Use bubbletea's viewport component for scroll optimization
- Cache rendered task strings to avoid repeated string formatting
- Lazy-load tasks from file on demand (streaming for very large files)

**Performance Targets**:
- Initial load: Parse all tasks to index (<30ms for 1000 tasks)
- View update: Render visible region only (<5ms)
- Scroll: Update viewport position and re-render (<5ms)

**Best Practices**:
- Precompute task display strings on load/update
- Use efficient data structures: slice for indexed access, map for filtered views
- Debounce search/filter operations to avoid blocking UI
- Background goroutines for expensive operations (file I/O)

### todo.txt Parsing Strategy

**Decision**: Implement streaming parser with regex-based validation

**Rationale**:
- todo.txt format is line-based, easy to parse incrementally
- Regex patterns for extracting priority, date, contexts, projects
- Graceful degradation for malformed lines (display as-is)
- Full specification compliance (https://github.com/todotxt/todo.txt)

**Parser Requirements**:
```go
// Task structure based on spec
type Task struct {
    Priority      string              // A-Z or empty
    CreationDate  time.Time           // YYYY-MM-DD or zero
    Completed     bool
    CompletionDate time.Time          // YYYY-MM-DD if completed
    Description   string
    Contexts      map[string]struct{} // @tags
    Projects      map[string]struct{} // +tags
    Metadata      map[string]string   // key:value pairs
    RawLine       string              // Original line for preservation
    LineNumber    int                 // For error reporting
}
```

**Parsing Pattern**:
```
^(x )?(\([A-Z]\) )?(\d{4}-\d{2}-\d{2} )?(.*)$

Priority: \([A-Z]\) at start (if not completed)
Creation date: YYYY-MM-DD after priority or start of line
Completed: "x " prefix
Completion date: after "x " prefix
Contexts: @word pattern
Projects: +word pattern
Metadata: key:value pairs in description
```

### Vim-style Keybinding Implementation

**Decision**: Implement central keymap registry with vim-style modes

**Rationale**:
- Centralized key mapping for easy customization
- Support multiple distinct modes: normal mode (navigation), command mode (text input/commands), and dialog mode (confirmation prompts)
- Clear separation of intent (what) from action (how)
- Easy to add keybindings in future

**Keymap Structure**:
```go
type Keymap struct {
    NormalMode    map[key.Binding]Msg
    CommandMode   map[key.Binding]Msg
    DialogMode    map[key.Binding]Msg
}
```

**Vim-style Mappings**:
- `j`/`k`: Move down/up (normal mode)
- `g`/`G`: Jump to top/bottom (normal mode)
- `h`/`l`: Navigate left/right in dialogs (dialog mode)
- `a`: Add task (normal mode)
- `e`: Edit task (normal mode)
- `Space`: Toggle completion (normal mode)
- `d`: Delete task (normal mode)
- `Enter`: Confirm action (all modes)
- `Esc`: Cancel/back (all modes)
- `q`: Quit application (all modes)

**Best Practices**:
- Use bubbletea's key.Binding for type safety
- Provide visual hints for available keys in help text
- Support key sequences (e.g., `gg` for top) if needed
- Configurable keymap (future enhancement)

### File I/O and Persistence

**Decision**: Implement atomic writes with file watching

**Rationale**:
- Atomic writes prevent data corruption on crash/interruption
- File watching detects external changes
- Graceful handling of permission errors and disk space issues
- Auto-save after every modification with manual save option

**Persistence Strategy**:
1. Load: Read entire file, parse all lines, build index
2. Save: Write to temporary file, rename on success (atomic)
3. Watch: Use fsnotify or poll for external modifications
4. Reload: Prompt user if file changed externally

**Error Handling**:
- File not found on load: Create empty file with message
- Permission denied: Display error, prevent modifications
- Disk full: Display error, queue changes for retry
- Corrupt file: Parse what's possible, display warnings

### Filtering and Search Implementation

**Decision**: Implement filter pipeline with bitmask indexing

**Rationale**:
- Efficient multi-criteria filtering (AND/OR combinations)
- Maintain separate indexes for fast lookups
- Support live search with debouncing
- Clear visual indication of active filters

**Filter Pipeline**:
1. Parse user input (regex for search, specific values for filters)
2. Build filter criteria object
3. Apply filters sequentially or in parallel
4. Return filtered view (index subset)

**Indexing Strategy**:
- Priority index: map[priority][]int (task indices)
- Context index: map[context][]int
- Project index: map[project][]int
- Text index: trigram or inverted index for search (optional for <10k tasks)

**Performance**:
- Build indexes on load: O(n)
- Filter by priority: O(k) where k = tasks with that priority
- Search by text: O(n*m) where n = tasks, m = average task length
- Combined filters: Apply most selective first

## Architecture Decisions

### Package Organization

**Decision**: Domain-driven package structure

**Rationale**:
- Packages organized by business domain (parser, filter, storage)
- Clear boundaries between components
- Easy to test in isolation
- Follows Go community conventions

**Package Breakdown**:
- `internal/config`: Configuration file parsing and defaults
- `internal/parser`: todo.txt parsing and Task struct
- `internal/storage`: File I/O, atomic writes, watching
- `internal/filter`: Filtering and search logic
- `internal/keymap`: Keyboard bindings and handling
- `internal/ui`: TUI models, components, views
- `tests/integration`: End-to-end tests

### Error Handling Strategy

**Decision**: Explicit error propagation with user-friendly messages

**Rationale**:
- Go convention: never ignore errors
- Clear distinction between internal errors (logs) and user errors (display)
- Graceful degradation where possible
- Contextual error messages with suggested fixes

**Error Types**:
- File errors: permission denied, not found, disk full, corrupt
- Parse errors: malformed task lines, invalid dates
- UI errors: terminal too small, color not supported
- User errors: invalid input, invalid filter criteria

**Display Strategy**:
- Non-blocking error toasts for minor issues
- Modal dialogs for critical errors
- Help text with suggestions
- Log internal errors to stderr

### Performance Optimization Strategy

**Decision**: Profile-driven optimization with progressive enhancement

**Rationale**:
- Optimize based on actual bottlenecks (pprof)
- Start simple, optimize when needed
- Maintain code readability
- Target specific performance goals

**Optimization Priorities**:
1. Startup time: Efficient parsing, lazy loading
2. Rendering: Viewport-based, cached strings
3. Filtering: Indexing, selective application
4. File I/O: Buffering, async operations
5. Memory: Reuse buffers, avoid allocations in hot paths

**Profiling Tools**:
- `go tool pprof` for CPU and memory profiling
- Benchmark tests for critical paths
- Flame graphs for visualization

## Unresolved Questions

### Configuration Format

**Status**: Needs decision in Phase 1

**Options**:
- YAML (easy to read, common for config)
- JSON (standard, built-in support)
- TOML (simple, good for flat config)
- Environment variables (no config file needed)

**Recommendation**: YAML with environment variable overrides for Docker/container usage

### Theme System

**Status**: Needs decision in Phase 1

**Options**:
- Predefined color schemes (light, dark, terminal default)
- User-configurable color palette
- Automatic theme detection based on terminal background

**Recommendation**: Start with predefined schemes, add user configuration in future

### Archive Support

**Status**: Clearly specified in spec, implement done.txt archiving

**Approach**: Optional configuration to move completed tasks to done.txt

## Dependencies Summary

**Core Dependencies**:
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/onsi/ginkgo/v2` - Testing framework
- `github.com/onsi/gomega` - Test matchers

**Optional Dependencies** (evaluate during implementation):
- `github.com/fsnotify/fsnotify` - File watching
- `github.com/mitchellh/mapstructure` - Config parsing
- `gopkg.in/yaml.v3` - YAML config support

## Next Steps

1. **Phase 1**: Define data model with Task, TodoFile, Filter, TaskList entities
2. **Phase 1**: Generate API contracts for internal interfaces
3. **Phase 1**: Create quickstart guide for developers
4. **Phase 2**: Break down implementation into tasks
5. **Phase 2**: Implement parser with full todo.txt compliance
6. **Phase 2**: Build basic TUI with task list display
7. **Phase 2**: Add vim-style navigation
8. **Phase 2**: Implement task CRUD operations
9. **Phase 2**: Add filtering and search
10. **Phase 2**: Optimize for large datasets

## References

- todo.txt specification: https://github.com/todotxt/todo.txt
- bubbletea documentation: https://github.com/charmbracelet/bubbletea
- lipgloss documentation: https://github.com/charmbracelet/lipgloss
- Ginkgo documentation: https://onsi.github.io/ginkgo/
- Effective Go: https://golang.org/doc/effective_go

## Appendix

### Big O Notation

Big O describes how an algorithm's execution time or space requirements grow as the input size increases, focusing on the dominant factor while ignoring constants and smaller terms.

**Common Orders** (from fastest to slowest):

- **O(1)** - Constant: Same time regardless of input size (array access)
- **O(log n)** - Logarithmic: Doubles input, adds one operation (binary search)
- **O(n)** - Linear: Double input, double time (scanning a list)
- **O(n log n)** - Linearithmic: Efficient sorting algorithms
- **O(n²)** - Quadratic: Nested loops (bubble sort)
- **O(2ⁿ)** - Exponential: Trying all combinations (very slow)

For the TUI application, we prioritize O(1) and O(log n) operations for interactive features (scrolling, navigation) and accept O(n) operations for background tasks (file parsing, initial indexing).
