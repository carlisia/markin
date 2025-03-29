package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	ProjectDir             string `yaml:"project_dir"`
	DailyNotePath          string `yaml:"daily_note_path"`
	DailyNoteName          string `yaml:"daily_note_name"`
	Section                string `yaml:"section"`
	Position               string `yaml:"position"`
	CreateSectionIfMissing bool   `yaml:"create_section_if_missing"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(configPath string) (*Config, error) {
	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to find home directory: %w", err)
	}

	// If no config path provided, use the default location
	if configPath == "" {
		configPath = filepath.Join(homeDir, ".config", "markin", ".markin.yaml")
	}

	// Read the file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Parse the YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse configuration file at %s: %w", configPath, err)
	}

	return &config, nil
}

// GenerateSampleConfig generates a sample configuration file
func GenerateSampleConfig(configPath string) error {
	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to find home directory: %w", err)
	}

	// If no config path provided, use the default location
	if configPath == "" {
		configPath = filepath.Join(homeDir, ".config", "markin", ".markin.yaml")
	}

	// Check if file exists
	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("configuration file already exists at %s", configPath)
	}

	sample := `# Markin Configuration

# The root directory of your project
project_dir: "~/Documents/notes"

# The path to your daily notes relative to project_dir
daily_note_path: "daily"

# The name of your daily note file (can include date format)
daily_note_name: "{{.Date}}.md"

# The section to insert lines into
section: "## 💭 ✍️ ✨ Notes"

# Where to insert new lines in the section
# Options: "after-heading" or "before-end"
position: "after-heading"

# Whether to create the section if it doesn't exist
create_section_if_missing: true
`

	// Create the config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create configuration directory at %s: %w", configDir, err)
	}

	// Write the sample config
	if err := os.WriteFile(configPath, []byte(sample), 0644); err != nil {
		return fmt.Errorf("failed to write sample configuration to %s: %w", configPath, err)
	}

	return nil
}
