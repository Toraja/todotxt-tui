# TodoTxt Constitution

<!--
Sync Impact Report:
Version change: 1.0.0 → 1.1.0
Principles added:
  - I. Code Quality Standards
  - II. Test-Driven Development (NON-NEGOTIABLE)
  - III. User Experience Consistency
  - IV. Performance Requirements
  - V. Code Review & Documentation
  - VI. Language-Specific Standards: Golang (NEW in 1.1.0)
Sections added:
  - Quality Gates & Enforcement
  - Development Workflow
  - Governance
Updates in 1.1.0:
  - Added Golang-specific conventions and standards
  - Updated test requirements to mandate Ginkgo testing framework
  - Added Go module and dependency management requirements
Templates status:
  - ✅ plan-template.md (to be validated)
  - ✅ spec-template.md (to be validated)
  - ✅ tasks-template.md (to be validated)
Follow-up TODOs: None
-->

## Core Principles

### I. Code Quality Standards

**Maintainability is paramount.** All code MUST adhere to the following non-negotiable standards:

- **Single Responsibility**: Each function, class, or module serves ONE clear purpose
- **DRY (Don't Repeat Yourself)**: Extract common logic into reusable components; no code duplication beyond 3 occurrences
- **Clear Naming**: Variables, functions, and types use descriptive names that reveal intent without requiring comments
- **Complexity Limits**: 
  - Functions MUST NOT exceed 50 lines (excluding comments/whitespace)
  - Cyclomatic complexity MUST stay below 10
  - Nesting depth MUST NOT exceed 3 levels
- **Type Safety**: Use strong typing; no implicit any/unknown types without explicit justification
- **Error Handling**: All error paths explicitly handled; no silent failures

**Rationale**: Code is read 10x more than written. Prioritizing readability and maintainability reduces technical debt and accelerates feature velocity over time.

### II. Test-Driven Development (NON-NEGOTIABLE)

**TDD is mandatory for all feature development.** The Red-Green-Refactor cycle MUST be strictly enforced:

1. **Red**: Write failing tests that define expected behavior
2. **Green**: Implement minimal code to make tests pass
3. **Refactor**: Improve code quality while keeping tests green

**Test Coverage Requirements**:
- **Minimum 80% code coverage** for all production code
- **100% coverage required** for:
  - Business logic and algorithms
  - Data transformation functions
  - Error handling paths
  - Public APIs and interfaces
- **Unit tests MUST**:
  - Run in isolation (no external dependencies)
  - Execute in under 100ms each
  - Test one behavior per test case
- **Integration tests required** for:
  - File I/O operations
  - External service interactions
  - Database operations
  - Cross-module contracts

**Test Framework Requirements**:
- **All tests MUST use the Ginkgo testing framework** with Gomega matchers
- Tests MUST follow BDD (Behavior-Driven Development) style with Describe/Context/It blocks
- Use descriptive test descriptions that read as specifications
- Example structure:
  ```go
  Describe("ComponentName", func() {
      Context("when specific condition", func() {
          It("should exhibit expected behavior", func() {
              // test implementation
          })
      })
  })
  ```

**Test File Naming Convention**: `[filename]_test.go` (Go standard)

**Rationale**: TDD ensures correctness by design, documents expected behavior, and enables fearless refactoring. Tests are living documentation that never goes stale. Ginkgo provides expressive BDD-style testing that improves test readability and maintainability.

### III. User Experience Consistency

**User interfaces MUST be predictable, intuitive, and accessible.** Every user-facing feature MUST adhere to:

**CLI Interface Standards**:
- **Input Protocol**: Accept data via stdin, command-line arguments, or config files
- **Output Protocol**: Results to stdout, errors to stderr
- **Format Support**: Provide both human-readable and machine-readable (JSON) outputs
- **Flag Consistency**: Use standard flags (-h/--help, -v/--version, -q/--quiet, -o/--output)
- **Exit Codes**: Use conventional exit codes (0=success, 1=general error, 2=misuse)

**Error Messages**:
- **Descriptive**: Clearly state what went wrong and why
- **Actionable**: Provide specific steps to resolve the issue
- **Contextual**: Include relevant details (file paths, line numbers, values)
- **User-Friendly**: Avoid technical jargon in user-facing messages

**Documentation Standards**:
- **Every public function/API** documented with purpose, parameters, return values, and examples
- **User guides** for all features with real-world examples
- **Change logs** maintained for all releases

**Rationale**: Consistency reduces cognitive load, improves productivity, and lowers the barrier to entry for new users. Predictable behavior builds user trust and confidence.

### IV. Performance Requirements

**Performance is a feature, not an afterthought.** All code MUST meet these performance standards:

**Response Time Targets**:
- **CLI commands**: Complete within 500ms for typical workloads
- **File operations**: Process files up to 10MB within 1 second
- **Search operations**: Return results within 200ms for datasets up to 10,000 items
- **Interactive operations**: Provide feedback within 100ms (perceived instant)

**Resource Constraints**:
- **Memory**: Peak usage MUST NOT exceed 256MB for typical workloads
- **CPU**: Avoid blocking operations; use asynchronous I/O where appropriate
- **Disk I/O**: Minimize file system operations; batch reads/writes when possible

**Performance Testing Requirements**:
- **Benchmark tests** for all performance-critical paths
- **Load tests** simulating realistic usage patterns
- **Profiling** before optimizing; measure, don't guess
- **Regression detection**: CI MUST fail if performance degrades by >10%

**Optimization Guidelines**:
- Profile first, optimize second
- Prioritize algorithmic improvements over micro-optimizations
- Document all performance trade-offs and decisions
- Use lazy evaluation and streaming for large datasets

**Rationale**: Performance directly impacts user satisfaction and adoption. Slow tools are abandoned tools. Setting clear performance budgets prevents degradation over time.

### V. Code Review & Documentation

**All changes MUST be reviewed and comprehensively documented.**

**Code Review Requirements**:
- **No direct commits to main**: All changes via pull requests
- **Minimum one approval** from a maintainer before merge
- **Review Checklist**:
  - Tests included and passing
  - Constitution compliance verified
  - Performance impact assessed
  - Documentation updated
  - No unnecessary complexity introduced

**Documentation Requirements**:
- **Inline comments** for non-obvious logic and algorithms
- **Architecture Decision Records (ADRs)** for significant design choices
- **API documentation** auto-generated from code comments
- **Examples** for all public APIs and features

**Rationale**: Peer review catches defects early, shares knowledge across the team, and maintains quality standards. Documentation ensures knowledge is preserved and accessible.

### VI. Language-Specific Standards: Golang

**This is a Golang project.** All code MUST strictly adhere to Go conventions and best practices:

**Go Coding Standards**:
- **Follow [Effective Go](https://go.dev/doc/effective_go)** guidelines
- **Use `gofmt` and `goimports`**: All code MUST be formatted with standard Go tools
- **Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)** conventions
- **Package Naming**: 
  - Use short, concise, lowercase package names
  - No underscores or mixedCaps in package names
  - Package name should be singular (e.g., `todo` not `todos`)
- **File Organization**:
  - One package per directory
  - Group related functionality in the same package
  - Keep internal implementation details in `internal/` packages
- **Exported vs Unexported**:
  - Only export what needs to be public (starts with uppercase)
  - Keep implementation details unexported (starts with lowercase)
  - Document all exported functions, types, and constants

**Go-Specific Code Quality**:
- **Error Handling**: 
  - Check ALL errors; never ignore `err` return values
  - Use `errors.Is()` and `errors.As()` for error comparison
  - Wrap errors with context using `fmt.Errorf("context: %w", err)`
  - Return errors rather than panicking (panic only for truly exceptional cases)
- **Interfaces**:
  - Accept interfaces, return concrete types
  - Keep interfaces small and focused (prefer many small interfaces)
  - Define interfaces where they are used, not where they are implemented
- **Struct Embedding**: Use composition over inheritance via struct embedding
- **Goroutines & Concurrency**:
  - Document when functions are safe for concurrent use
  - Use channels to communicate, avoid shared memory where possible
  - Always provide a way to stop goroutines (context cancellation)
- **Context Usage**: 
  - First parameter in functions should be `context.Context` when needed
  - Pass context through call chains, don't store in structs
- **Testing**:
  - **All tests MUST use `_test` package** (e.g., `package todo_test` for testing `package todo`)
  - This enforces testing through the public API and prevents testing internal implementation details
  - Prefer Ginkgo's BDD style with separate It blocks for different scenarios
  - Table-driven tests are acceptable ONLY when appropriate (e.g., argument validation, parsing multiple input formats)
  - When using table-driven tests with Ginkgo, use `DescribeTable` and `Entry` constructs
  - Use testdata/ directory for test fixtures

**Dependency Management**:
- **Use Go modules** (`go.mod` and `go.sum`)
- Pin dependencies to specific versions
- Run `go mod tidy` regularly to clean up dependencies
- Avoid unnecessary external dependencies; prefer standard library
- Document why each major dependency is needed

**Go Tooling Requirements**:
- **Linting**: Use `golangci-lint` with strict configuration
- **Formatting**: `gofmt` and `goimports` MUST pass
- **Vet**: `go vet` MUST pass with zero warnings
- **Static Analysis**: Use `staticcheck` for additional checks
- **Security**: Run `gosec` for security vulnerability scanning

**Testing with Ginkgo**:
- **All tests MUST use Ginkgo BDD framework** with Gomega matchers
- Run tests with: `ginkgo -r` (recursive) or `go test ./...`
- Use `ginkgo bootstrap` to set up test suites
- Use `ginkgo generate` to create test files
- Leverage Gomega's expressive matchers (e.g., `Expect().To()`, `Eventually()`, `Consistently()`)
- Use `BeforeEach` and `AfterEach` for test setup/teardown
- Group related tests using `Context` blocks

**Documentation**:
- Every exported identifier MUST have a doc comment
- Doc comments start with the name of the identifier
- Use complete sentences in doc comments
- Example:
  ```go
  // Task represents a todo item with its metadata.
  // It implements the Stringer interface for easy printing.
  type Task struct { ... }
  ```

**Rationale**: Go has strong conventions that make code predictable and maintainable across the ecosystem. Following these standards ensures compatibility with Go tooling, improves code review efficiency, and makes the codebase accessible to any Go developer. Ginkgo provides a mature, expressive BDD testing framework that enhances test clarity and maintainability.

## Quality Gates & Enforcement

All code changes MUST pass these automated gates before merge:

**Continuous Integration (CI) Requirements**:
1. **Linting**: Code style and quality checks MUST pass (zero warnings tolerated)
   - For Go: `golangci-lint`, `go vet`, `gofmt`, `goimports`, `staticcheck`
2. **Tests**: All tests MUST pass with required coverage thresholds
   - Run with: `ginkgo -r --cover` or `go test -v -cover ./...`
3. **Build**: Project MUST compile/build without errors
   - Go: `go build ./...` MUST succeed
4. **Performance**: Benchmarks MUST not regress beyond threshold
   - Go: Run `go test -bench=.` for performance tests
5. **Security**: Dependency vulnerability scans MUST show no high/critical issues
   - Go: Use `gosec` and `go list -json -m all | nancy sleuth`

**Pre-commit Hooks**:
- Format code automatically
- Run fast unit tests
- Check for common errors (TODOs without issues, hardcoded secrets)

**Branch Protection**:
- Require CI passing before merge
- Require code review approval
- Require up-to-date with base branch
- No force pushes to protected branches

## Development Workflow

**Standard Development Process**:

1. **Planning**: Create specification using `/speckit.specify` before implementation
2. **Task Breakdown**: Decompose feature into testable units using `/speckit.tasks`
3. **Test Writing**: Write failing tests that define expected behavior
4. **Implementation**: Write minimal code to make tests pass
5. **Refactoring**: Improve code quality while maintaining green tests
6. **Documentation**: Update user docs and inline comments
7. **Review**: Submit PR with checklist completed
8. **Merge**: After approval and CI passing

**Branch Naming Convention**: `type/short-description`
- Types: feature, fix, refactor, docs, test, perf

**Commit Message Format**:
```
type(scope): brief description

- Detailed explanation if needed
- Reference to related issues/specs

Refs: #123
```

**Types**: feat, fix, refactor, docs, test, perf, chore

## Governance

**This constitution supersedes all other development practices and guidelines.**

**Amendment Process**:
1. Propose change with rationale in issue or ADR
2. Discuss and refine with team
3. Require consensus (or majority for non-critical changes)
4. Document decision in constitution
5. Update version following semantic versioning
6. Create migration plan if backward-incompatible
7. Communicate changes to all contributors

**Version Bump Rules**:
- **MAJOR**: Backward-incompatible principle changes or removals
- **MINOR**: New principle/section added or significant expansion
- **PATCH**: Clarifications, wording fixes, non-semantic improvements

**Compliance Review**:
- All pull requests MUST verify constitution compliance
- Complexity MUST be justified with concrete rationale
- Violations may be accepted with documented technical debt and remediation plan

**Constitution Authority**:
- Overrides style guides, coding standards, and process documents
- Defines minimum quality bar for all contributions
- Violations in existing code should be tracked and remediated incrementally

**Runtime Development Guidance**: See `.specify/templates/` for detailed templates and workflows aligned with these principles.

**Version**: 1.1.0 | **Ratified**: 2026-01-23 | **Last Amended**: 2026-01-23
