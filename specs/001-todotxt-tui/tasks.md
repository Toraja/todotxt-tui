# Tasks: todotxt-tui

**Input**: Design documents from `/specs/001-todotxt-tui/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Testing is REQUIRED for all code (Constitution Principle II - Testing with Ginkgo is NON-NEGOTIABLE).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3, US4)
- Include exact file paths in descriptions

## Path Conventions

Based on plan.md project structure:
- **cmd/todotxt-tui/**: Main application entry point
- **internal/**: Private packages (config, ui, parser, filter, storage, keymap)
  - Unit tests adjacent to source files (e.g., internal/parser/task_test.go)
- **tests/**: Integration tests and fixtures only
  - tests/integration/: Integration tests
  - tests/fixtures/: Test data and sample files

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Create project directory structure: cmd/todotxt-tui/, internal/{config,ui,parser,filter,storage,keymap}/, tests/{unit,integration,fixtures}/, docs/
- [X] T002 Initialize Go module with go mod init and add dependencies: bubbletea, lipgloss, ginkgo/v2, gomega
- [X] T003 [P] Create Makefile with targets: build, run, test, lint, clean
- [X] T004 [P] Setup .gitignore for Go project (bin/, *.test, coverage.out)
- [X] T005 [P] Create initial README.md with project description and build instructions
- [X] T006 [P] Setup golangci-lint configuration in .golangci.yml
- [X] T007 [P] Create default config directory structure at ~/.local/share/todotxt-tui/ and ~/.config/todotxt-tui/

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

### Parser - todo.txt Format Support

- [X] T008 [P] Create Task struct in internal/parser/task.go with all todo.txt fields (Priority, CreationDate, Completed, CompletionDate, Description, Contexts, Projects, Metadata, RawLine, LineNumber, Modified)
- [X] T009 [P] Write Ginkgo unit tests for Task struct in internal/parser/task_test.go
- [X] T010 [P] Implement Parser interface in internal/parser/parser.go with ParseLine, ParseFile, Serialize, Validate methods
- [X] T011 [P] Write Ginkgo unit tests for Parser.ParseLine in internal/parser/parser_test.go with table-driven tests for all todo.txt format cases (priority, dates, contexts, projects, completion)
- [X] T012 [P] Write Ginkgo unit tests for Parser.ParseFile in internal/parser/parser_test.go
- [X] T013 Implement ParseLine method with regex patterns for priority, dates, contexts, projects, metadata extraction in internal/parser/parser.go
- [X] T014 Implement ParseFile method for parsing entire file content in internal/parser/parser.go
- [X] T015 Implement Serialize method to convert Task to todo.txt format string in internal/parser/parser.go
- [X] T016 Implement date parsing helper ParseTodoDate(dateStr string) (time.Time, error) in internal/parser/date.go
- [X] T017 [P] Write Ginkgo unit tests for date parsing in internal/parser/date_test.go
- [X] T018 Implement context/project tag validation helpers (ValidateContextTag, ValidateProjectTag) in internal/parser/validation.go
- [X] T019 [P] Write Ginkgo unit tests for tag validation in internal/parser/validation_test.go

### Storage - File I/O and Persistence

- [X] T020 [P] Create Storage interface in internal/storage/storage.go with Load, Save, Watch, Exists, Create, GetModificationTime methods
- [X] T021 [P] Create FileEvent struct and EventType enum in internal/storage/event.go
- [ ] T022 [P] Write Ginkgo unit tests for Storage interface in internal/storage/storage_test.go
- [X] T023 Implement FileStorage struct with atomic write logic (temp file + rename) in internal/storage/file_storage.go
- [X] T024 Implement Load method with UTF-8 encoding and error handling in internal/storage/file_storage.go
- [X] T025 Implement Save method with atomic write (write to temp, rename on success) in internal/storage/file_storage.go
- [X] T026 Implement Watch method using fsnotify or polling in internal/storage/file_storage.go
- [X] T027 Implement file existence and creation helpers in internal/storage/file_storage.go
- [ ] T028 [P] Write Ginkgo integration tests for file I/O operations in tests/integration/storage/file_storage_test.go with fixtures

### Configuration Management

- [X] T029 [P] Create Config struct in internal/config/config.go with all settings (TodoFilePath, DoneFilePath, Theme, ShowCompleted, ArchiveCompleted, ConfirmDelete, AutoSave, FileWatchEnabled)
- [ ] T030 [P] Write Ginkgo unit tests for Config in internal/config/config_test.go
- [X] T031 Implement config loading from YAML file in internal/config/loader.go with environment variable overrides
- [X] T032 Implement default configuration values in internal/config/defaults.go
- [X] T033 Implement config validation in internal/config/validation.go
- [ ] T034 [P] Create sample config.yaml in tests/fixtures/config/sample_config.yaml

### Data Model Core

- [X] T035 [P] Create TodoFile struct in internal/storage/todofile.go with Tasks, Path, LastModified, Modified, Loaded, LoadError fields
- [ ] T036 [P] Write Ginkgo unit tests for TodoFile in internal/storage/todofile_test.go
- [X] T037 Implement TodoFile methods: Load, Save, AddTask, UpdateTask, DeleteTask, GetTask, Count, CountCompleted, CountActive, ReloadIfChanged in internal/storage/todofile.go
- [X] T038 [P] Create TaskList struct in internal/filter/tasklist.go with AllTasks, ActiveFilters, FilterLogic, SortBy, SortOrder, VisibleTasks, SelectedIndex, ScrollOffset, ViewportSize fields
- [ ] T039 [P] Write Ginkgo unit tests for TaskList in internal/filter/tasklist_test.go

### Filter and Search Infrastructure

- [X] T040 [P] Create Filter interface in internal/filter/filter.go with Apply, BuildIndex, Search, FilterByPriority, FilterByContext, FilterByProject, FilterByCompletion methods
- [X] T041 [P] Create FilterCriteria struct and FilterLogic enum in internal/filter/criteria.go
- [X] T042 [P] Create Index struct with priority, context, project, completion indexes in internal/filter/index.go
- [ ] T043 [P] Write Ginkgo unit tests for Filter interface in internal/filter/filter_test.go
- [X] T044 Implement Filter implementation with indexing logic in internal/filter/filter_impl.go
- [X] T045 Implement BuildIndex method to create lookup maps in internal/filter/index.go
- [X] T046 Implement Apply method with AND/OR combination logic in internal/filter/filter_impl.go
- [X] T047 Implement Search method with case-insensitive text matching in internal/filter/filter_impl.go
- [X] T048 Implement individual filter methods (FilterByPriority, FilterByContext, FilterByProject, FilterByCompletion) in internal/filter/filter_impl.go

### Keymap and Input Handling

- [X] T049 [P] Create Keymap interface in internal/keymap/keymap.go with GetBinding, SetBinding, GetAvailableActions, GetKeysForAction methods
- [X] T050 [P] Create Mode enum (Normal, Insert, Dialog, Search) and Action enum in internal/keymap/types.go
- [ ] T051 [P] Write Ginkgo unit tests for Keymap in internal/keymap/keymap_test.go
- [X] T052 Implement default vim-style keybindings in internal/keymap/defaults.go (j/k/g/G/h/l/a/e/Space/d/+/-/0///f/c/r/s/q/Esc/?/F1)
- [X] T053 Implement Keymap implementation in internal/keymap/keymap_impl.go
- [ ] T054 Implement Handler interface for processing key events in internal/keymap/handler.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - View and Navigate Tasks (Priority: P1) 🎯 MVP

**Goal**: Users can view their todo.txt file as an interactive list and navigate using vim-style keys (j/k/g/G/h/l)

**Independent Test**: Launch app with sample todo.txt file, verify tasks load correctly, all navigation keys work (j/k/g/G/h/l), UI displays tasks with proper formatting and priority indicators

### Tests for User Story 1 (REQUIRED - Ginkgo BDD) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T055 [P] [US1] Write Ginkgo unit tests for TaskListModel in internal/ui/models/tasklist_model_test.go
- [ ] T056 [P] [US1] Write Ginkgo unit tests for task rendering in internal/ui/components/task_renderer_test.go
- [ ] T057 [P] [US1] Write Ginkgo integration test for loading todo.txt and displaying tasks in tests/integration/ui/view_tasks_test.go
- [ ] T058 [P] [US1] Write Ginkgo integration test for navigation (j/k/g/G) in tests/integration/ui/navigation_test.go

### Implementation for User Story 1

- [ ] T059 [P] [US1] Create TaskListModel interface in internal/ui/models/tasklist.go implementing bubbletea Model (Init, Update, View)
- [ ] T060 [P] [US1] Create task renderer component in internal/ui/components/task_renderer.go for formatting tasks with priority indicators using lipgloss
- [ ] T061 [P] [US1] Create viewport component wrapper in internal/ui/components/viewport.go for scroll management
- [ ] T062 [US1] Implement TaskListModel.Init method to load tasks from TodoFile in internal/ui/models/tasklist.go
- [ ] T063 [US1] Implement TaskListModel.Update method to handle navigation messages (MoveDown, MoveUp, MoveTop, MoveBottom) in internal/ui/models/tasklist.go
- [ ] T064 [US1] Implement TaskListModel.View method to render visible tasks in viewport in internal/ui/models/tasklist.go
- [ ] T065 [US1] Implement navigation methods (SelectNext, SelectPrev, SelectFirst, SelectLast) in internal/ui/models/tasklist.go
- [ ] T066 [US1] Implement scroll methods (ScrollDown, ScrollUp, ScrollTo) with viewport management in internal/ui/models/tasklist.go
- [ ] T067 [US1] Create main application model in internal/ui/app.go that coordinates TaskListModel and keymap
- [ ] T068 [US1] Implement main.go entry point in cmd/todotxt-tui/main.go that initializes config, loads todo.txt, starts bubbletea program
- [ ] T069 [US1] Add styling with lipgloss for priority indicators (A=red, B=yellow, C=green, etc.) in internal/ui/styles/styles.go
- [ ] T070 [US1] Add terminal resize handling in internal/ui/models/tasklist.go Update method
- [ ] T071 [US1] Implement empty state display when no tasks exist in internal/ui/components/empty_state.go
- [ ] T072 [US1] Integrate keymap for navigation keys (j/k/g/G) with TaskListModel in internal/ui/app.go
- [ ] T072a [US1] Implement detail view for displaying full untruncated task text when Enter key pressed in internal/ui/models/detail_view.go

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Create and Edit Tasks (Priority: P2)

**Goal**: Users can create new tasks and edit existing tasks with full todo.txt format support (priority, contexts, projects, dates)

**Independent Test**: Create new task with priority/context/project, edit existing task, verify persistence to file and proper display

### Tests for User Story 2 (REQUIRED - Ginkgo BDD) ⚠️

- [ ] T073 [P] [US2] Write Ginkgo unit tests for DialogModel in internal/ui/models/dialog_model_test.go
- [ ] T074 [P] [US2] Write Ginkgo unit tests for task creation logic in internal/ui/actions/create_task_test.go
- [ ] T075 [P] [US2] Write Ginkgo unit tests for task editing logic in internal/ui/actions/edit_task_test.go
- [ ] T076 [P] [US2] Write Ginkgo integration test for creating task with all format elements in tests/integration/ui/create_task_test.go
- [ ] T077 [P] [US2] Write Ginkgo integration test for editing task and persistence in tests/integration/ui/edit_task_test.go

### Implementation for User Story 2

- [ ] T078 [P] [US2] Create DialogModel interface in internal/ui/models/dialog.go implementing bubbletea Model with SetPrompt, SetValue, GetValue, Show, Hide, Confirm, Cancel methods
- [ ] T079 [P] [US2] Create text input component wrapper in internal/ui/components/text_input.go using bubbletea textinput
- [ ] T080 [P] [US2] Create priority selector component in internal/ui/components/priority_selector.go for A-Z selection
- [ ] T081 [US2] Implement DialogModel for task creation in internal/ui/models/create_dialog.go with fields for description, priority, context, project
- [ ] T082 [US2] Implement DialogModel for task editing in internal/ui/models/edit_dialog.go pre-populating with existing task values
- [ ] T083 [US2] Add message types (TaskAddedMsg, TaskEditedMsg) in internal/ui/messages.go
- [ ] T084 [US2] Implement AddTask action handler in internal/ui/actions/add_task.go that parses input and calls TodoFile.AddTask
- [ ] T085 [US2] Implement EditTask action handler in internal/ui/actions/edit_task.go that updates selected task
- [ ] T086 [US2] Integrate 'a' key binding to show create dialog in internal/ui/app.go
- [ ] T087 [US2] Integrate 'e' key binding to show edit dialog for selected task in internal/ui/app.go
- [ ] T088 [US2] Implement auto-save after task creation/edit in internal/ui/app.go using Storage.Save
- [ ] T089 [US2] Add creation date automatically to new tasks in internal/ui/actions/add_task.go
- [ ] T090 [US2] Implement tag auto-completion for contexts and projects in internal/ui/components/tag_input.go (extract existing tags from all tasks)
- [ ] T091 [US2] Add validation for task input (non-empty description) in internal/ui/models/create_dialog.go
- [ ] T092 [US2] Update TaskListModel to refresh after task add/edit in internal/ui/models/tasklist.go

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Complete and Delete Tasks (Priority: P2)

**Goal**: Users can mark tasks as complete (x + date) and delete tasks with confirmation

**Independent Test**: Mark task as complete, verify "x" prefix and date added, delete task with confirmation, verify removal from file

### Tests for User Story 3 (REQUIRED - Ginkgo BDD) ⚠️

- [ ] T093 [P] [US3] Write Ginkgo unit tests for task completion logic in internal/ui/actions/complete_task_test.go
- [ ] T094 [P] [US3] Write Ginkgo unit tests for task deletion logic in internal/ui/actions/delete_task_test.go
- [ ] T095 [P] [US3] Write Ginkgo unit tests for confirmation dialog in internal/ui/models/confirm_dialog_test.go
- [ ] T096 [P] [US3] Write Ginkgo integration test for completing task and verification in tests/integration/ui/complete_task_test.go
- [ ] T097 [P] [US3] Write Ginkgo integration test for deleting task with confirmation in tests/integration/ui/delete_task_test.go
- [ ] T098 [P] [US3] Write Ginkgo unit tests for archive logic in internal/storage/archive_test.go

### Implementation for User Story 3

- [ ] T099 [P] [US3] Create confirmation dialog model in internal/ui/models/confirm_dialog.go for delete confirmation
- [ ] T100 [P] [US3] Implement Task.Complete method to set Completed=true, CompletionDate=today in internal/parser/task.go
- [ ] T101 [P] [US3] Implement Task.Uncomplete method to toggle back to incomplete in internal/parser/task.go
- [ ] T102 [US3] Implement CompleteTask action handler in internal/ui/actions/complete_task.go that calls Task.Complete and TodoFile.UpdateTask
- [ ] T103 [US3] Implement DeleteTask action handler in internal/ui/actions/delete_task.go that calls TodoFile.DeleteTask after confirmation
- [ ] T104 [US3] Integrate Space key binding to toggle task completion in internal/ui/app.go
- [ ] T105 [US3] Integrate 'd' key binding to show delete confirmation dialog in internal/ui/app.go
- [ ] T106 [US3] Implement single-level undo for delete (store last deleted task) in internal/ui/actions/delete_task.go
- [ ] T107 [US3] Integrate 'u' key binding for undo delete in internal/ui/app.go
- [ ] T108 [US3] Implement archive support: move completed tasks to done.txt if ArchiveCompleted config enabled in internal/storage/archive.go
- [ ] T109 [US3] Add visual indicator for completed tasks (strikethrough or muted color) in internal/ui/components/task_renderer.go
- [ ] T110 [US3] Implement sorting completed tasks to bottom of list in internal/filter/sort.go
- [ ] T111 [US3] Add message types (TaskCompletedMsg, TaskDeletedMsg, TaskUndoMsg) in internal/ui/messages.go
- [ ] T112 [US3] Update TaskListModel to refresh and maintain selection after complete/delete in internal/ui/models/tasklist.go

**Checkpoint**: All core task management features (view, create, edit, complete, delete) are now functional

---

## Phase 6: User Story 4 - Filter and Search Tasks (Priority: P3)

**Goal**: Users can filter by priority/context/project and search by text to find relevant tasks quickly

**Independent Test**: Apply priority A filter (only A tasks shown), filter by @phone context, search for "meeting", combine filters with AND/OR logic

### Tests for User Story 4 (REQUIRED - Ginkgo BDD) ⚠️

- [ ] T113 [P] [US4] Write Ginkgo unit tests for FilterPanelModel in internal/ui/models/filter_panel_test.go
- [ ] T114 [P] [US4] Write Ginkgo unit tests for search functionality in internal/filter/search_test.go
- [ ] T115 [P] [US4] Write Ginkgo unit tests for filter combination (AND/OR) in internal/filter/combination_test.go
- [ ] T116 [P] [US4] Write Ginkgo integration test for filtering by priority in tests/integration/ui/filter_priority_test.go
- [ ] T117 [P] [US4] Write Ginkgo integration test for filtering by context/project in tests/integration/ui/filter_tags_test.go
- [ ] T118 [P] [US4] Write Ginkgo integration test for text search in tests/integration/ui/search_test.go
- [ ] T119 [P] [US4] Write Ginkgo integration test for combined filters in tests/integration/ui/combined_filters_test.go

### Implementation for User Story 4

- [ ] T120 [P] [US4] Create FilterPanelModel in internal/ui/models/filter_panel.go for filter UI with checkboxes/selectors
- [ ] T121 [P] [US4] Create search dialog model in internal/ui/models/search_dialog.go for text search input
- [ ] T122 [P] [US4] Create filter display component in internal/ui/components/filter_display.go showing active filters
- [ ] T123 [US4] Implement TaskList.ApplyFilters method in internal/filter/tasklist.go that uses Filter.Apply with criteria
- [ ] T124 [US4] Implement TaskList.AddFilter method in internal/filter/tasklist.go
- [ ] T125 [US4] Implement TaskList.RemoveFilter method in internal/filter/tasklist.go
- [ ] T126 [US4] Implement TaskList.ClearFilters method in internal/filter/tasklist.go
- [ ] T127 [US4] Integrate 'f' key binding to show filter panel in internal/ui/app.go
- [ ] T128 [US4] Integrate '/' key binding to show search dialog in internal/ui/app.go
- [ ] T129 [US4] Integrate 'c' key binding to clear all filters in internal/ui/app.go
- [ ] T130 [US4] Integrate 'n' and 'p' key bindings for next/previous priority task navigation in internal/ui/app.go
- [ ] T131 [US4] Add FilterChangedMsg message type in internal/ui/messages.go
- [ ] T132 [US4] Update TaskListModel to display filtered tasks in internal/ui/models/tasklist.go
- [ ] T133 [US4] Implement filter combination logic (AND vs OR) selection in filter panel in internal/ui/models/filter_panel.go
- [ ] T134 [US4] Add active filter indicators in header/status bar in internal/ui/components/status_bar.go
- [ ] T135 [US4] Implement filter suggestions based on existing tags in tasks in internal/ui/models/filter_panel.go
- [ ] T136 [US4] Add debouncing for search input to avoid lag in internal/ui/models/search_dialog.go

**Checkpoint**: All user stories (P1, P2, P3) are now independently functional

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

### Error Handling and Messages

- [ ] T137 [P] Create ErrorModel in internal/ui/models/error.go for displaying errors with SetError, Show, Hide methods
- [ ] T138 [P] Write Ginkgo unit tests for ErrorModel in internal/ui/models/error_test.go
- [ ] T139 Implement error display as non-blocking toast in internal/ui/components/error_toast.go
- [ ] T140 Add ErrorMsg message type in internal/ui/messages.go
- [ ] T141 Add error handling for file not found (create empty file) in internal/storage/file_storage.go
- [ ] T142 Add error handling for permission denied in internal/storage/file_storage.go
- [ ] T143 Add error handling for malformed task lines (display as-is with warning) in internal/parser/parser.go
- [ ] T144 Add error handling for terminal too small in internal/ui/app.go

### File Operations and External Changes

- [ ] T145 [P] Implement file watcher integration in internal/ui/app.go using Storage.Watch
- [ ] T146 [P] Write Ginkgo integration tests for external file changes in tests/integration/storage/file_watch_test.go
- [ ] T147 Integrate 'r' key binding for manual reload in internal/ui/app.go
- [ ] T148 Integrate 's' key binding for manual save in internal/ui/app.go
- [ ] T149 Add FileChangedMsg message type in internal/ui/messages.go
- [ ] T150 Implement reload prompt when external changes detected in internal/ui/models/reload_dialog.go
- [ ] T151 Add visual indicator when file has unsaved changes in internal/ui/components/status_bar.go

### Help and Documentation

- [ ] T152 [P] Create HelpModel in internal/ui/models/help.go displaying all keyboard shortcuts
- [ ] T153 [P] Write Ginkgo unit tests for HelpModel in internal/ui/models/help_test.go
- [ ] T154 Integrate '?' and 'F1' key bindings to show help screen in internal/ui/app.go
- [ ] T155 Add HelpRequestedMsg and HelpClosedMsg message types in internal/ui/messages.go
- [ ] T156 Create keyboard shortcuts reference in internal/ui/components/help_screen.go with categories (Navigation, Task Management, Filtering, File Operations, System)
- [ ] T157 [P] Create docs/architecture.md documenting component architecture and data flow

### Themes and Styling

- [ ] T158 [P] Implement light theme color scheme in internal/ui/styles/light.go
- [ ] T159 [P] Implement dark theme color scheme in internal/ui/styles/dark.go
- [ ] T160 [P] Implement terminal default theme in internal/ui/styles/default.go
- [ ] T161 Create theme selector based on config in internal/ui/styles/theme.go
- [ ] T162 Add terminal background color detection (future enhancement placeholder) in internal/ui/styles/detection.go

### Status Bar and UI Polish

- [ ] T163 [P] Create status bar component in internal/ui/components/status_bar.go showing task count, filter status, file path
- [ ] T164 Add timestamp display for last save in internal/ui/components/status_bar.go
- [ ] T165 Add progress indicator for loading large files in internal/ui/components/progress.go
- [ ] T166 Implement line truncation with visual indicator "…" for long task descriptions exceeding terminal width in internal/ui/components/task_renderer.go

### Performance Optimization

- [ ] T167 [P] Implement display string caching in Task struct in internal/parser/task.go
- [ ] T168 [P] Write benchmark tests for parsing 10,000 tasks in internal/parser/parser_bench_test.go
- [ ] T169 [P] Write benchmark tests for filtering 10,000 tasks in internal/filter/filter_bench_test.go
- [ ] T170 Implement lazy loading for files >1000 tasks with background goroutine in internal/storage/lazy_loader.go
- [ ] T171 Optimize index rebuilding to only rebuild on task changes in internal/filter/index.go
- [ ] T172 Profile application with pprof and optimize hot paths in cmd/todotxt-tui/main.go (add --profile flag)

### Edge Cases and Validation

- [ ] T173 [P] Handle empty task list gracefully (show empty state) in internal/ui/models/tasklist.go
- [ ] T174 [P] Handle task with multiple malformed priorities (use first valid or none) in internal/parser/parser.go
- [ ] T175 [P] Handle special characters and emoji in task descriptions in internal/parser/parser.go
- [ ] T176 Handle Esc key outside active operation (no-op, don't quit) in internal/ui/app.go
- [ ] T177 Handle increasing priority beyond A (stay at A) in internal/ui/actions/priority_actions.go
- [ ] T178 Handle decreasing priority beyond no priority (stay at none) in internal/ui/actions/priority_actions.go
- [ ] T179 Handle completing already completed task (toggle back to incomplete) in internal/ui/actions/complete_task.go
- [ ] T180 Handle deleting last remaining task in internal/ui/actions/delete_task.go
- [ ] T181 Handle editing task to empty string (reject with error) in internal/ui/actions/edit_task.go
- [ ] T182 [P] Write Ginkgo tests for all edge cases in internal/ui/edge_cases_test.go

### Build and Deployment

- [ ] T183 [P] Add cross-platform build targets to Makefile (linux, macos, windows)
- [ ] T184 [P] Create sample todo.txt files in tests/fixtures/todo/ for quickstart
- [ ] T185 [P] Update README.md with keyboard shortcuts, configuration, and quickstart instructions
- [ ] T186 [P] Add version flag and version display in cmd/todotxt-tui/main.go (--version)
- [ ] T187 [P] Add command-line flags for config file path, todo file path in cmd/todotxt-tui/main.go

### Code Quality and Testing

- [ ] T188 Run ginkgo -r --cover and verify 80% coverage threshold
- [ ] T189 Run golangci-lint run ./... and fix all issues
- [ ] T190 Run goimports -w . to format all code
- [ ] T191 Run go mod tidy to clean up dependencies
- [ ] T192 Verify quickstart.md instructions work end-to-end in tests/integration/quickstart_test.go
- [ ] T193 Run performance benchmarks and verify <100ms startup, <50MB memory targets
- [ ] T194 Test application with sample 1000-task and 10,000-task files from tests/fixtures/todo/large_*.txt

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Story 1 (Phase 3)**: Depends on Foundational phase completion
- **User Story 2 (Phase 4)**: Depends on Foundational phase completion (can run parallel to US1)
- **User Story 3 (Phase 5)**: Depends on Foundational phase completion (can run parallel to US1/US2)
- **User Story 4 (Phase 6)**: Depends on Foundational phase completion (can run parallel to US1/US2/US3)
- **Polish (Phase 7)**: Depends on desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: No dependencies on other stories - fully independent
- **User Story 2 (P2)**: No dependencies on other stories - integrates with US1 but independently testable
- **User Story 3 (P2)**: No dependencies on other stories - integrates with US1 but independently testable
- **User Story 4 (P3)**: No dependencies on other stories - integrates with US1 but independently testable

### Within Each User Story

- Tests MUST be written and FAIL before implementation (TDD - NON-NEGOTIABLE)
- Parser/Storage/Config before UI models
- UI models before actions
- Actions before keybinding integration
- Core implementation before edge cases
- Story complete before moving to next priority
- Verify test coverage meets 80% threshold per story

### Parallel Opportunities

- **Setup**: T003, T004, T005, T006, T007 (all different files)
- **Foundational Parser**: T008-T009, T010-T012, T016-T017, T018-T019 can run in parallel
- **Foundational Storage**: T020-T022, T028 can run in parallel with Parser tasks
- **Foundational Config**: T029-T030, T034 can run in parallel with Parser and Storage
- **Foundational Data**: T035-T036, T038-T039 can run in parallel
- **Foundational Filter**: T040-T043 can run in parallel
- **Foundational Keymap**: T049-T051 can run in parallel
- **User Story 1**: T055-T058 (all tests), T059-T061 (models/components) can run in parallel
- **User Story 2**: T073-T077 (all tests), T078-T080 (models/components) can run in parallel
- **User Story 3**: T093-T098 (all tests), T099-T101 (models/methods) can run in parallel
- **User Story 4**: T113-T119 (all tests), T120-T122 (models/components) can run in parallel
- **Polish**: Many polish tasks can run in parallel (T137-T138, T145-T146, T152-T153, T158-T160, T168-T169, T183-T187)
- **Once Foundational completes**: User Stories 1, 2, 3, 4 can ALL proceed in parallel by different developers

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together:
Task: "T055 [P] [US1] Write Ginkgo unit tests for TaskListModel in internal/ui/models/tasklist_model_test.go"
Task: "T056 [P] [US1] Write Ginkgo unit tests for task rendering in internal/ui/components/task_renderer_test.go"
Task: "T057 [P] [US1] Write Ginkgo integration test for loading todo.txt and displaying tasks in tests/integration/ui/view_tasks_test.go"
Task: "T058 [P] [US1] Write Ginkgo integration test for navigation (j/k/g/G) in tests/integration/ui/navigation_test.go"

# After tests are written, launch all models/components together:
Task: "T059 [P] [US1] Create TaskListModel interface in internal/ui/models/tasklist.go"
Task: "T060 [P] [US1] Create task renderer component in internal/ui/components/task_renderer.go"
Task: "T061 [P] [US1] Create viewport component wrapper in internal/ui/components/viewport.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001-T007)
2. Complete Phase 2: Foundational (T008-T054) - CRITICAL blocks all stories
3. Complete Phase 3: User Story 1 (T055-T072)
4. **STOP and VALIDATE**: Test User Story 1 independently
5. Deploy/demo if ready

**MVP Deliverable**: View and navigate tasks in TUI with vim keys - fully functional todo.txt viewer

### Incremental Delivery

1. Complete Setup + Foundational (T001-T054) → Foundation ready
2. Add User Story 1 (T055-T072) → Test independently → Deploy/Demo (MVP! Read-only viewer)
3. Add User Story 2 (T073-T092) → Test independently → Deploy/Demo (Task creation/editing)
4. Add User Story 3 (T093-T112) → Test independently → Deploy/Demo (Complete task lifecycle)
5. Add User Story 4 (T113-T136) → Test independently → Deploy/Demo (Full-featured with filtering)
6. Add Polish (T137-T195) → Final production release
7. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together (T001-T054)
2. Once Foundational is done:
   - Developer A: User Story 1 (T055-T072)
   - Developer B: User Story 2 (T073-T092)
   - Developer C: User Story 3 (T093-T112)
   - Developer D: User Story 4 (T113-T136)
3. Stories complete and integrate independently
4. Team collaborates on Polish phase (T137-T195)

---

## Task Summary

**Total Tasks**: 196
- **Phase 1 Setup**: 7 tasks (T001-T007)
- **Phase 2 Foundational**: 47 tasks (T008-T054)
- **Phase 3 User Story 1**: 19 tasks (T055-T072a)
- **Phase 4 User Story 2**: 20 tasks (T073-T092)
- **Phase 5 User Story 3**: 20 tasks (T093-T112)
- **Phase 6 User Story 4**: 24 tasks (T113-T136)
- **Phase 7 Polish**: 59 tasks (T137-T195)

**Parallel Opportunities**: ~80 tasks marked [P] can run in parallel with proper coordination

**MVP Scope (Recommended)**: Phases 1-3 (T001-T072a) = 73 tasks for basic working application

**Test Coverage**: 57 dedicated test tasks ensuring 80%+ coverage (constitution requirement)

---

## Notes

- [P] tasks = different files, no dependencies, can run in parallel
- [US1]/[US2]/[US3]/[US4] labels map task to specific user story for traceability
- Each user story is independently completable and testable
- Ginkgo BDD tests are NON-NEGOTIABLE (constitution requirement)
- Tests must be written FIRST and FAIL before implementation (TDD)
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that break independence
- All file paths are concrete and match plan.md structure
