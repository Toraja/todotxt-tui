package filter

import "github.com/Toraja/todotxt-tui/internal/parser"

// SortCriteria defines the field to sort by.
type SortCriteria int

const (
	SortNone SortCriteria = iota
	SortPriority
	SortCreationDate
	SortCompletionDate
	SortDescription
)

// SortOrder defines ascending or descending sort.
type SortOrder int

const (
	SortAscending SortOrder = iota
	SortDescending
)

// TaskList represents the filtered and sorted view of tasks.
type TaskList struct {
	// Source data
	AllTasks []*parser.Task // All tasks from TodoFile (immutable reference)

	// Filter state
	ActiveFilters []*FilterCriteria // Currently active filters
	FilterLogic   FilterLogic       // How to combine filters (AND/OR)
	filter        Filter            // Filter implementation

	// Sort state
	SortBy    SortCriteria // Current sort field
	SortOrder SortOrder    // Ascending or descending

	// View state
	VisibleTasks  []*parser.Task // Filtered and sorted subset of tasks
	SelectedIndex int            // Index of currently selected task in VisibleTasks
	ScrollOffset  int            // Scroll offset for viewport
	ViewportSize  int            // Number of tasks visible in viewport

	// Computed
	TotalCount   int // Total tasks in AllTasks
	VisibleCount int // Tasks in VisibleTasks
}

// NewTaskList creates a new TaskList.
func NewTaskList(tasks []*parser.Task) *TaskList {
	tl := &TaskList{
		AllTasks:      tasks,
		ActiveFilters: []*FilterCriteria{},
		FilterLogic:   FilterAnd,
		filter:        NewFilter(),
		SortBy:        SortNone,
		SortOrder:     SortAscending,
		VisibleTasks:  tasks,
		SelectedIndex: 0,
		ScrollOffset:  0,
		ViewportSize:  10, // Default viewport size
		TotalCount:    len(tasks),
		VisibleCount:  len(tasks),
	}

	// Build index
	tl.filter.BuildIndex(tasks)

	return tl
}

// ApplyFilters recomputes VisibleTasks based on ActiveFilters.
func (tl *TaskList) ApplyFilters() {
	tl.VisibleTasks = tl.filter.Apply(tl.AllTasks, tl.ActiveFilters, tl.FilterLogic)
	tl.VisibleCount = len(tl.VisibleTasks)

	// Apply sorting
	tl.sort()

	// Adjust selected index if out of range
	if tl.SelectedIndex >= tl.VisibleCount {
		if tl.VisibleCount > 0 {
			tl.SelectedIndex = tl.VisibleCount - 1
		} else {
			tl.SelectedIndex = -1
		}
	}

	// Adjust scroll offset
	tl.adjustScrollOffset()
}

// AddFilter adds a filter to ActiveFilters and recomputes.
func (tl *TaskList) AddFilter(filter *FilterCriteria) {
	tl.ActiveFilters = append(tl.ActiveFilters, filter)
	tl.ApplyFilters()
}

// RemoveFilter removes a filter by type and recomputes.
func (tl *TaskList) RemoveFilter(filterType FilterType) {
	newFilters := make([]*FilterCriteria, 0)
	for _, f := range tl.ActiveFilters {
		if f.Type != filterType {
			newFilters = append(newFilters, f)
		}
	}
	tl.ActiveFilters = newFilters
	tl.ApplyFilters()
}

// ClearFilters clears all ActiveFilters and shows all tasks.
func (tl *TaskList) ClearFilters() {
	tl.ActiveFilters = []*FilterCriteria{}
	tl.ApplyFilters()
}

// SetSort sorts VisibleTasks by the given criteria.
func (tl *TaskList) SetSort(by SortCriteria, order SortOrder) {
	tl.SortBy = by
	tl.SortOrder = order
	tl.sort()
}

// sort sorts VisibleTasks based on SortBy and SortOrder.
func (tl *TaskList) sort() {
	if tl.SortBy == SortNone {
		return
	}

	// Simple bubble sort (fine for typical todo.txt sizes)
	// TODO: Use more efficient sorting for very large files
	for i := 0; i < len(tl.VisibleTasks)-1; i++ {
		for j := i + 1; j < len(tl.VisibleTasks); j++ {
			if tl.shouldSwap(tl.VisibleTasks[i], tl.VisibleTasks[j]) {
				tl.VisibleTasks[i], tl.VisibleTasks[j] = tl.VisibleTasks[j], tl.VisibleTasks[i]
			}
		}
	}
}

