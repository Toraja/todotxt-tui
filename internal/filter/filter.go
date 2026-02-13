package filter

import "github.com/Toraja/todotxt-tui/internal/parser"

// Filter defines the interface for filtering tasks.
type Filter interface {
	// Apply applies the filter to a list of tasks and returns matching tasks.
	Apply(tasks []*parser.Task, criteria []*FilterCriteria, logic FilterLogic) []*parser.Task

	// BuildIndex creates an index from the given tasks for fast lookups.
	BuildIndex(tasks []*parser.Task)

	// Search performs case-insensitive text search in task descriptions.
	Search(tasks []*parser.Task, query string) []*parser.Task

	// FilterByPriority returns tasks with the given priority.
	FilterByPriority(tasks []*parser.Task, priority string) []*parser.Task

	// FilterByContext returns tasks with the given context.
	FilterByContext(tasks []*parser.Task, context string) []*parser.Task

	// FilterByProject returns tasks with the given project.
	FilterByProject(tasks []*parser.Task, project string) []*parser.Task

	// FilterByCompletion returns tasks based on completion status.
	FilterByCompletion(tasks []*parser.Task, completed bool) []*parser.Task

	// GetIndex returns the current index.
	GetIndex() *Index
}

// NewFilter creates a new Filter implementation.
func NewFilter() Filter {
	return &filterImpl{
		index: NewIndex(),
	}
}

// filterImpl is the concrete implementation of the Filter interface.
type filterImpl struct {
	index *Index
}

// Apply applies multiple filter criteria to tasks.
func (f *filterImpl) Apply(tasks []*parser.Task, criteria []*FilterCriteria, logic FilterLogic) []*parser.Task {
	if len(criteria) == 0 {
		return tasks
	}

	// Filter only enabled criteria
	enabledCriteria := make([]*FilterCriteria, 0, len(criteria))
	for _, c := range criteria {
		if c.Enabled {
			enabledCriteria = append(enabledCriteria, c)
		}
	}

	if len(enabledCriteria) == 0 {
		return tasks
	}

	// Apply filters based on logic
	result := make([]*parser.Task, 0)

	if logic == FilterAnd {
		// AND logic: task must match ALL criteria
		for _, task := range tasks {
			matchesAll := true
			for _, c := range enabledCriteria {
				if !c.Matches(task) {
					matchesAll = false
					break
				}
			}
			if matchesAll {
				result = append(result, task)
			}
		}
	} else {
		// OR logic: task must match ANY criterion
		for _, task := range tasks {
			matchesAny := false
			for _, c := range enabledCriteria {
				if c.Matches(task) {
					matchesAny = true
					break
				}
			}
			if matchesAny {
				result = append(result, task)
			}
		}
	}

	return result
}

// BuildIndex builds the index from tasks.
func (f *filterImpl) BuildIndex(tasks []*parser.Task) {
	f.index.Build(tasks)
}

// Search performs case-insensitive text search.
func (f *filterImpl) Search(tasks []*parser.Task, query string) []*parser.Task {
	if query == "" {
		return tasks
	}

	result := make([]*parser.Task, 0)
	for _, task := range tasks {
		if containsIgnoreCase(task.Description, query) {
			result = append(result, task)
		}
	}
	return result
}

// FilterByPriority filters tasks by priority.
func (f *filterImpl) FilterByPriority(tasks []*parser.Task, priority string) []*parser.Task {
	result := make([]*parser.Task, 0)
	for _, task := range tasks {
		if task.Priority == priority {
			result = append(result, task)
		}
	}
	return result
}

// FilterByContext filters tasks by context.
func (f *filterImpl) FilterByContext(tasks []*parser.Task, context string) []*parser.Task {
	result := make([]*parser.Task, 0)
	for _, task := range tasks {
		if task.HasContext(context) {
			result = append(result, task)
		}
	}
	return result
}

// FilterByProject filters tasks by project.
func (f *filterImpl) FilterByProject(tasks []*parser.Task, project string) []*parser.Task {
	result := make([]*parser.Task, 0)
	for _, task := range tasks {
		if task.HasProject(project) {
			result = append(result, task)
		}
	}
	return result
}

// FilterByCompletion filters tasks by completion status.
func (f *filterImpl) FilterByCompletion(tasks []*parser.Task, completed bool) []*parser.Task {
	result := make([]*parser.Task, 0)
	for _, task := range tasks {
		if task.IsComplete() == completed {
			result = append(result, task)
		}
	}
	return result
}

// GetIndex returns the current index.
func (f *filterImpl) GetIndex() *Index {
	return f.index
}
