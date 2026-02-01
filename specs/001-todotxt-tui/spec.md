# Feature Specification: todotxt-tui

**Feature Branch**: `001-todotxt-tui`
**Created**: 2026-02-01
**Status**: Draft
**Input**: User description: "Develop todotxt-tui, TODO list TUI that manages TODOs in todo.txt format (https://github.com/todotxt/todo.txt) and vim-like keybind."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - View and Navigate Tasks (Priority: P1)

Users can view their todo.txt file as an interactive list in the terminal and navigate through tasks using vim-style keyboard shortcuts. The application loads tasks from a default location and displays them in a clean, readable format with priority indicators visible at a glance.

**Why this priority**: This is the MVP foundation. Without the ability to see and navigate tasks, no other functionality is useful. Users need to quickly scan their task list to understand what needs to be done.

**Independent Test**: Can be fully tested by launching the application with a sample todo.txt file and verifying: tasks load correctly, all vim-style navigation keys work (j/k, g/G, h/l), and the UI displays tasks with proper formatting and priority indicators.

**Acceptance Scenarios**:

1. **Given** a user has a todo.txt file with 10 tasks, **When** they launch the application, **Then** all tasks display in the terminal window in list format
2. **Given** the application is open with 20 tasks, **When** the user presses `j`, **Then** the cursor moves down to the next task
3. **Given** the application is open, **When** the user presses `k`, **Then** the cursor moves up to the previous task
4. **Given** the application is open, **When** the user presses `g`, **Then** the cursor moves to the first task
5. **Given** the application is open, **When** the user presses `G`, **Then** the cursor moves to the last task
6. **Given** the application is displaying a task with priority (A), **When** viewing the list, **Then** the priority indicator is visually distinct from non-priority tasks

---

### User Story 2 - Create and Edit Tasks (Priority: P2)

Users can create new tasks and edit existing tasks with full support for todo.txt format including priorities (A-Z), contexts (@context), projects (+project), and creation dates. The interface provides intuitive prompts and validation to ensure tasks conform to the standard format.

**Why this priority**: After viewing tasks, users need to be able to add new tasks and modify existing ones. This enables full task management and makes the application functional for daily use.

**Independent Test**: Can be fully tested by creating a new task with priority, context, and project tags, then editing it to change each attribute. Verifies the task persists correctly to the file and displays properly.

**Acceptance Scenarios**:

1. **Given** the application is open, **When** the user presses `a` to add a new task, **Then** a prompt appears to enter task details
2. **Given** the user is adding a task, **When** they enter "Buy groceries @store +household", **Then** the task is saved with context and project tags
3. **Given** the user is adding a task, **When** they select priority (A) before entering text, **Then** the task is saved with "(A)" prefix
4. **Given** the application has a task selected, **When** the user presses `e` to edit, **Then** the task text becomes editable in a prompt
5. **Given** the user is editing a task, **When** they change the priority from (B) to (A), **Then** the task is saved with the new priority and repositioned in the list
6. **Given** the user is editing a task, **When** they add a new context tag, **Then** the updated task displays the new context

---

### User Story 3 - Complete and Delete Tasks (Priority: P2)

Users can mark tasks as complete and delete tasks they no longer need. Completed tasks are automatically prefixed with "x" and completion date according to todo.txt format, and optionally can be archived to a separate done.txt file.

**Why this priority**: Task completion and deletion are core lifecycle operations. Without them, users cannot manage their task list effectively or clear completed items.

**Independent Test**: Can be fully tested by marking tasks as complete and verifying they are properly formatted with "x" and date, then deleting tasks and confirming they are removed from the file and UI.

**Acceptance Scenarios**:

