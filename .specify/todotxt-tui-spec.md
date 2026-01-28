# Feature Specification: TodoTxt TUI with Vim-like Keybindings

**Feature Branch**: `001-todotxt-tui`  
**Created**: 2026-01-26  
**Status**: Draft  
**Input**: User description: "TODO list TUI that uses todo.txt format and vim-like keybind"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - View and Navigate Todo List (Priority: P1)

As a user, I want to view my todo.txt file in a terminal user interface and navigate through tasks using vim-like keybindings so that I can quickly review my tasks without leaving the terminal.

**Why this priority**: This is the foundational MVP feature. Without the ability to view and navigate tasks, no other functionality is useful. This provides immediate value by offering a keyboard-driven interface for todo.txt files.

**Independent Test**: Can be fully tested by launching the TUI with a sample todo.txt file, navigating with j/k keys, and verifying the correct task is highlighted. Delivers value as a read-only todo.txt viewer.

**Acceptance Scenarios**:

1. **Given** a todo.txt file exists with multiple tasks, **When** I launch the TUI, **Then** all tasks are displayed in a list view with the first task highlighted
2. **Given** the TUI is displaying tasks, **When** I press 'j', **Then** the highlight moves down to the next task
3. **Given** the TUI is displaying tasks, **When** I press 'k', **Then** the highlight moves up to the previous task
4. **Given** I am on the first task, **When** I press 'k', **Then** the highlight stays on the first task (no wrapping)
5. **Given** I am on the last task, **When** I press 'j', **Then** the highlight stays on the last task (no wrapping)
6. **Given** the TUI is displaying tasks, **When** I press 'gg', **Then** the highlight jumps to the first task
7. **Given** the TUI is displaying tasks, **When** I press 'G', **Then** the highlight jumps to the last task
8. **Given** the TUI is displaying tasks, **When** I press 'q', **Then** the application exits cleanly

---

### User Story 2 - Mark Tasks as Complete/Incomplete (Priority: P1)

As a user, I want to toggle task completion status using a simple keybinding so that I can quickly check off completed tasks and maintain my todo.txt file in the standard format.

**Why this priority**: Core functionality for a todo application. Completing tasks is the primary action users perform, making this essential for the MVP.

**Independent Test**: Can be tested by opening a todo.txt file, pressing 'x' on an incomplete task, verifying the task is marked with 'x' and the date prefix, and confirming the file is saved correctly. Delivers immediate value for task management.

**Acceptance Scenarios**:

1. **Given** an incomplete task is highlighted, **When** I press 'x', **Then** the task is marked as complete with 'x ' prefix and completion date in format 'x YYYY-MM-DD'
2. **Given** a completed task is highlighted, **When** I press 'x', **Then** the 'x' prefix and completion date are removed and the task becomes incomplete
3. **Given** I toggle task completion, **When** I save the file, **Then** the changes are persisted to the todo.txt file
4. **Given** a task is marked complete, **When** viewing the list, **Then** completed tasks are visually distinct (e.g., strikethrough, different color, or moved to bottom)

---

### User Story 3 - Add New Tasks (Priority: P1)

As a user, I want to add new tasks using a vim-like insert mode so that I can quickly capture new todos without switching to an external editor.

**Why this priority**: Task creation is a fundamental operation. Without the ability to add tasks, users must exit the TUI to modify their todo.txt file, defeating the purpose of a TUI.

**Independent Test**: Can be tested by pressing 'o' to enter insert mode, typing a new task, pressing ESC to exit insert mode, and verifying the task appears in the list and is saved to the file.

**Acceptance Scenarios**:

1. **Given** the TUI is displaying tasks, **When** I press 'o', **Then** a new task entry line appears below the current task and insert mode is activated
2. **Given** I am in insert mode for a new task, **When** I type text, **Then** the text appears in the input field
3. **Given** I am in insert mode with text entered, **When** I press ESC or Ctrl+C, **Then** the new task is saved to the list and added to the todo.txt file, and I return to normal mode
4. **Given** the TUI is displaying tasks, **When** I press 'O' (capital O), **Then** a new task entry line appears above the current task and insert mode is activated
5. **Given** I am in insert mode, **When** I press ESC without entering any text, **Then** insert mode is cancelled and no empty task is created

