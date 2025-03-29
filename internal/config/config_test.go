package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".markin.yaml")

	// Create a test configuration file
	content := `project_dir: "~/Documents/notes"
daily_note_path: "daily"
daily_note_name: "test.md"
section: "## 💭 ✍️ ✨ Notes"
position: "after-heading"
create_section_if_missing: true
`
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	// Test loading the configuration
	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify the configuration values
	expected := &Config{
		ProjectDir:             "~/Documents/notes",
		DailyNotePath:          "daily",
		DailyNoteName:          "test.md",
		Section:                "## 💭 ✍️ ✨ Notes",
		Position:               "after-heading",
		CreateSectionIfMissing: true,
	}

	if cfg.ProjectDir != expected.ProjectDir {
		t.Errorf("Expected ProjectDir %s, got %s", expected.ProjectDir, cfg.ProjectDir)
	}
	if cfg.DailyNotePath != expected.DailyNotePath {
		t.Errorf("Expected DailyNotePath %s, got %s", expected.DailyNotePath, cfg.DailyNotePath)
	}
	if cfg.DailyNoteName != expected.DailyNoteName {
		t.Errorf("Expected DailyNoteName %s, got %s", expected.DailyNoteName, cfg.DailyNoteName)
	}
	if cfg.Section != expected.Section {
		t.Errorf("Expected Section %s, got %s", expected.Section, cfg.Section)
	}
	if cfg.Position != expected.Position {
		t.Errorf("Expected Position %s, got %s", expected.Position, cfg.Position)
	}
	if cfg.CreateSectionIfMissing != expected.CreateSectionIfMissing {
		t.Errorf("Expected CreateSectionIfMissing %v, got %v", expected.CreateSectionIfMissing, cfg.CreateSectionIfMissing)
	}
}

func TestLoadConfigMissingFile(t *testing.T) {
	// Test loading a non-existent configuration file
	_, err := LoadConfig("nonexistent.yaml")
	if err == nil {
		t.Error("Expected error when loading non-existent config file")
	}
}

func TestGenerateSampleConfig(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".markin.yaml")

	// Generate a sample configuration
	if err := GenerateSampleConfig(configPath); err != nil {
		t.Fatalf("Failed to generate sample config: %v", err)
	}

	// Verify the configuration file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("Sample configuration file was not created")
	}

	// Read and verify the content
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read sample config: %v", err)
	}

	// Verify the content contains expected fields
	expectedFields := []string{
		"project_dir:",
		"daily_note_path:",
		"daily_note_name:",
		"section:",
		"position:",
		"create_section_if_missing:",
	}

	for _, field := range expectedFields {
		if !strings.Contains(string(content), field) {
			t.Errorf("Sample config missing expected field: %s", field)
		}
	}
}

func TestGenerateSampleConfigExistingFile(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".markin.yaml")

	// Create an existing configuration file
	content := `project_dir: "~/Documents/notes"
daily_note_path: "daily"
daily_note_name: "test.md"
section: "## 💭 ✍️ ✨ Notes"
position: "after-heading"
create_section_if_missing: true
`
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	// Try to generate a sample configuration
	err := GenerateSampleConfig(configPath)
	if err == nil {
		t.Error("Expected error when generating sample config with existing file")
	}

	// Verify the content was not changed
	updatedContent, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read updated config: %v", err)
	}

	if string(updatedContent) != content {
		t.Error("Sample config overwrote existing configuration")
	}
}
