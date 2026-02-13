package config

// Config holds all application configuration settings.
type Config struct {
	// File paths
	TodoFilePath string `yaml:"todo_file_path"`
	DoneFilePath string `yaml:"done_file_path"`
	ConfigPath   string `yaml:"-"` // Not loaded from file

	// Display settings
	Theme            string `yaml:"theme"`
	ShowCompleted    bool   `yaml:"show_completed"`
	ArchiveCompleted bool   `yaml:"archive_completed"`

	// UI settings
	ConfirmDelete    bool `yaml:"confirm_delete"`
	AutoSave         bool `yaml:"auto_save"`
	FileWatchEnabled bool `yaml:"file_watch_enabled"`
}
