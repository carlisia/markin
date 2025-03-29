package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	testConfigPath := filepath.Join(tmpDir, "test_config.yaml")

	// Test data
	testConfig := `markdown_dir: /test/dir
markdown_file: test.md
section: "## Test Section"
position: after-heading
create_section_if_missing: true`

	// Write test config file
	if err := os.WriteFile(testConfigPath, []byte(testConfig), 0644); err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}

	// Test loading config
	cfg, err := LoadConfig(testConfigPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify config values
	expected := Config{
		MarkdownDir:            "/test/dir",
		MarkdownFile:           "test.md",
		Section:                "## Test Section",
		Position:               "after-heading",
		CreateSectionIfMissing: true,
	}

	if *cfg != expected {
		t.Errorf("Config mismatch:\nGot:      %+v\nExpected: %+v", *cfg, expected)
	}
}

func TestLoadConfigWithEnvVars(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	testConfigPath := filepath.Join(tmpDir, "test_config.yaml")

	// Set environment variable
	os.Setenv("TEST_DIR", "/test/dir")
	defer os.Unsetenv("TEST_DIR")

	// Test data with environment variable
	testConfig := `markdown_dir: $TEST_DIR
markdown_file: test.md
section: "## Test Section"
position: after-heading
create_section_if_missing: true`

	// Write test config file
	if err := os.WriteFile(testConfigPath, []byte(testConfig), 0644); err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}

	// Test loading config
	cfg, err := LoadConfig(testConfigPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify config values
	expected := Config{
		MarkdownDir:            "/test/dir",
		MarkdownFile:           "test.md",
		Section:                "## Test Section",
		Position:               "after-heading",
		CreateSectionIfMissing: true,
	}

	if *cfg != expected {
		t.Errorf("Config mismatch:\nGot:      %+v\nExpected: %+v", *cfg, expected)
	}
}

func TestLoadConfigMissingFile(t *testing.T) {
	// Test loading non-existent config file
	_, err := LoadConfig("/non/existent/file.yaml")
	if err == nil {
		t.Error("LoadConfig should fail for non-existent file")
	}
}

func TestGenerateSampleConfig(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Generate sample config
	err := GenerateSampleConfig()
	if err != nil {
		t.Fatalf("GenerateSampleConfig failed: %v", err)
	}

	// Verify config file was created
	configPath := filepath.Join(tmpDir, ".config", "markin", ".markin.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("Sample config file was not created at %s", configPath)
	}

	// Read and verify config contents
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read sample config file: %v", err)
	}

	expectedContent := `markdown_dir: $VAULT_MAIN
markdown_file: daily.md
section: "## Tasks"
position: after-heading
create_section_if_missing: true`

	if string(content) != expectedContent {
		t.Errorf("Sample config content mismatch:\nGot:      %s\nExpected: %s", string(content), expectedContent)
	}
}

func TestGenerateSampleConfigWithExistingFile(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create config directory and file
	configDir := filepath.Join(tmpDir, ".config", "markin")
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	configPath := filepath.Join(configDir, ".markin.yaml")
	if err := os.WriteFile(configPath, []byte("existing content"), 0644); err != nil {
		t.Fatalf("Failed to create existing config file: %v", err)
	}

	// Generate sample config
	err := GenerateSampleConfig()
	if err != nil {
		t.Fatalf("GenerateSampleConfig failed: %v", err)
	}

	// Verify config file was overwritten
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read sample config file: %v", err)
	}

	expectedContent := `markdown_dir: $VAULT_MAIN
markdown_file: daily.md
section: "## Tasks"
position: after-heading
create_section_if_missing: true`

	if string(content) != expectedContent {
		t.Errorf("Sample config content mismatch:\nGot:      %s\nExpected: %s", string(content), expectedContent)
	}
}
