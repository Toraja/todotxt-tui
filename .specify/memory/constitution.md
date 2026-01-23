# TodoTxt Constitution

<!--
Sync Impact Report:
Version change: Initial constitution → 1.0.0
Principles added:
  - I. Code Quality Standards
  - II. Test-Driven Development (NON-NEGOTIABLE)
  - III. User Experience Consistency
  - IV. Performance Requirements
  - V. Code Review & Documentation
Sections added:
  - Quality Gates & Enforcement
  - Development Workflow
  - Governance
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

**Test Naming Convention**: `test_[unit]_[scenario]_[expectedBehavior]`

**Rationale**: TDD ensures correctness by design, documents expected behavior, and enables fearless refactoring. Tests are living documentation that never goes stale.

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

## Quality Gates & Enforcement

All code changes MUST pass these automated gates before merge:

**Continuous Integration (CI) Requirements**:
1. **Linting**: Code style and quality checks MUST pass (zero warnings tolerated)
2. **Tests**: All tests MUST pass with required coverage thresholds
3. **Build**: Project MUST compile/build without errors
4. **Performance**: Benchmarks MUST not regress beyond threshold
5. **Security**: Dependency vulnerability scans MUST show no high/critical issues

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

**Version**: 1.0.0 | **Ratified**: 2026-01-23 | **Last Amended**: 2026-01-23
