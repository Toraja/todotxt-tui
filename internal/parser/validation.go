package parser

import (
	"fmt"
	"strings"
	"unicode"
)

// ValidateContextTag validates a context tag format.
// Context tags must start with @ and be followed by a valid word.
func ValidateContextTag(tag string) error {
	if !strings.HasPrefix(tag, "@") {
		return fmt.Errorf("context must start with @")
	}
	if len(tag) < 2 {
		return fmt.Errorf("context tag too short")
	}
	rest := tag[1:]
	if !isValidTagWord(rest) {
		return fmt.Errorf("context must be single word (no spaces)")
	}
	return nil
}

// ValidateProjectTag validates a project tag format.
// Project tags must start with + and be followed by a valid word.
func ValidateProjectTag(tag string) error {
	if !strings.HasPrefix(tag, "+") {
		return fmt.Errorf("project must start with +")
	}
	if len(tag) < 2 {
		return fmt.Errorf("project tag too short")
	}
	rest := tag[1:]
	if !isValidTagWord(rest) {
		return fmt.Errorf("project must be single word (no spaces)")
	}
	return nil
}

// isValidTagWord checks if a string is a valid tag word.
// Valid tag words contain only letters, digits, underscores, and hyphens.
func isValidTagWord(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
			return false
		}
	}
	return true
}
