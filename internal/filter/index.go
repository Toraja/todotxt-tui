package filter

import "github.com/Toraja/todotxt-tui/internal/parser"

// Index provides fast lookups for filtering tasks.
type Index struct {
	// Maps from priority/context/project to task indices
	PriorityIndex map[string][]int // priority letter → task indices
	ContextIndex  map[string][]int // @context → task indices
	ProjectIndex  map[string][]int // +project → task indices
	CompletedIdx  []int            // Indices of completed tasks
	ActiveIdx     []int            // Indices of active (non-completed) tasks
}

// NewIndex creates a new empty index.
func NewIndex() *Index {
	return &Index{
		PriorityIndex: make(map[string][]int),
		ContextIndex:  make(map[string][]int),
		ProjectIndex:  make(map[string][]int),
		CompletedIdx:  []int{},
		ActiveIdx:     []int{},
	}
}

// Build constructs the index from a list of tasks.
// This is an O(n) operation where n is the number of tasks.
func (idx *Index) Build(tasks []*parser.Task) {
	// Clear existing index
	idx.PriorityIndex = make(map[string][]int)
	idx.ContextIndex = make(map[string][]int)
	idx.ProjectIndex = make(map[string][]int)
	idx.CompletedIdx = []int{}
	idx.ActiveIdx = []int{}

	// Build indexes
	for i, task := range tasks {
		// Priority index
		if task.Priority != "" {
			idx.PriorityIndex[task.Priority] = append(idx.PriorityIndex[task.Priority], i)
		}

		// Context index
		for ctx := range task.Contexts {
			idx.ContextIndex[ctx] = append(idx.ContextIndex[ctx], i)
		}

		// Project index
		for proj := range task.Projects {
			idx.ProjectIndex[proj] = append(idx.ProjectIndex[proj], i)
		}

		// Completion index
		if task.IsComplete() {
			idx.CompletedIdx = append(idx.CompletedIdx, i)
		} else {
			idx.ActiveIdx = append(idx.ActiveIdx, i)
		}
	}
}

// GetByPriority returns indices of tasks with the given priority.
func (idx *Index) GetByPriority(priority string) []int {
	if indices, ok := idx.PriorityIndex[priority]; ok {
		return indices
	}
	return []int{}
}

// GetByContext returns indices of tasks with the given context.
func (idx *Index) GetByContext(context string) []int {
	if indices, ok := idx.ContextIndex[context]; ok {
		return indices
	}
	return []int{}
}

// GetByProject returns indices of tasks with the given project.
func (idx *Index) GetByProject(project string) []int {
	if indices, ok := idx.ProjectIndex[project]; ok {
		return indices
	}
	return []int{}
}

// GetCompleted returns indices of completed tasks.
func (idx *Index) GetCompleted() []int {
	return idx.CompletedIdx
}

// GetActive returns indices of active (non-completed) tasks.
func (idx *Index) GetActive() []int {
	return idx.ActiveIdx
}

// GetAllPriorities returns all unique priorities in the index.
func (idx *Index) GetAllPriorities() []string {
	priorities := make([]string, 0, len(idx.PriorityIndex))
	for p := range idx.PriorityIndex {
		priorities = append(priorities, p)
	}
	return priorities
}

// GetAllContexts returns all unique contexts in the index.
func (idx *Index) GetAllContexts() []string {
	contexts := make([]string, 0, len(idx.ContextIndex))
	for c := range idx.ContextIndex {
		contexts = append(contexts, c)
	}
	return contexts
}

// GetAllProjects returns all unique projects in the index.
func (idx *Index) GetAllProjects() []string {
	projects := make([]string, 0, len(idx.ProjectIndex))
	for p := range idx.ProjectIndex {
		projects = append(projects, p)
	}
	return projects
}