1. **Given** the application has a task selected, **When** the user presses `Space` to complete it, **Then** the task is marked with "x" and today's date
2. **Given** a task was just completed, **When** the user views the task list, **Then** the completed task appears at the bottom of the list or in a completed section
3. **Given** the application has a task selected, **When** the user presses `d` to delete, **Then** the task is removed from the file and list after confirmation
4. **Given** a completed task exists, **When** the user deletes it, **Then** the task is permanently removed from the todo.txt file
5. **Given** archive is configured, **When** a task is completed, **Then** the task is moved to done.txt with proper formatting

---

### User Story 4 - Filter and Search Tasks (Priority: P3)

Users can filter the task list by priority, context, project, or search by text content to quickly find relevant tasks. Filters can be combined and toggled to help users focus on specific subsets of their task list.

**Why this priority**: As task lists grow, filtering and search become essential for productivity. This is a power-user feature that enhances the application's utility but is not required for basic functionality.

**Independent Test**: Can be fully tested by applying various filters (priority A only, @phone context, +project project) and verifying only matching tasks display. Search by text and confirm results match the query.

**Acceptance Scenarios**:

1. **Given** the application has 50 tasks, **When** the user applies a filter for priority (A), **Then** only tasks with (A) priority are displayed
2. **Given** the application has tasks with contexts, **When** the user filters by @phone, **Then** only tasks containing @phone are displayed
3. **Given** the application has tasks with projects, **When** the user filters by +project, **Then** only tasks with that project tag are displayed
4. **Given** the application has a filter active, **When** the user clears the filter, **Then** all tasks are displayed again
5. **Given** the application has 100 tasks, **When** the user searches for "meeting", **Then** only tasks containing "meeting" are displayed
6. **Given** priority (A) and context @phone filters are active, **When** both filters are applied, **Then** only tasks matching both criteria are displayed

---

### Edge Cases

- What happens when the todo.txt file does not exist at startup?
- What happens when the todo.txt file has invalid or malformed task lines?
- What happens when the user tries to edit a task to an empty string?
- What happens when the file is externally modified while the application is running?
- What happens when the terminal window is resized during use?
- What happens when a task has multiple priorities or malformed priority tags?
- What happens when special characters or emoji are used in task text?
- What happens when the file cannot be saved due to permissions or disk space?
- What happens when a task line exceeds the terminal width?
- What happens when there are no tasks in the file?
- What happens when an unrecognized keyboard shortcut is pressed?
- What happens when the user tries to delete the last remaining task?
- What happens when trying to increase priority of a task already at (A)?
- What happens when trying to complete an already completed task?

### Keyboard Shortcuts

**Navigation**:
- `j` / `k` - Move cursor down/up through tasks
- `g` / `G` - Jump to first/last task
- `h` / `l` - Navigate left/right (in dialogs, filter panels)

**Task Management**:
- `a` - Add/create new task
- `e` - Edit selected task
- `Space` - Toggle task completion (mark as complete/incomplete)
- `d` - Delete selected task (confirm before deletion)
- `u` - Undo last delete (single-level undo)

**Priority Editing**:
- `+` / `-` - Increase/decrease priority of selected task (A→B→C→...→no priority)
- `0` - Remove priority from selected task

**Filtering and Searching**:
- `/` - Open search prompt to filter by text
- `f` - Open filter menu (priority, context, project filters)
- `c` - Clear all active filters
- `n` / `p` - Navigate to next/previous priority task

**File Operations**:
- `r` - Reload file from disk (detect external changes)
- `s` - Save changes manually (auto-save also occurs after modifications)

**System**:
- `q` or `Esc` - Quit application (Esc may cancel current operation instead)
- `?` or `F1` - Display help screen with all keyboard shortcuts

