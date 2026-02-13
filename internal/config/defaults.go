package config

import (
	"os"
	"path/filepath"
)

// Defaults returns the default configuration.
func Defaults() *Config {
	homeDir, _ := os.UserHomeDir()

	return &Config{
		TodoFilePath:     filepath.Join(homeDir, ".local", "share", "todotxt-tui", "todo.txt"),
		DoneFilePath:     "", // Optional, empty by default
		Theme:            "default",
		ShowCompleted:    false,
		ArchiveCompleted: false,
		ConfirmDelete:    true,
		AutoSave:         true,
		FileWatchEnabled: true,
	}
}