---

### User Story 4 - Delete Tasks (Priority: P2)

As a user, I want to delete tasks using vim-like delete commands so that I can remove tasks that are no longer relevant.

**Why this priority**: Task deletion is important for maintenance but less critical than viewing, completing, and adding tasks. Users can work around this by manually editing the file if needed.

**Independent Test**: Can be tested by highlighting a task, pressing 'dd', confirming deletion, and verifying the task is removed from both the list and the file.

**Acceptance Scenarios**:

1. **Given** a task is highlighted, **When** I press 'dd', **Then** the task is deleted from the list and removed from the todo.txt file
2. **Given** a task was just deleted, **When** I press 'u' (undo), **Then** the task is restored to its original position
3. **Given** the TUI is displaying tasks, **When** I delete the last task, **Then** the highlight moves to the new last task
4. **Given** only one task exists, **When** I delete it, **Then** the list becomes empty with a message "No tasks found"

---

### User Story 5 - Edit Existing Tasks (Priority: P2)

As a user, I want to edit task text using vim-like edit commands so that I can update task descriptions, priorities, projects, and contexts.

**Why this priority**: Editing tasks is common but not as frequent as viewing, adding, or completing tasks. Users can work around this by deleting and re-adding tasks if needed.

**Independent Test**: Can be tested by highlighting a task, pressing 'e' or 'i', modifying the text, and verifying the changes are saved to the file.

**Acceptance Scenarios**:

1. **Given** a task is highlighted, **When** I press 'e', **Then** the task enters edit mode with the full task text editable
2. **Given** I am in edit mode, **When** I modify the text and press ESC, **Then** the changes are saved to the todo.txt file and I return to normal mode
3. **Given** I am in edit mode, **When** I press Ctrl+C, **Then** the edit is cancelled, no changes are saved, and I return to normal mode
4. **Given** I am in edit mode, **When** I clear all text and press ESC, **Then** the task is not deleted but remains as an empty line (user must use 'dd' to delete)

---

### User Story 6 - Filter and Search Tasks (Priority: P2)

As a user, I want to filter tasks by project, context, or search for keywords so that I can focus on specific subsets of my todo list.

**Why this priority**: Filtering is valuable for power users with many tasks but not essential for basic functionality. Users can manually scan their lists if needed.

**Independent Test**: Can be tested by pressing '/' to enter search mode, typing a search term, and verifying that only matching tasks are displayed. Pressing ESC clears the filter.

**Acceptance Scenarios**:

1. **Given** the TUI is displaying tasks, **When** I press '/', **Then** a search prompt appears at the bottom
2. **Given** I am in search mode, **When** I type a keyword and press ENTER, **Then** only tasks containing that keyword are displayed
3. **Given** tasks are filtered, **When** I press ESC, **Then** the filter is cleared and all tasks are displayed
4. **Given** the TUI is displaying tasks, **When** I press '@' followed by a context name, **Then** only tasks with that context are displayed
5. **Given** the TUI is displaying tasks, **When** I press '+' followed by a project name, **Then** only tasks with that project are displayed

---

### User Story 7 - Set Task Priority (Priority: P3)

As a user, I want to quickly set or change task priorities using keybindings so that I can maintain task importance according to todo.txt format (A-Z).

**Why this priority**: Priority management is useful but less critical than core CRUD operations. Many users don't use priorities extensively, and they can be edited manually.

**Independent Test**: Can be tested by highlighting a task, pressing 'p' followed by a priority letter (A-Z), and verifying the priority is added/updated in the format '(A) task text'.

**Acceptance Scenarios**:

1. **Given** a task without priority is highlighted, **When** I press 'p' followed by 'a', **Then** the task is updated with '(A) ' prefix
2. **Given** a task with priority (B) is highlighted, **When** I press 'p' followed by 'a', **Then** the priority is updated to '(A)'
3. **Given** a task with priority is highlighted, **When** I press 'p' followed by '0' or DEL, **Then** the priority is removed
4. **Given** the TUI is displaying tasks, **When** tasks have different priorities, **Then** they are optionally sorted by priority (A > B > ... > no priority)