**Context Actions** (when task selected):
- `Enter` - View task details or confirm action in dialogs
- `Tab` - Cycle through context tags or filter options

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST load and parse tasks from a todo.txt file located at a configurable default path
- **FR-002**: System MUST display tasks in an interactive terminal user interface (TUI)
- **FR-003**: System MUST support vim-style navigation: j/k for down/up, g/G for top/bottom, h/l for left/right where applicable
- **FR-004**: System MUST parse todo.txt format including: priority (A-Z), creation date (YYYY-MM-DD), contexts (@tag), projects (+tag), and task description
- **FR-005**: System MUST display task priority indicators visually (color-coded or symbol-based)
- **FR-006**: System MUST allow users to create new tasks with support for priority, context, and project tags
- **FR-007**: System MUST allow users to edit existing tasks with support for all todo.txt format elements
- **FR-008**: System MUST mark tasks as complete by prefixing with "x" and completion date (YYYY-MM-DD)
- **FR-009**: System MUST allow users to delete tasks from the file
- **FR-010**: System MUST persist all changes to the todo.txt file immediately after modification
- **FR-011**: System MUST support filtering tasks by priority (A-Z)
- **FR-012**: System MUST support filtering tasks by context (@tag)
- **FR-013**: System MUST support filtering tasks by project (+tag)
- **FR-014**: System MUST support searching tasks by text content
- **FR-015**: System MUST detect external file changes and reload when prompted
- **FR-016**: System MUST handle terminal resize events and redraw the interface appropriately
- **FR-017**: System MUST display error messages for file read/write failures
- **FR-018**: System MUST handle invalid task lines gracefully (display as-is or show warning)
- **FR-019**: System MUST support keyboard shortcuts for common actions (create, edit, complete, delete, filter, search, quit)
- **FR-020**: System MUST display a help screen showing available keyboard shortcuts

### Key Entities

- **Task**: Represents a single todo item with attributes: priority (A-Z or none), creation date (YYYY-MM-DD or none), completion status (complete/incomplete), completion date (if complete), description (text), contexts (list of @tags), projects (list of +tags), additional metadata (key:value pairs)

- **TodoFile**: Represents the todo.txt text file containing all tasks, with properties: file path, list of tasks, last modified timestamp, encoding (UTF-8)

- **Filter**: Represents an active filter criteria with properties: type (priority/context/project/search), value (specific priority letter, tag, or search text), combination logic (AND/OR)

- **TaskList**: Represents the filtered and sorted view of tasks displayed to users, with properties: all tasks list, active filters, sort order, currently selected task index, visible task count

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can load a todo.txt file with 1000 tasks and view them in under 100 milliseconds
- **SC-002**: Users can navigate through a task list of 100 tasks using vim-style keys with visible cursor updates within 16 milliseconds per action
- **SC-003**: Users can complete a task (mark with "x" and date) with 3 or fewer key presses
- **SC-004**: Users can create a new task with priority, context, and project tags in under 10 seconds
- **SC-005**: System correctly parses and displays 100% of valid todo.txt format tasks according to the official specification
- **SC-006**: Users can filter 1000 tasks by priority, context, or project and see results in under 50 milliseconds
- **SC-007**: System persists all changes to the file with 100% reliability (no data loss on save)
- **SC-008**: 90% of new users can successfully complete the core workflow (view, create, edit, complete tasks) within 5 minutes without documentation
- **SC-009**: System handles terminal resize events without losing cursor position or scroll position
- **SC-010**: System uses less than 50MB of memory during normal operation with 1000 tasks

## Assumptions

- The todo.txt file uses UTF-8 encoding (standard for plain text files)
- Default todo.txt file location is `~/.local/share/todotxt-tui/todo.txt` unless configured otherwise
- Users are familiar with vim-style navigation patterns
- Terminal supports at least 80x24 character dimensions and ANSI color codes
- System clock provides accurate current date in YYYY-MM-DD format for task completion
- File system permissions allow read/write access to the todo.txt file location
- Tasks are limited to reasonable length for display (no hard limit enforced, but 500+ characters may need truncation or wrapping)
- Contexts and projects are single words without spaces (per todo.txt specification)
- Completed tasks should be sorted to the bottom of the list by default
- Application should not require network connectivity or external services
