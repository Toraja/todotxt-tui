# Quickstart Guide: todotxt-tui Development

**Feature**: todotxt-tui  
**Date**: 2026-02-02  
**Branch**: 001-todotxt-tui

## Prerequisites

- **Go 1.21+**: Install from https://go.dev/dl/
- **Git**: For version control
- **Terminal**: Any modern terminal with ANSI color support

## Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd todotxt-tui
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build the Application

```bash
make build
```

Or manually:

```bash
go build -o bin/todotxt-tui ./cmd/todotxt-tui
```

### 4. Run the Application

```bash
./bin/todotxt-tui
```

Or use make:

```bash
make run
```

The application will:
- Create a default todo.txt file at `~/.local/share/todotxt-tui/todo.txt` if it doesn't exist
- Display an empty task list
- Show help with `?` or `F1`

## Development Workflow

### Running Tests

```bash
# Run all tests with Ginkgo
make test

# Or manually
ginkgo -r --randomize-all --randomize-suites --fail-on-pending --cover

# Run tests for specific package
ginkgo ./internal/parser

# Run tests with verbose output
ginkgo -r -v
```

### Code Quality

Before committing, run:

```bash
make lint
```

This runs:
1. `go mod tidy` - Clean up dependencies
2. `goimports -w .` - Format code
3. `golangci-lint run` - Lint code (includes govet checks)
4. `ginkgo -r --fail-on-pending --cover` - Run tests

### Adding Features

1. Create a feature branch:
```bash
git checkout -b feature/your-feature-name
```

2. Write tests first (TDD):
```bash
ginkgo bootstrap
ginkgo generate ./internal/parser
```

3. Implement the feature
4. Run tests and fix failures
5. Run lint and fix issues
6. Commit and push

### Project Structure

```
.
├── cmd/
│   └── todotxt-tui/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration handling
│   ├── ui/                      # TUI components
│   │   ├── models/              # Bubbletea models
│   │   ├── components/          # Reusable UI components
│   │   └── views/               # Screen views
│   ├── parser/                  # todo.txt parsing
│   ├── filter/                  # Filtering and search
│   ├── storage/                 # File I/O
│   └── keymap/                  # Keyboard handling
├── tests/
│   ├── unit/                    # Unit tests
│   ├── integration/             # Integration tests
│   └── fixtures/                # Test data
├── specs/
│   └── 001-todotxt-tui/         # Feature specs and plans
│       ├── spec.md
│       ├── plan.md
│       ├── research.md
│       ├── data-model.md
│       ├── quickstart.md
│       └── contracts/
└── go.mod
```

## Common Tasks

### Adding a New Task

```go
// Example: Adding a task programmatically
task := &parser.Task{
    Priority:      "A",
    CreationDate:  time.Now(),
    Description:   "Buy groceries @store +household",
    Contexts:      map[string]struct{}{"@store", {}},
    Projects:      map[string]struct{}{"+household", {}},
}

// Add to TodoFile
todoFile.AddTask(task)
```

### Running with Custom Config

```bash
# Create config file
mkdir -p ~/.config/todotxt-tui
cat > ~/.config/todotxt-tui/config.yaml << EOF
todo_file_path: ~/todo.txt
theme: dark
show_completed: false
confirm_delete: true
EOF

# Run with config
./bin/todotxt-tui --config ~/.config/todotxt-tui/config.yaml
```

### Testing with Sample Data

```bash
# Create sample todo.txt
mkdir -p ~/.local/share/todotxt-tui
cat > ~/.local/share/todotxt-tui/todo.txt << EOF
(A) 2026-02-02 Buy groceries @store +household
(B) 2026-02-01 Call mom @family
(C) 2026-02-02 Review PR @work +project
x 2026-01-31 2026-02-01 Old completed task
EOF

# Run application
./bin/todotxt-tui
```

### Debugging

```bash
# Run with debug logging
DEBUG=1 ./bin/todotxt-tui

# Run with pprof for performance profiling
go run ./cmd/todotxt-tui --profile
# Then visit http://localhost:6060/debug/pprof/

# Check coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Keyboard Shortcuts

### Navigation
- `j` / `k` - Move down/up
- `g` / `G` - Jump to top/bottom
- `h` / `l` - Navigate left/right in dialogs

### Task Management
- `a` - Add new task
- `e` - Edit selected task
- `Space` - Toggle completion
- `d` - Delete task (confirm)
- `+` / `-` - Increase/decrease priority
- `0` - Remove priority

### Filtering
- `/` - Search tasks
- `f` - Open filter menu
- `c` - Clear all filters

### File Operations
- `r` - Reload file
- `s` - Save manually

### System
- `q` - Quit
- `Esc` - Cancel/back
- `?` or `F1` - Help

## Configuration

Configuration is loaded from (in order of priority):
1. Command-line flags
2. Environment variables
3. Config file: `~/.config/todotxt-tui/config.yaml`
4. Default values

### Environment Variables

```bash
export TODOTXT_FILE="~/todo.txt"
export TODOTXT_THEME="dark"
export TODOTXT_DEBUG="1"
```

### Config File Format (YAML)

```yaml
todo_file_path: "~/.local/share/todotxt-tui/todo.txt"
done_file_path: ""  # Optional, for archiving
theme: "dark"
show_completed: false
archive_completed: false
confirm_delete: true
auto_save: true
file_watch_enabled: true
```

## Testing Guide

### Writing Ginkgo Tests

```go
package parser_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "todotxt-tui/internal/parser"
)

