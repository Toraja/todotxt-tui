# todotxt-tui

A terminal user interface (TUI) application for managing [todo.txt](https://github.com/todotxt/todo.txt) files with vim-style keyboard navigation.

## Features

- **Vim-style Navigation**: Navigate tasks with familiar j/k/g/G keys
- **Full todo.txt Support**: Complete support for priority, contexts, projects, dates, and metadata
- **Filtering & Search**: Filter tasks by priority, context, project, or search text
- **Live Editing**: Create, edit, complete, and delete tasks with instant feedback
- **File Watching**: Auto-reload when file changes externally
- **High Performance**: Efficiently handles files with 10,000+ tasks
- **Customizable**: Configuration file support for themes and preferences

## Installation

### Prerequisites

- Go 1.21+ (or latest stable)
- Make (optional, for build targets)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/Toraja/todotxt-tui.git
cd todotxt-tui

# Install dependencies
go mod download

# Build the application
make build

# Or build manually
go build -o bin/todotxt-tui ./cmd/todotxt-tui
```

## Usage

```bash
# Run with default todo.txt location (~/.local/share/todotxt-tui/todo.txt)
./bin/todotxt-tui

# Or use make
make run

# Specify custom todo.txt file
./bin/todotxt-tui --todo-file ~/Documents/todo.txt

# Use custom config file
./bin/todotxt-tui --config ~/.config/todotxt-tui/config.yaml
```

## Keyboard Shortcuts

### Navigation
- `j` / `k` - Move down/up
- `g` / `G` - Jump to top/bottom
- `h` / `l` - Navigate left/right in dialogs
- `Enter` - Show task detail view

### Task Management
- `a` - Add new task
- `e` - Edit selected task
- `Space` - Toggle completion
- `d` - Delete task (with confirmation)
- `+` / `-` - Increase/decrease priority
- `0` - Remove priority
- `u` - Undo last delete

### Filtering & Search
- `/` - Search tasks
- `f` - Open filter menu
- `c` - Clear all filters
- `n` / `p` - Next/previous priority task

### File Operations
- `r` - Reload file
- `s` - Save manually

### System
- `q` - Quit
- `Esc` - Cancel/back
- `?` or `F1` - Help

## Configuration

Configuration file location: `~/.config/todotxt-tui/config.yaml`

Example configuration:

```yaml
todo_file_path: "~/.local/share/todotxt-tui/todo.txt"
done_file_path: ""  # Optional, for archiving completed tasks
theme: "dark"       # Options: light, dark, default
show_completed: false
archive_completed: false
confirm_delete: true
auto_save: true
file_watch_enabled: true
```

## Development

### Running Tests

```bash
# Run all tests with Ginkgo
make test

# Run specific package tests
ginkgo ./internal/parser

# Run tests with verbose output
ginkgo -r -v
```

### Code Quality

```bash
# Run linters and formatters
make lint

# Install development tools
make install-tools
```

### Project Structure

```
.
├── cmd/
│   └── todotxt-tui/        # Main application entry point
├── internal/
│   ├── config/             # Configuration handling
│   ├── ui/                 # TUI components
│   │   ├── models/         # Bubbletea models
│   │   ├── components/     # Reusable UI components
│   │   └── views/          # Screen views
│   ├── parser/             # todo.txt parsing
│   ├── filter/             # Filtering and search
│   ├── storage/            # File I/O and persistence
│   └── keymap/             # Keyboard handling
├── tests/                  # Unit tests are adjacent to source files (e.g., internal/parser/parser_test.go)
│   ├── integration/        # Integration tests
│   └── fixtures/           # Test data
└── docs/                   # Documentation
```

## Contributing

See [quickstart.md](specs/001-todotxt-tui/quickstart.md) for development setup and guidelines.

## License

See LICENSE file for details.

## Acknowledgments

- [todo.txt format](https://github.com/todotxt/todo.txt)
- [Charm - Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Charm - Lipgloss](https://github.com/charmbracelet/lipgloss)
