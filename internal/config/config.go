package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	MarkdownDir            string `mapstructure:"markdown_dir"`
	MarkdownFile           string `mapstructure:"markdown_file"`
	Section                string `mapstructure:"section"`
	Position               string `mapstructure:"position"`
	CreateSectionIfMissing bool   `mapstructure:"create_section_if_missing"`
}

// LoadConfig reads and parses the configuration file
func LoadConfig(cfgFile string) (*Config, error) {
	if cfgFile == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("error finding home directory: %w", err)
		}
		cfgFile = filepath.Join(homeDir, ".config", "markin", ".markin.yaml")
	}

	configData, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	expanded := os.ExpandEnv(string(configData))
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(expanded)); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &cfg, nil
}

// GenerateSampleConfig creates a sample configuration file
func GenerateSampleConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "markin")
	configPath := filepath.Join(configDir, ".markin.yaml")

	sample := `markdown_dir: $VAULT_MAIN
markdown_file: daily.md
section: "## 💭 ✍️ ✨ Notes"
position: after-heading
create_section_if_missing: true`

	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	if err := os.WriteFile(configPath, []byte(sample), 0644); err != nil {
		return fmt.Errorf("error writing sample config file: %w", err)
	}

	return nil
}