// shouldSwap returns true if task i should come after task j.
func (tl *TaskList) shouldSwap(i, j *parser.Task) bool {
	var cmp int

	switch tl.SortBy {
	case SortPriority:
		// Empty priority sorts last
		if i.Priority == "" && j.Priority != "" {
			cmp = 1
		} else if i.Priority != "" && j.Priority == "" {
			cmp = -1
		} else {
			// A < B < C ... < Z
			if i.Priority < j.Priority {
				cmp = -1
			} else if i.Priority > j.Priority {
				cmp = 1
			}
		}

	case SortCreationDate:
		if i.CreationDate.Before(j.CreationDate) {
			cmp = -1
		} else if i.CreationDate.After(j.CreationDate) {
			cmp = 1
		}

	case SortCompletionDate:
		if i.CompletionDate.Before(j.CompletionDate) {
			cmp = -1
		} else if i.CompletionDate.After(j.CompletionDate) {
			cmp = 1
		}

	case SortDescription:
		if i.Description < j.Description {
			cmp = -1
		} else if i.Description > j.Description {
			cmp = 1
		}
	}

	// Apply sort order
	if tl.SortOrder == SortDescending {
		cmp = -cmp
	}

	return cmp > 0
}

// SelectNext moves selection down (cyclic).
func (tl *TaskList) SelectNext() {
	if tl.VisibleCount == 0 {
		tl.SelectedIndex = -1
		return
	}

	tl.SelectedIndex = (tl.SelectedIndex + 1) % tl.VisibleCount
	tl.adjustScrollOffset()
}

// SelectPrev moves selection up (cyclic).
func (tl *TaskList) SelectPrev() {
	if tl.VisibleCount == 0 {
		tl.SelectedIndex = -1
		return
	}

	tl.SelectedIndex--
	if tl.SelectedIndex < 0 {
		tl.SelectedIndex = tl.VisibleCount - 1
	}
	tl.adjustScrollOffset()
}

// SelectFirst moves selection to first task.
func (tl *TaskList) SelectFirst() {
	if tl.VisibleCount > 0 {
		tl.SelectedIndex = 0
		tl.adjustScrollOffset()
	}
}

// SelectLast moves selection to last task.
func (tl *TaskList) SelectLast() {
	if tl.VisibleCount > 0 {
		tl.SelectedIndex = tl.VisibleCount - 1
		tl.adjustScrollOffset()
	}
}

// GetSelectedTask returns the currently selected task or nil.
func (tl *TaskList) GetSelectedTask() *parser.Task {
	if tl.SelectedIndex >= 0 && tl.SelectedIndex < tl.VisibleCount {
		return tl.VisibleTasks[tl.SelectedIndex]
	}
	return nil
}

// ScrollDown moves scroll offset down.
func (tl *TaskList) ScrollDown() {
	maxOffset := tl.VisibleCount - tl.ViewportSize
	if maxOffset < 0 {
		maxOffset = 0
	}

	if tl.ScrollOffset < maxOffset {
		tl.ScrollOffset++
	}
}

// ScrollUp moves scroll offset up.
func (tl *TaskList) ScrollUp() {
	if tl.ScrollOffset > 0 {
		tl.ScrollOffset--
	}
}

// GetVisibleTasks returns tasks in current viewport.
func (tl *TaskList) GetVisibleTasks() []*parser.Task {
	start := tl.ScrollOffset
	end := start + tl.ViewportSize

	if start >= tl.VisibleCount {
		return []*parser.Task{}
	}

	if end > tl.VisibleCount {
		end = tl.VisibleCount
	}

	return tl.VisibleTasks[start:end]
}

// GetTaskAtViewportPosition returns the task at the given viewport position.
func (tl *TaskList) GetTaskAtViewportPosition(pos int) *parser.Task {
	index := tl.ScrollOffset + pos
	if index >= 0 && index < tl.VisibleCount {
		return tl.VisibleTasks[index]
	}
	return nil
}

// adjustScrollOffset ensures selected task is visible in viewport.
func (tl *TaskList) adjustScrollOffset() {
	if tl.SelectedIndex < tl.ScrollOffset {
		// Selected task is above viewport, scroll up
		tl.ScrollOffset = tl.SelectedIndex
	} else if tl.SelectedIndex >= tl.ScrollOffset+tl.ViewportSize {
		// Selected task is below viewport, scroll down
		tl.ScrollOffset = tl.SelectedIndex - tl.ViewportSize + 1
		if tl.ScrollOffset < 0 {
			tl.ScrollOffset = 0
		}
	}
}

// SetViewportSize updates the viewport size.
func (tl *TaskList) SetViewportSize(size int) {
	if size > 0 {
		tl.ViewportSize = size
		tl.adjustScrollOffset()
	}
}

// Refresh rebuilds the index and reapplies filters.
func (tl *TaskList) Refresh() {
	tl.TotalCount = len(tl.AllTasks)
	tl.filter.BuildIndex(tl.AllTasks)
	tl.ApplyFilters()
}