---

### User Story 8 - Sort and Organize Tasks (Priority: P3)

As a user, I want to sort tasks by various criteria (priority, date, project, completion status) so that I can organize my todo list according to my workflow preferences.

**Why this priority**: Sorting enhances usability but is not essential for basic task management. Users can manually organize their todo.txt file if needed.

**Independent Test**: Can be tested by pressing 's' to cycle through sort modes (priority, date, alphabetical) and verifying tasks are reordered accordingly.

**Acceptance Scenarios**:

1. **Given** the TUI is displaying tasks, **When** I press 's', **Then** a sort menu appears with options (priority, date, alphabetical, none)
2. **Given** I select priority sort, **When** tasks are displayed, **Then** they are ordered by priority (A-Z, then unprioritized)
3. **Given** I select date sort, **When** tasks are displayed, **Then** they are ordered by creation date (newest/oldest first)
4. **Given** completed tasks exist, **When** I press 'h', **Then** completed tasks are hidden/shown (toggle)

---

### User Story 9 - View Task Details and Metadata (Priority: P3)

As a user, I want to view detailed task metadata (creation date, completion date, projects, contexts) in a detail pane so that I can see all task information at a glance.

**Why this priority**: Detail view enhances usability but the main list view can display most essential information. Power users will appreciate this, but it's not critical for MVP.

**Independent Test**: Can be tested by highlighting a task and pressing ENTER or 'v' to toggle a detail pane showing all task metadata parsed according to todo.txt format.

**Acceptance Scenarios**:

1. **Given** a task is highlighted, **When** I press ENTER, **Then** a detail pane appears showing task metadata (priority, creation date, projects, contexts, key-value pairs)
2. **Given** the detail pane is open, **When** I press ESC or ENTER again, **Then** the detail pane closes
3. **Given** the detail pane is open, **When** I navigate with j/k, **Then** the detail pane updates to show the currently highlighted task
4. **Given** a task has multiple projects and contexts, **When** viewing details, **Then** all projects (prefixed with +) and contexts (prefixed with @) are listed separately

---

### Edge Cases

- What happens when the todo.txt file is empty? → Display "No tasks found. Press 'o' to add a new task."
- What happens when the todo.txt file doesn't exist? → Create a new empty file and display the empty state message.
- What happens when the todo.txt file is modified externally while the TUI is open? → Detect file changes and prompt to reload.
- What happens when the terminal window is resized? → Automatically reflow the layout to fit the new dimensions.
- What happens when invalid todo.txt format is encountered? → Display the line as-is with a warning indicator, but don't crash.
- What happens when very long task descriptions exceed terminal width? → Truncate with ellipsis or wrap to multiple lines based on configuration.
- What happens when there are thousands of tasks? → Implement virtual scrolling for performance; only render visible tasks.
- What happens when the user presses an unmapped key? → Show a brief "Unknown command" message in the status bar.
- What happens when the file cannot be saved (permissions, disk full)? → Show error message and keep changes in memory with option to retry or save to alternate location.
- What happens when trying to undo with no undo history? → Show "Nothing to undo" message in status bar.

## Requirements *(mandatory)*

### Functional Requirements

