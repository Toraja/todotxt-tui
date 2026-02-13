package filter

import "github.com/Toraja/todotxt-tui/internal/parser"

// FilterType represents the type of filter to apply.
type FilterType int

const (
	FilterNone FilterType = iota
	FilterPriority
	FilterContext
	FilterProject
	FilterSearch
	FilterCompleted
)

// FilterLogic defines how multiple filters are combined.
type FilterLogic int

const (
	FilterAnd FilterLogic = iota // All filters must match
	FilterOr                     // Any filter can match
)

// FilterCriteria represents a single filter criterion.
type FilterCriteria struct {
	Type    FilterType  // Type of filter
	Value   string      // Filter value (priority letter, @context, +project, search text)
	Enabled bool        // True if filter is active
	Logic   FilterLogic // AND/OR for combining with other filters
}

// Matches checks if a task matches this filter criterion.
func (fc *FilterCriteria) Matches(task *parser.Task) bool {
	if !fc.Enabled {
		return true // Disabled filters match everything
	}

	switch fc.Type {
	case FilterNone:
		return true

	case FilterPriority:
		return task.Priority == fc.Value

	case FilterContext:
		return task.HasContext(fc.Value)

	case FilterProject:
		return task.HasProject(fc.Value)

	case FilterSearch:
		// Case-insensitive substring search in description
		// TODO: Support regex in future
		return containsIgnoreCase(task.Description, fc.Value)

	case FilterCompleted:
		if fc.Value == "true" {
			return task.IsComplete()
		}
		return !task.IsComplete()

	default:
		return false
	}
}

// String returns a human-readable description of the filter.
func (fc *FilterCriteria) String() string {
	if !fc.Enabled {
		return "disabled"
	}

	switch fc.Type {
	case FilterNone:
		return "no filter"
	case FilterPriority:
		return "priority:" + fc.Value
	case FilterContext:
		return "context:" + fc.Value
	case FilterProject:
		return "project:" + fc.Value
	case FilterSearch:
		return "search:" + fc.Value
	case FilterCompleted:
		if fc.Value == "true" {
			return "completed"
		}
		return "active"
	default:
		return "unknown"
	}
}

// Equals checks if two filters are equivalent.
func (fc *FilterCriteria) Equals(other *FilterCriteria) bool {
	return fc.Type == other.Type &&
		fc.Value == other.Value &&
		fc.Enabled == other.Enabled &&
		fc.Logic == other.Logic
}

// containsIgnoreCase performs case-insensitive substring search.
func containsIgnoreCase(s, substr string) bool {
	// Simple implementation - convert both to lowercase
	// TODO: Use more efficient algorithm for large strings
	sLower := toLower(s)
	substrLower := toLower(substr)
	return contains(sLower, substrLower)
}

// toLower converts string to lowercase (simple ASCII version).
func toLower(s string) string {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}

// contains checks if s contains substr.
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(substr) > len(s) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
