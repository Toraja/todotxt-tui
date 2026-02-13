# Implementation Plan: todotxt-tui

**Branch**: `001-todotxt-tui` | **Date**: 2026-02-02 | **Spec**: spec.md
**Input**: Feature specification from `/specs/001-todotxt-tui/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

Develop a terminal user interface (TUI) application for managing todo.txt files with vim-style keyboard navigation. The application will parse and display tasks, support creation/editing/completion/deletion, provide filtering and search capabilities, and maintain high performance with files containing 10,000+ tasks. Technical approach uses Go with bubbletea TUI framework, lipgloss for styling, and Ginkgo for BDD testing.

## Technical Context

**Language/Version**: Go 1.21+ (or latest stable)
**Primary Dependencies**: bubbletea (TUI framework), lipgloss (styling), ginkgo (testing), gomega (matchers)
**Storage**: todo.txt text files (local filesystem at ~/.local/share/todotxt-tui/todo.txt)
**Testing**: Ginkgo BDD framework + gomega matchers
**Target Platform**: Linux, macOS, Windows (terminal-based)
**Project Type**: Single TUI application
**Performance Goals**: <100ms startup, <50MB memory, 60fps rendering, <50ms file I/O
**Constraints**: Offline-capable, no blocking UI operations, responsive to terminal resize
**Scale/Scope**: Support 10,000+ tasks with progressive rendering

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify compliance with todotxt-tui Constitution:

- [x] **Go Code Quality**: Design follows Go idioms with minimal dependencies (bubbletea, lipgloss, ginkgo). Uses standard library for file I/O and parsing.
- [x] **Testing with Ginkgo**: All components (parser, UI models, file storage, filters) are testable with Ginkgo BDD. 80% coverage achievable for core functionality.
- [x] **TUI UX Consistency**: Design follows vim-style navigation (j/k, g/G, h/l) as specified. Standard TUI conventions (Enter/Esc/q) included.
- [x] **Performance Requirements**: Design meets <100ms startup (efficient parsing), <50MB memory (lazy loading), 60fps rendering (bubbletea event loop), <50ms file I/O.

No violations detected.

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)
<!--
  ACTION REQUIRED: Replace the placeholder tree below with the concrete layout
  for this feature. Delete unused options and expand the chosen structure with
  real paths (e.g., apps/admin, packages/something). The delivered plan must
  not include Option labels.
-->

```text
cmd/
└── todotxt-tui/          # Main application entry point
    └── main.go

internal/
├── config/               # Configuration and settings
├── ui/                   # TUI components and screens
│   ├── models/           # Bubbletea models
│   ├── components/       # Reusable UI components
│   └── views/            # Screen views
├── parser/               # todo.txt parsing logic
├── filter/               # Filtering and searching
├── storage/              # File I/O and persistence
└── keymap/               # Keyboard handling and shortcuts

tests/
├── integration/          # Integration tests with Ginkgo
└── fixtures/             # Test data files

docs/
└── architecture.md

go.mod
go.sum
Makefile                 # Build and test commands
```

**Structure Decision**: Standard Go project layout with cmd/ for entry point, internal/ for private packages, and tests/ co-located with implementation following Go conventions.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |
