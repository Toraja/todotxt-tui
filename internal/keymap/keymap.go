package keymap

// Keymap defines the interface for managing keybindings.
type Keymap interface {
	// GetBinding returns the action for a key in a given mode.
	GetBinding(mode Mode, key string) Action

	// SetBinding sets a key binding for a mode and action.
	SetBinding(mode Mode, key string, action Action)

	// GetAvailableActions returns all available actions for a mode.
	GetAvailableActions(mode Mode) []Action

	// GetKeysForAction returns all keys bound to an action in a mode.
	GetKeysForAction(mode Mode, action Action) []string

	// ClearBinding removes a key binding.
	ClearBinding(mode Mode, key string)

	// ResetToDefaults resets all bindings to default values.
	ResetToDefaults()
}

// KeyBinding represents a single key binding.
type KeyBinding struct {
	Key         string // Key string (e.g., "j", "ctrl+d", "enter")
	Action      Action // Action to trigger
	Description string // Human-readable description
}

// NewKeymap creates a new Keymap with default bindings.
func NewKeymap() Keymap {
	km := &keymapImpl{
		bindings: make(map[Mode]map[string]Action),
	}
	km.ResetToDefaults()
	return km
}

// keymapImpl is the concrete implementation of the Keymap interface.
type keymapImpl struct {
	bindings map[Mode]map[string]Action // mode → key → action
}

// GetBinding returns the action for a key in a given mode.
func (km *keymapImpl) GetBinding(mode Mode, key string) Action {
	if modeBindings, ok := km.bindings[mode]; ok {
		if action, ok := modeBindings[key]; ok {
			return action
		}
	}
	return ActionNone
}

// SetBinding sets a key binding for a mode and action.
func (km *keymapImpl) SetBinding(mode Mode, key string, action Action) {
	if km.bindings[mode] == nil {
		km.bindings[mode] = make(map[string]Action)
	}
	km.bindings[mode][key] = action
}

// GetAvailableActions returns all available actions for a mode.
func (km *keymapImpl) GetAvailableActions(mode Mode) []Action {
	if modeBindings, ok := km.bindings[mode]; ok {
		actions := make([]Action, 0, len(modeBindings))
		seen := make(map[Action]bool)
		for _, action := range modeBindings {
			if !seen[action] {
				actions = append(actions, action)
				seen[action] = true
			}
		}
		return actions
	}
	return []Action{}
}

// GetKeysForAction returns all keys bound to an action in a mode.
func (km *keymapImpl) GetKeysForAction(mode Mode, action Action) []string {
	keys := make([]string, 0)
	if modeBindings, ok := km.bindings[mode]; ok {
		for key, act := range modeBindings {
			if act == action {
				keys = append(keys, key)
			}
		}
	}
	return keys
}

// ClearBinding removes a key binding.
func (km *keymapImpl) ClearBinding(mode Mode, key string) {
	if modeBindings, ok := km.bindings[mode]; ok {
		delete(modeBindings, key)
	}
}

// ResetToDefaults resets all bindings to default values.
func (km *keymapImpl) ResetToDefaults() {
	km.bindings = make(map[Mode]map[string]Action)
	loadDefaultBindings(km)
}
