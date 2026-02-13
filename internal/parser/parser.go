package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Parser defines the interface for parsing todo.txt files.
type Parser interface {
	// ParseLine parses a single line of text into a Task.
	// lineNumber is 1-indexed for user-friendly error reporting.
	ParseLine(line string, lineNumber int) (*Task, error)

	// ParseFile parses an entire file's contents into a slice of Tasks.
	ParseFile(reader io.Reader) ([]*Task, error)

	// Serialize converts a Task back to todo.txt format string.
	Serialize(task *Task) string

	// Validate checks if a Task conforms to todo.txt format rules.
	Validate(task *Task) error
}

// NewParser creates a new Parser implementation.
func NewParser() Parser {
	return &parserImpl{
		priorityRe: regexp.MustCompile(`^\(([A-Z])\)\s+`),
		dateRe:     regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})\s+`),
		contextRe:  regexp.MustCompile(`@\S+`),
		projectRe:  regexp.MustCompile(`\+\S+`),
		metadataRe: regexp.MustCompile(`(\w+):(\S+)`),
	}
}

// parserImpl is the concrete implementation of the Parser interface.
type parserImpl struct {
	priorityRe *regexp.Regexp
	dateRe     *regexp.Regexp
	contextRe  *regexp.Regexp
	projectRe  *regexp.Regexp
	metadataRe *regexp.Regexp
}

// ParseLine parses a single line of todo.txt format into a Task.
func (p *parserImpl) ParseLine(line string, lineNumber int) (*Task, error) {
	if line == "" {
		return nil, fmt.Errorf("empty line")
	}

	task := &Task{
		RawLine:    line,
		LineNumber: lineNumber,
		Contexts:   make(map[string]struct{}),
		Projects:   make(map[string]struct{}),
		Metadata:   make(map[string]string),
	}

	remaining := strings.TrimSpace(line)

	// Check for completed marker
	if strings.HasPrefix(remaining, "x ") {
		task.Completed = true
		remaining = strings.TrimSpace(remaining[2:])

		// Next might be completion date
		if matches := p.dateRe.FindStringSubmatch(remaining); len(matches) > 0 {
			if date, err := ParseTodoDate(matches[1]); err == nil {
				task.CompletionDate = date
				remaining = strings.TrimSpace(remaining[len(matches[0]):])
			}
		}
	}

	// Check for priority (only if not completed)
	if !task.Completed {
		if matches := p.priorityRe.FindStringSubmatch(remaining); len(matches) > 0 {
			task.Priority = matches[1]
			remaining = strings.TrimSpace(remaining[len(matches[0]):])
		}
	}

	// Check for creation date
	if matches := p.dateRe.FindStringSubmatch(remaining); len(matches) > 0 {
		if date, err := ParseTodoDate(matches[1]); err == nil {
			task.CreationDate = date
			remaining = strings.TrimSpace(remaining[len(matches[0]):])
		}
	}

	// Remaining is the description
	task.Description = remaining

	// Extract contexts
	contexts := p.contextRe.FindAllString(remaining, -1)
	for _, c := range contexts {
		task.Contexts[c] = struct{}{}
	}

	// Extract projects
	projects := p.projectRe.FindAllString(remaining, -1)
	for _, proj := range projects {
		task.Projects[proj] = struct{}{}
	}

	// Extract metadata key:value pairs
	metadataMatches := p.metadataRe.FindAllStringSubmatch(remaining, -1)
	for _, match := range metadataMatches {
		if len(match) >= 3 {
			task.Metadata[match[1]] = match[2]
		}
	}

	return task, nil
}

// ParseFile parses an entire file into a slice of Tasks.
func (p *parserImpl) ParseFile(reader io.Reader) ([]*Task, error) {
	var tasks []*Task
	scanner := bufio.NewScanner(reader)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			lineNumber++
			continue
		}

		task, err := p.ParseLine(line, lineNumber)
		if err != nil {
			// Log warning but continue parsing
			// This allows graceful degradation for malformed lines
			lineNumber++
			continue
		}

		tasks = append(tasks, task)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return tasks, nil
}

// Serialize converts a Task back to todo.txt format string.
func (p *parserImpl) Serialize(task *Task) string {
	return task.String()
}

// Validate checks if a Task conforms to todo.txt format rules.
func (p *parserImpl) Validate(task *Task) error {
	// Check priority
	if task.Priority != "" {
		if len(task.Priority) != 1 {
			return fmt.Errorf("priority must be single letter A-Z")
		}
		if task.Priority[0] < 'A' || task.Priority[0] > 'Z' {
			return fmt.Errorf("priority must be letter A-Z")
		}
	}

	// Check dates
	if !task.CreationDate.IsZero() {
		// Creation date is valid if it parses (already validated)
	}

	if task.Completed && task.CompletionDate.IsZero() {
		return fmt.Errorf("completed task must have completion date")
	}

	// Check description
	if strings.TrimSpace(task.Description) == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	// Validate contexts
	for ctx := range task.Contexts {
		if err := ValidateContextTag(ctx); err != nil {
			return fmt.Errorf("invalid context %s: %w", ctx, err)
		}
	}

	// Validate projects
	for proj := range task.Projects {
		if err := ValidateProjectTag(proj); err != nil {
			return fmt.Errorf("invalid project %s: %w", proj, err)
		}
	}

	return nil
}
