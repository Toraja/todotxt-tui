package keymap

// Mode represents the current input mode of the application.
type Mode int

const (
	ModeNormal Mode = iota // Normal mode (navigation, commands)
	ModeInsert             // Insert mode (editing task text)
	ModeDialog             // Dialog mode (confirmation prompts)
	ModeSearch             // Search mode (filtering tasks)
)

// String returns the string representation of the mode.
func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "NORMAL"
	case ModeInsert:
		return "INSERT"
	case ModeDialog:
		return "DIALOG"
	case ModeSearch:
		return "SEARCH"
	default:
		return "UNKNOWN"
	}
}

// Action represents a user action that can be triggered by a key.
type Action int

const (
	ActionNone Action = iota

	// Navigation
	ActionMoveDown   // Move selection down (j)
	ActionMoveUp     // Move selection up (k)
	ActionMoveTop    // Move to top (g)
	ActionMoveBottom // Move to bottom (G)
	ActionPageDown   // Page down (Ctrl+d)
	ActionPageUp     // Page up (Ctrl+u)
	ActionScrollDown // Scroll down (Ctrl+e)
	ActionScrollUp   // Scroll up (Ctrl+y)
	ActionMoveLeft   // Move left (h)
	ActionMoveRight  // Move right (l)
	ActionSelectTask // Select/activate task (Enter)
	ActionDetailView // Show detail view (Enter)

	// Task operations
	ActionAddTask      // Add new task (a)
	ActionEditTask     // Edit selected task (e)
	ActionToggleTask   // Toggle task completion (Space)
	ActionDeleteTask   // Delete task (d)
	ActionIncreasePrio // Increase priority (+)
	ActionDecreasePrio // Decrease priority (-)
	ActionClearPrio    // Clear priority (0)

	// Filter operations
	ActionFilterToggle   // Toggle filter (f)
	ActionFilterContext  // Filter by context (c)
	ActionFilterProject  // Filter by project (p)
	ActionFilterPriority // Filter by priority (P)
	ActionSearch         // Search tasks (/)
	ActionClearFilters   // Clear all filters (F)

	// File operations
	ActionSave   // Save file (s)
	ActionReload // Reload file (r)

	// Application
	ActionQuit    // Quit application (q)
	ActionHelp    // Show help (?)
	ActionCancel  // Cancel/escape (Esc)
	ActionConfirm // Confirm action (y/Enter)
)

// String returns the string representation of the action.
func (a Action) String() string {
	switch a {
	case ActionNone:
		return "None"
	case ActionMoveDown:
		return "MoveDown"
	case ActionMoveUp:
		return "MoveUp"
	case ActionMoveTop:
		return "MoveTop"
	case ActionMoveBottom:
		return "MoveBottom"
	case ActionPageDown:
		return "PageDown"
	case ActionPageUp:
		return "PageUp"
	case ActionScrollDown:
		return "ScrollDown"
	case ActionScrollUp:
		return "ScrollUp"
	case ActionMoveLeft:
		return "MoveLeft"
	case ActionMoveRight:
		return "MoveRight"
	case ActionSelectTask:
		return "SelectTask"
	case ActionDetailView:
		return "DetailView"
	case ActionAddTask:
		return "AddTask"
	case ActionEditTask:
		return "EditTask"
	case ActionToggleTask:
		return "ToggleTask"
	case ActionDeleteTask:
		return "DeleteTask"
	case ActionIncreasePrio:
		return "IncreasePrio"
	case ActionDecreasePrio:
		return "DecreasePrio"
	case ActionClearPrio:
		return "ClearPrio"
	case ActionFilterToggle:
		return "FilterToggle"
	case ActionFilterContext:
		return "FilterContext"
	case ActionFilterProject:
		return "FilterProject"
	case ActionFilterPriority:
		return "FilterPriority"
	case ActionSearch:
		return "Search"
	case ActionClearFilters:
		return "ClearFilters"
	case ActionSave:
		return "Save"
	case ActionReload:
		return "Reload"
	case ActionQuit:
		return "Quit"
	case ActionHelp:
		return "Help"
	case ActionCancel:
		return "Cancel"
	case ActionConfirm:
		return "Confirm"
	default:
		return "Unknown"
	}
}

// Description returns a human-readable description of the action.
func (a Action) Description() string {
	switch a {
	case ActionNone:
		return "No action"
	case ActionMoveDown:
		return "Move selection down"
	case ActionMoveUp:
		return "Move selection up"
	case ActionMoveTop:
		return "Move to top"
	case ActionMoveBottom:
		return "Move to bottom"
	case ActionPageDown:
		return "Page down"
	case ActionPageUp:
		return "Page up"
	case ActionScrollDown:
		return "Scroll down"
	case ActionScrollUp:
		return "Scroll up"
	case ActionMoveLeft:
		return "Move left"
	case ActionMoveRight:
		return "Move right"
	case ActionSelectTask:
		return "Select task"
	case ActionDetailView:
		return "Show detail view"
	case ActionAddTask:
		return "Add new task"
	case ActionEditTask:
		return "Edit task"
	case ActionToggleTask:
		return "Toggle task completion"
	case ActionDeleteTask:
		return "Delete task"
	case ActionIncreasePrio:
		return "Increase priority"
	case ActionDecreasePrio:
		return "Decrease priority"
	case ActionClearPrio:
		return "Clear priority"
	case ActionFilterToggle:
		return "Toggle filter"
	case ActionFilterContext:
		return "Filter by context"
	case ActionFilterProject:
		return "Filter by project"
	case ActionFilterPriority:
		return "Filter by priority"
	case ActionSearch:
		return "Search tasks"
	case ActionClearFilters:
		return "Clear all filters"
	case ActionSave:
		return "Save file"
	case ActionReload:
		return "Reload file"
	case ActionQuit:
		return "Quit application"
	case ActionHelp:
		return "Show help"
	case ActionCancel:
		return "Cancel/escape"
	case ActionConfirm:
		return "Confirm action"
	default:
		return "Unknown action"
	}
}