#### Core Data & File Operations
- **FR-001**: System MUST read and parse todo.txt files following the [todo.txt format specification](https://github.com/todotxt/todo.txt)
- **FR-002**: System MUST support all standard todo.txt elements: completion status (x), priority (A-Z), dates, projects (+project), contexts (@context), and key:value metadata
- **FR-003**: System MUST preserve the exact format and ordering of unrecognized metadata when saving
- **FR-004**: System MUST write changes back to the todo.txt file atomically (temp file + rename) to prevent corruption
- **FR-005**: System MUST handle file encoding as UTF-8
- **FR-006**: System MUST auto-save changes immediately after each action

#### Navigation & Interaction
- **FR-007**: System MUST implement vim-like normal mode with keybindings: j (down), k (up), gg (top), G (bottom), Ctrl+d (page down), Ctrl+u (page up)
- **FR-008**: System MUST implement insert mode for adding/editing tasks (i, o, O keys to enter, ESC to exit)
- **FR-009**: System MUST display a visual indicator of the current mode (normal/insert/search)
- **FR-010**: System MUST highlight the currently selected task with distinct visual styling
- **FR-011**: System MUST implement vim-like command mode (press ':' to enter) for advanced operations

#### Task Operations
- **FR-012**: System MUST allow toggling task completion status with 'x' key
- **FR-013**: System MUST automatically add/remove completion date (YYYY-MM-DD format) when toggling completion
- **FR-014**: System MUST allow adding new tasks (o/O keys) with immediate entry into insert mode
- **FR-015**: System MUST allow editing existing tasks (e/i keys) with full text editing capabilities
- **FR-016**: System MUST allow deleting tasks (dd key combination) with confirmation
- **FR-017**: System MUST support undo (u) and redo (Ctrl+r) for destructive operations
- **FR-018**: System MUST allow setting/changing task priority (p key followed by a-z or 0 to remove)

#### Filtering & Search
- **FR-019**: System MUST implement search functionality (/ key) with live filtering as user types
- **FR-020**: System MUST support filtering by project (+ prefix) and context (@ prefix)
- **FR-021**: System MUST allow clearing filters (ESC key) to return to full list view
- **FR-022**: System MUST display filter status in the status bar

#### Display & Visualization
- **FR-023**: System MUST visually distinguish completed tasks from incomplete tasks (strikethrough, color, or dimmed)
- **FR-024**: System MUST syntax-highlight projects (+), contexts (@), and priorities ((A))
- **FR-025**: System MUST display a status bar with file name, task count, current mode, and help hint
- **FR-026**: System MUST handle terminal resize events and reflow layout accordingly
- **FR-027**: System MUST support color schemes (at minimum: default and monochrome for accessibility)

#### Sorting & Organization
- **FR-028**: System MUST support multiple sort modes: priority, creation date, alphabetical, and original file order
- **FR-029**: System MUST allow toggling visibility of completed tasks (h key)
- **FR-030**: System MUST persist sort and filter preferences in config file

#### Error Handling & Validation
- **FR-031**: System MUST display actionable error messages for file I/O failures (permission denied, disk full, etc.)
- **FR-032**: System MUST validate priority input (a-z only) and reject invalid input with feedback
- **FR-033**: System MUST handle malformed todo.txt lines gracefully (display as-is with warning indicator)
- **FR-034**: System MUST prevent data loss by prompting to save unsaved changes before exit

#### Configuration & Customization
- **FR-035**: System MUST support configuration file for keybindings in YAML format
- **FR-036**: System MUST support configuration for color schemes and visual styling
- **FR-037**: System MUST accept todo.txt file path as command-line argument or use default location specified in the configuration file

### Key Entities

- **Task**: Represents a single todo item with the following attributes:
  - **Completion Status**: Boolean (complete/incomplete) with optional completion date
  - **Priority**: Single letter A-Z (optional)
  - **Creation Date**: Date in YYYY-MM-DD format (optional, typically added automatically)
  - **Completion Date**: Date in YYYY-MM-DD format (optional, added when task is completed)
  - **Description**: Free-form text content
  - **Projects**: List of project tags (strings prefixed with +)
  - **Contexts**: List of context tags (strings prefixed with @)
  - **Metadata**: Key-value pairs (e.g., due:2026-01-31)
  - **Raw Line**: Original file line for preserving unknown formats

- **TodoFile**: Represents the todo.txt file being edited
  - **Path**: File system path to todo.txt
  - **Tasks**: Ordered list of Task entities
  - **Modified**: Boolean flag indicating unsaved changes
  - **Encoding**: Character encoding (UTF-8)

- **ViewState**: Represents the current UI state
  - **Selected Index**: Currently highlighted task index
  - **Scroll Offset**: Vertical scroll position for viewport
  - **Mode**: Current interaction mode (normal/insert/command/search)
  - **Filter**: Active filter criteria (search term, project, context)
  - **Sort Mode**: Active sort method (priority/date/alpha/none)
  - **Show Completed**: Boolean flag for completed task visibility

- **Configuration**: User preferences and customization
  - **Keybindings**: Map of actions to key combinations
  - **Colors**: Color scheme definitions
  - **Default File Path**: Default todo.txt location
  - **Auto-save**: Boolean flag for automatic saving

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can navigate through a 100-task todo.txt file using vim keybindings with < 100ms response time per keystroke
- **SC-002**: Users can complete the core workflow (add task, mark complete, delete task) within 30 seconds of first launch
- **SC-003**: System handles todo.txt files up to 10,000 tasks without noticeable performance degradation (< 500ms load time)
- **SC-004**: System preserves 100% of existing todo.txt format including unknown metadata and custom key-value pairs
- **SC-005**: UI is responsive and usable in terminal windows as small as 80x24 characters
- **SC-006**: All keybindings follow vim conventions such that users familiar with vim can use the TUI without reading documentation
- **SC-007**: System achieves 0% data loss rate across file save operations (atomic writes with error handling)
- **SC-008**: 90% of users can successfully toggle task completion status on first attempt without help
- **SC-009**: Application memory usage stays below 256MB even with 10,000 tasks loaded
- **SC-010**: System provides meaningful error messages with actionable recovery steps for 100% of error scenarios

## Technical Considerations *(optional)*

### Technology Stack (Implementation Detail - Not Prescribed)

While implementation details are outside the scope of this specification, the following technical considerations should inform the implementation:

- **Language**: Golang (per project constitution)
- **TUI Framework**: Consider libraries like bubbletea, tview, or termui for building the terminal interface
- **File Watching**: Implement file system watching to detect external changes to todo.txt
- **Testing**: All functionality MUST be unit tested with Ginkgo/Gomega achieving 80%+ coverage
- **Performance**: Use virtual scrolling for large lists (only render visible portion)
- **Concurrency**: File I/O operations should be non-blocking where appropriate

### Configuration File Format (Recommendation)

Suggest using YAML or TOML for configuration file with structure:

```yaml
todo_file: ~/.local/share/todotxt/todo.txt
auto_save: true
show_completed: true
default_sort: priority
colors:
  completed: gray
  priority_a: red
  priority_b: yellow
  project: blue
  context: green
keybindings:
  quit: q
  toggle_complete: x
  delete: dd
  # ... additional bindings
```

### Accessibility Considerations

- Support monochrome mode for users who cannot distinguish colors
- Ensure all functionality is keyboard-accessible (no mouse required)
- Provide alternative visual indicators beyond color (e.g., symbols, strikethrough)

### Future Enhancements (Out of Scope for MVP)

- Multi-file support (multiple todo.txt files)
- Cloud sync integration
- Due date reminders and notifications
- Recurrence rules for repeating tasks
- Bulk operations (select multiple tasks)
- Export to other formats (JSON, CSV, Markdown)
- Archive completed tasks to done.txt
- Statistics and productivity metrics
- Split view for multiple perspectives

## Definition of Done

This feature is considered complete when:

1. All P1 user stories are fully implemented and tested
2. Unit tests achieve 80%+ code coverage using Ginkgo/Gomega
3. Integration tests validate file I/O operations
4. Manual testing confirms all vim-like keybindings work as expected
5. Performance benchmarks meet defined criteria (SC-001, SC-003, SC-009)
6. Error handling covers all edge cases with meaningful messages
7. README documentation includes usage guide with keybinding reference
8. Code passes all CI checks (gofmt, golangci-lint, go vet, tests)
9. The application successfully manages real-world todo.txt files without data loss

## References

- [Todo.txt Format Specification](https://github.com/todotxt/todo.txt)
- [Vim Keybinding Reference](https://vim.rtorr.com/)
- [Go TUI Libraries Comparison](https://charm.sh/)
- Project Constitution: `.specify/memory/constitution.md`
