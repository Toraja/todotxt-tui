package parser

import (
	"fmt"
	"time"
)

// ParseTodoDate parses a date string in YYYY-MM-DD format.
// Returns zero time if the date string is empty.
// Returns an error if the date format is invalid.
func ParseTodoDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %s (expected YYYY-MM-DD)", dateStr)
	}
	return t, nil
}

// FormatTodoDate formats a time.Time into YYYY-MM-DD format.
// Returns empty string if the time is zero.
func FormatTodoDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}
