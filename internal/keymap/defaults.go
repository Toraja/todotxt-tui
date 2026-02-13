package keymap

// loadDefaultBindings loads the default vim-style keybindings.
func loadDefaultBindings(km *keymapImpl) {
	// Normal mode bindings
	normalMode := make(map[string]Action)

	// Navigation (vim-style)
	normalMode["j"] = ActionMoveDown
	normalMode["k"] = ActionMoveUp
	normalMode["g"] = ActionMoveTop
	normalMode["G"] = ActionMoveBottom
	normalMode["h"] = ActionMoveLeft
	normalMode["l"] = ActionMoveRight
	normalMode["ctrl+d"] = ActionPageDown
	normalMode["ctrl+u"] = ActionPageUp
	normalMode["ctrl+e"] = ActionScrollDown
	normalMode["ctrl+y"] = ActionScrollUp

	// Task selection
	normalMode["enter"] = ActionDetailView

	// Task operations
	normalMode["a"] = ActionAddTask
	normalMode["e"] = ActionEditTask
	normalMode[" "] = ActionToggleTask // Space
	normalMode["d"] = ActionDeleteTask
	normalMode["+"] = ActionIncreasePrio
	normalMode["-"] = ActionDecreasePrio
	normalMode["0"] = ActionClearPrio

	// Filter operations
	normalMode["f"] = ActionFilterToggle
	normalMode["c"] = ActionFilterContext
	normalMode["p"] = ActionFilterProject
	normalMode["P"] = ActionFilterPriority
	normalMode["/"] = ActionSearch
	normalMode["F"] = ActionClearFilters

	// File operations
	normalMode["s"] = ActionSave
	normalMode["r"] = ActionReload

	// Application
	normalMode["q"] = ActionQuit
	normalMode["?"] = ActionHelp
	normalMode["esc"] = ActionCancel

	km.bindings[ModeNormal] = normalMode

	// Insert mode bindings
	insertMode := make(map[string]Action)
	insertMode["esc"] = ActionCancel
	insertMode["enter"] = ActionConfirm
	km.bindings[ModeInsert] = insertMode

	// Dialog mode bindings
	dialogMode := make(map[string]Action)
	dialogMode["y"] = ActionConfirm
	dialogMode["n"] = ActionCancel
	dialogMode["enter"] = ActionConfirm
	dialogMode["esc"] = ActionCancel
	km.bindings[ModeDialog] = dialogMode

	// Search mode bindings
	searchMode := make(map[string]Action)
	searchMode["esc"] = ActionCancel
	searchMode["enter"] = ActionConfirm
	km.bindings[ModeSearch] = searchMode
}

// GetDefaultBindings returns the default keybindings for documentation.
func GetDefaultBindings() map[Mode][]KeyBinding {
	bindings := make(map[Mode][]KeyBinding)

	// Normal mode
	bindings[ModeNormal] = []KeyBinding{
		{"j", ActionMoveDown, "Move selection down"},
		{"k", ActionMoveUp, "Move selection up"},
		{"g", ActionMoveTop, "Move to top"},
		{"G", ActionMoveBottom, "Move to bottom"},
		{"h", ActionMoveLeft, "Move left"},
		{"l", ActionMoveRight, "Move right"},
		{"ctrl+d", ActionPageDown, "Page down"},
		{"ctrl+u", ActionPageUp, "Page up"},
		{"ctrl+e", ActionScrollDown, "Scroll down"},
		{"ctrl+y", ActionScrollUp, "Scroll up"},
		{"enter", ActionDetailView, "Show detail view"},
		{"a", ActionAddTask, "Add new task"},
		{"e", ActionEditTask, "Edit task"},
		{"space", ActionToggleTask, "Toggle task completion"},
		{"d", ActionDeleteTask, "Delete task"},
		{"+", ActionIncreasePrio, "Increase priority"},
		{"-", ActionDecreasePrio, "Decrease priority"},
		{"0", ActionClearPrio, "Clear priority"},
		{"f", ActionFilterToggle, "Toggle filter"},
		{"c", ActionFilterContext, "Filter by context"},
		{"p", ActionFilterProject, "Filter by project"},
		{"P", ActionFilterPriority, "Filter by priority"},
		{"/", ActionSearch, "Search tasks"},
		{"F", ActionClearFilters, "Clear all filters"},
		{"s", ActionSave, "Save file"},
		{"r", ActionReload, "Reload file"},
		{"q", ActionQuit, "Quit application"},
		{"?", ActionHelp, "Show help"},
		{"esc", ActionCancel, "Cancel/escape"},
	}

	// Insert mode
	bindings[ModeInsert] = []KeyBinding{
		{"esc", ActionCancel, "Cancel editing"},
		{"enter", ActionConfirm, "Confirm edit"},
	}

	// Dialog mode
	bindings[ModeDialog] = []KeyBinding{
		{"y", ActionConfirm, "Confirm (yes)"},
		{"n", ActionCancel, "Cancel (no)"},
		{"enter", ActionConfirm, "Confirm"},
		{"esc", ActionCancel, "Cancel"},
	}

	// Search mode
	bindings[ModeSearch] = []KeyBinding{
		{"esc", ActionCancel, "Cancel search"},
		{"enter", ActionConfirm, "Apply search"},
	}

	return bindings
}
