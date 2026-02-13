package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Validate checks if the configuration is valid.
func Validate(cfg *Config) error {
	// Validate TodoFilePath
	if cfg.TodoFilePath == "" {
		return fmt.Errorf("todo_file_path cannot be empty")
	}

	// Expand ~ to home directory
	if strings.HasPrefix(cfg.TodoFilePath, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		cfg.TodoFilePath = filepath.Join(homeDir, cfg.TodoFilePath[2:])
	}

	// Make path absolute
	if !filepath.IsAbs(cfg.TodoFilePath) {
		absPath, err := filepath.Abs(cfg.TodoFilePath)
		if err != nil {
			return fmt.Errorf("failed to get absolute path for todo_file_path: %w", err)
		}
		cfg.TodoFilePath = absPath
	}

	// Validate DoneFilePath if specified
	if cfg.DoneFilePath != "" {
		if strings.HasPrefix(cfg.DoneFilePath, "~/") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err)
			}
			cfg.DoneFilePath = filepath.Join(homeDir, cfg.DoneFilePath[2:])
		}

		if !filepath.IsAbs(cfg.DoneFilePath) {
			absPath, err := filepath.Abs(cfg.DoneFilePath)
			if err != nil {
				return fmt.Errorf("failed to get absolute path for done_file_path: %w", err)
			}
			cfg.DoneFilePath = absPath
		}
	}

	// Validate theme
	validThemes := map[string]bool{
		"light":   true,
		"dark":    true,
		"default": true,
	}
	if !validThemes[cfg.Theme] {
		return fmt.Errorf("invalid theme: %s (must be light, dark, or default)", cfg.Theme)
	}

	// AutoSave must be true (non-negotiable for data safety)
	if !cfg.AutoSave {
		cfg.AutoSave = true // Force enable
	}

	return nil
}