var _ = Describe("Parser", func() {
    var p parser.Parser

    BeforeEach(func() {
        p = parser.NewParser()
    })

    Describe("ParseLine", func() {
        Context("with valid priority task", func() {
            It("parses priority correctly", func() {
                task, err := p.ParseLine("(A) 2026-02-02 Buy milk @store", 1)
                Expect(err).ToNot(HaveOccurred())
                Expect(task.Priority).To(Equal("A"))
                Expect(task.Description).To(Equal("Buy milk @store"))
                Expect(task.Contexts).To(ContainElement("@store"))
            })
        })

        Context("with completed task", func() {
            It("parses completed status correctly", func() {
                task, err := p.ParseLine("x 2026-01-31 2026-02-01 Old task", 1)
                Expect(err).ToNot(HaveOccurred())
                Expect(task.Completed).To(BeTrue())
                Expect(task.Priority).To(BeEmpty())
            })
        })
    })

    DescribeTable("Priority parsing",
        func(line string, expectedPriority string) {
            task, err := p.ParseLine(line, 1)
            Expect(err).ToNot(HaveOccurred())
            Expect(task.Priority).To(Equal(expectedPriority))
        },
        Entry("(A) priority", "(A) Task text", "A"),
        Entry("(B) priority", "(B) Task text", "B"),
        Entry("no priority", "Task text", ""),
    )
})
```

### Running Specific Tests

```bash
# Run specific test
ginkgo -focus="ParseLine" ./internal/parser

# Run tests matching pattern
ginkgo -focus="priority" ./internal/parser

# Skip tests
ginkgo -skip="completed" ./internal/parser
```

## Troubleshooting

### Application Won't Start

**Problem**: "Permission denied" error
```bash
# Check file permissions
ls -la ~/.local/share/todotxt-tui/

# Fix permissions
chmod 644 ~/.local/share/todotxt-tui/todo.txt
```

### Tasks Not Loading

**Problem**: Empty task list but file exists
```bash
# Check file content
cat ~/.local/share/todotxt-tui/todo.txt

# Validate file format
# Run parser in test mode
go test ./internal/parser -v -run TestParseFile
```

### Terminal Display Issues

**Problem**: Colors not displaying correctly
```bash
# Check terminal color support
echo $TERM

# Force 256-color mode
export TERM=xterm-256color

./bin/todotxt-tui
```

### Performance Issues

**Problem**: Slow with large task lists
```bash
# Profile the application
go build -o bin/todotxt-tui ./cmd/todotxt-tui
./bin/todotxt-tui --profile

# Analyze pprof output
go tool pprof bin/todotxt-tui cpu.prof
```

## Contributing

### Code Style

- Follow Effective Go guidelines
- Use `goimports` for formatting
- Use meaningful variable names
- Add comments for exported functions
- Keep functions small and focused

### Commit Messages

```
feat(parser): add support for metadata parsing

- Parse key:value pairs from task description
- Add tests for metadata extraction
- Update documentation

Fixes #123
```

### Pull Request Process

1. Update documentation
2. Add tests for new features
3. Ensure all tests pass
4. Run `make lint` and fix issues
5. Update CHANGELOG if applicable
6. Submit PR with clear description

## Useful Resources

- [todo.txt specification](https://github.com/todotxt/todo.txt)
- [bubbletea documentation](https://github.com/charmbracelet/bubbletea)
- [lipgloss documentation](https://github.com/charmbracelet/lipgloss)
- [Ginkgo documentation](https://onsi.github.io/ginkgo/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Project constitution](../../.specify/memory/constitution.md)

## Getting Help

- Check `?` or `F1` in the application for keyboard shortcuts
- Review feature specs in `specs/001-todotxt-tui/spec.md`
- Check architecture docs in `docs/architecture.md`
- Open an issue for bugs or feature requests

## Next Steps

After setting up:

1. Read the [feature specification](spec.md) for full requirements
2. Review the [data model](data-model.md) for entity definitions
3. Check [API contracts](contracts/api.md) for component interfaces
4. Explore [research findings](research.md) for technical decisions
5. Start implementing from [tasks.md](tasks.md) (created in Phase 2)

Happy hacking! 🚀
