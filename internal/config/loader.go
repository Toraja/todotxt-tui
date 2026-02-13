package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Load loads configuration from a YAML file.
// If the file doesn't exist, returns default configuration.
// Environment variables can override file settings (TODO: implement).
func Load(configPath string) (*Config, error) {
	// Start with defaults
	cfg := Defaults()
	cfg.ConfigPath = configPath

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// No config file, use defaults
		return cfg, nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// TODO: Apply environment variable overrides

	// Validate configuration
	if err := Validate(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}
