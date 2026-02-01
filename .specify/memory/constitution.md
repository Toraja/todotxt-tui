<!--
Sync Impact Report:
- Version change: 1.0.0 → 1.0.1
- Modified principles: None
- Modified sections: Development Workflow (gofmt → goimports)
- Added sections: None
- Removed sections: None
- Templates requiring updates:
  - ✅ All templates remain compatible
- Follow-up TODOs: None
-->

# todotxt-tui Constitution

## Core Principles

### I. Go Code Quality Standards

All code MUST follow Go best practices and idiomatic patterns:
- Run `goimports -w .` on all Go files before committing
- Run `golangci-lint run` and fix all reported issues
- Run `go vet ./...` before commits
- Follow Effective Go guidelines (https://golang.org/doc/effective_go)
- Use clear, descriptive names; avoid abbreviations unless widely known
- Export functions and types only when necessary; prefer package-private for internal implementation
- Handle errors explicitly; NEVER ignore errors with `\_`
- Use `const` for compile-time constants, NEVER magic numbers or strings
- Document exported packages, types, functions, and constants with godoc comments
- Structure packages by domain logic, not by architectural layers

**Rationale**: Go's simplicity and tooling ecosystem enable high-quality, maintainable code when conventions are followed consistently. These standards ensure code readability, catch common mistakes early, and align with the broader Go community.

### II. Testing with Ginkgo (NON-NEGOTIABLE)

All code MUST be tested using the Ginkgo BDD testing framework:
- Use Ginkgo's Describe/Context/It structure for test organization
- Write table-driven tests using Ginkgo's `DescribeTable` or `Entry` for multiple scenarios
- Maintain minimum 80% code coverage (measured by `go test -cover`)
- Write tests BEFORE implementation (TDD approach)
- Each test MUST be independent and runnable in any order
- Use Ginkgo's `BeforeEach` and `AfterEach` for setup/teardown, NEVER `Setup`/`TeardownSuite` except for global resources
- Mock external dependencies using `gomock` or test doubles
- Test both happy paths and error cases
- Integration tests MUST use actual file I/O where appropriate for todotxt format compliance
- Run full test suite with `ginkgo -r --randomize-all --randomize-suites --fail-on-pending --cover`

**Rationale**: Ginkgo's BDD style produces readable, maintainable tests that serve as living documentation. High test coverage ensures reliability of TUI operations and file parsing. Testing-first prevents bugs and clarifies requirements before implementation.

### III. TUI User Experience Consistency

All TUI interactions MUST provide a consistent, predictable user experience:
- Use vim-style keyboard navigation: `j`/`k` for down/up, `h`/`l` for left/right, `g`/`G` for top/bottom
- Standardized action keys: `Enter` to select/confirm, `Esc` to cancel/back, `q` to quit
- Consistent color scheme: use a defined palette for all UI elements (success, warning, error, muted, accent)
- Handle terminal resize events gracefully; preserve state and redraw appropriately
- Display clear, actionable error messages; include context and suggested fixes when possible
- Provide keyboard shortcuts for all common actions; display them in help text
- Support mouse input where appropriate but NEVER require it
- Maintain cursor visibility and positioning; never leave cursor in unexpected locations
- Use loading indicators for operations >200ms
- Ensure accessibility with clear visual hierarchy and sufficient contrast

**Rationale**: TUI users expect consistency across terminal applications. Following established conventions (especially vim-style navigation) reduces cognitive load and makes the application intuitive. Proper error handling and responsiveness build user trust.

### IV. Performance Requirements

All code MUST meet performance standards to ensure responsive TUI operation:
- Application startup time MUST be <100ms for typical todo.txt files (<1000 tasks)
- Memory usage MUST stay <50MB during normal operation
- UI rendering MUST maintain 60fps for smooth scrolling and updates
- File loading and saving operations MUST complete in <50ms for files up to 1000 tasks
- All user input MUST be acknowledged within 16ms (one frame at 60fps)
- Never block the main event loop; run heavy I/O or computation in goroutines
- Use efficient data structures: prefer slices for dynamic collections, maps for lookups
- Avoid unnecessary allocations in hot paths; profile with `pprof` to identify bottlenecks
- Cache computed results where appropriate; invalidate cache on data changes
- Implement progressive rendering for large task lists (render visible region first)

**Rationale**: TUI applications must feel snappy and responsive to be usable. Performance issues cause user frustration and abandonment. These standards ensure smooth operation even on resource-constrained systems and large todo.txt files.

## Performance Standards

Detailed performance benchmarks that MUST be maintained:

**Startup Performance**:
- Cold start (no cache): <100ms for 100 tasks, <200ms for 1000 tasks
- Warm start (with cache): <50ms for any file size
- Measure with `time todotxt-tui --version` and actual file loads

**Memory Usage**:
- Baseline (no tasks): <10MB
- 100 tasks: <20MB
- 1000 tasks: <50MB
- 10000 tasks: <100MB (progressive rendering required)
- Profile with `go tool pprof -http=:8080`

**Rendering Performance**:
- Frame time (single screen refresh): <16ms (60fps target)
- Scroll by one line: <5ms
- Redraw entire screen: <10ms
- Search/filter on 1000 tasks: <20ms

**File I/O Performance**:
- Parse 1000 tasks: <30ms
- Save 1000 tasks: <20ms
- Watch file change detection: <50ms latency

**Rationale**: Quantifiable performance benchmarks allow regression testing and ensure the application remains performant as features are added. Use Go's built-in `testing` package benchmarks and `pprof` for validation.

## Development Workflow

All development MUST follow this workflow to ensure quality:

**Before Committing**:
1. Run `go mod tidy` to ensure dependencies are clean
2. Run `goimports -w .` to format all code
3. Run `golangci-lint run` and fix all issues
4. Run `go vet ./...` and fix all warnings
5. Run `ginkgo -r --randomize-all --fail-on-pending --cover` and ensure all tests pass
6. Verify test coverage meets 80% threshold

**Branching Strategy**:
- Use feature branches: `feature/short-description`
- Branch names MUST be lowercase with hyphens
- Keep branches small and focused; 3-5 commits maximum
- Rebase main branch before creating pull requests

**Code Review Requirements**:
- All code MUST be reviewed before merging
- Reviewer must verify constitution compliance
- At least one approval required for feature branches
- Reviewer must run full test suite on their machine

**Continuous Integration**:
- CI pipeline MUST run on every pull request
- CI MUST execute: `go mod tidy`, `goimports`, `golangci-lint`, `go vet`, `ginkgo -r`
- CI MUST measure and report test coverage
- CI MUST fail if any step fails
- CI MUST include performance benchmarks for critical paths

**Rationale**: A consistent workflow prevents bugs, maintains code quality, and ensures all code meets the constitution standards. Automated checks in CI enforce these standards consistently.

## Governance

This constitution supersedes all other development practices and guidelines. All code and documentation MUST comply with these principles.

**Amendment Procedure**:
- Amendments require documented rationale and impact analysis
- Proposed amendments must be reviewed and approved by project maintainers
- Amendments MUST update `CONSTITUTION_VERSION` according to semantic versioning:
  - MAJOR: Backward incompatible principle removals or redefinitions
  - MINOR: New principle or section added, or materially expanded guidance
  - PATCH: Clarifications, wording fixes, non-semantic refinements
- Update `LAST_AMENDED_DATE` on all amendments
- Migration plans required for MAJOR version changes

**Compliance Review**:
- All pull requests MUST verify compliance with relevant principles
- Complexity beyond these principles must be explicitly justified in pull request description
- Use `.specify/memory/constitution.md` as the single source of truth for development standards

**Versioning Policy**:
- Follow semantic versioning: MAJOR.MINOR.PATCH
- Increment version based on amendment type (see Amendment Procedure)
- Document version changes in the Sync Impact Report

**Rationale**: A clear governance process ensures the constitution evolves deliberately while maintaining stability. Versioning tracks changes and communicates impact to developers.

**Version**: 1.0.1 | **Ratified**: 2026-02-01 | **Last Amended**: 2026-02-01
