package markdown

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInsertLine(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	line := "- Test item"
	position := "after-heading"

	// Test inserting a line
	err := InsertLine(markdownDir, markdownFile, section, line, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	// Verify file was created
	fullPath := filepath.Join(markdownDir, markdownFile)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}

	// Verify section and line format
	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		t.Fatalf("Expected at least 2 lines, got %d", len(lines))
	}
	if lines[0] != section {
		t.Errorf("First line should be section, got: %s", lines[0])
	}
	if !strings.Contains(lines[1], "Fleeting") {
		t.Errorf("Second line should contain 'Fleeting', got: %s", lines[1])
	}

	// Test inserting another line
	newLine := "- Another item"
	err = InsertLine(markdownDir, markdownFile, section, newLine, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	content, err = os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	lines = strings.Split(string(content), "\n")
	if len(lines) < 3 {
		t.Fatalf("Expected at least 3 lines, got %d", len(lines))
	}
	if lines[0] != section {
		t.Errorf("First line should be section, got: %s", lines[0])
	}
	if !strings.Contains(lines[1], "Fleeting") {
		t.Errorf("Second line should contain 'Fleeting', got: %s", lines[1])
	}
	if !strings.Contains(lines[2], "Fleeting") {
		t.Errorf("Third line should contain 'Fleeting', got: %s", lines[2])
	}
}

func TestInsertLineEmptyLine(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	position := "after-heading"

	// Test inserting an empty line
	err := InsertLine(markdownDir, markdownFile, section, "", position, true)
	if err != nil {
		t.Fatalf("InsertLine failed with empty line: %v", err)
	}

	// Verify file was not created
	fullPath := filepath.Join(markdownDir, markdownFile)
	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		t.Error("File should not be created when inserting empty line")
	}
}

func TestInsertLineEmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	line := "- Test item"
	position := "after-heading"

	// Create empty file
	fullPath := filepath.Join(markdownDir, markdownFile)
	if err := os.MkdirAll(markdownDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(fullPath, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create empty file: %v", err)
	}

	// Test inserting into empty file
	err := InsertLine(markdownDir, markdownFile, section, line, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	// Verify content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		t.Fatalf("Expected at least 2 lines, got %d", len(lines))
	}
	if lines[0] != section {
		t.Errorf("First line should be section, got: %s", lines[0])
	}
	if !strings.Contains(lines[1], "Fleeting") {
		t.Errorf("Second line should contain 'Fleeting', got: %s", lines[1])
	}
}

func TestInsertLineEndOfSection(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	line := "- Test item"
	position := "end-of-section"

	// Create initial content
	initialContent := section + "\n- Existing item\n## Next Section\n"
	fullPath := filepath.Join(markdownDir, markdownFile)
	if err := os.MkdirAll(markdownDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(fullPath, []byte(initialContent), 0644); err != nil {
		t.Fatalf("Failed to write initial content: %v", err)
	}

	// Insert line at end of section
	err := InsertLine(markdownDir, markdownFile, section, line, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	// Verify content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 4 {
		t.Fatalf("Expected at least 4 lines, got %d", len(lines))
	}
	if lines[0] != section {
		t.Errorf("First line should be section, got: %s", lines[0])
	}
	if lines[1] != "- Existing item" {
		t.Errorf("Second line should be existing item, got: %s", lines[1])
	}
	if !strings.Contains(lines[2], "Fleeting") {
		t.Errorf("Third line should contain 'Fleeting', got: %s", lines[2])
	}
	if lines[3] != "## Next Section" {
		t.Errorf("Fourth line should be next section, got: %s", lines[3])
	}
}

func TestInsertLineEndOfFile(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	line := "- Test item"
	position := "end-of-section"

	// Create initial content without next section
	initialContent := section + "\n- Existing item\n"
	fullPath := filepath.Join(markdownDir, markdownFile)
	if err := os.MkdirAll(markdownDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(fullPath, []byte(initialContent), 0644); err != nil {
		t.Fatalf("Failed to write initial content: %v", err)
	}

	// Insert line at end of file
	err := InsertLine(markdownDir, markdownFile, section, line, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	// Verify content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 3 {
		t.Fatalf("Expected at least 3 lines, got %d", len(lines))
	}
	if lines[0] != section {
		t.Errorf("First line should be section, got: %s", lines[0])
	}
	if lines[1] != "- Existing item" {
		t.Errorf("Second line should be existing item, got: %s", lines[1])
	}
	if !strings.Contains(lines[2], "Fleeting") {
		t.Errorf("Third line should contain 'Fleeting', got: %s", lines[2])
	}
}

func TestInsertLineCreateSectionIfMissing(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	line := "- Test item"
	position := "after-heading"

	// Create initial content without the section
	initialContent := "## Other Section\n- Existing item\n"
	fullPath := filepath.Join(markdownDir, markdownFile)
	if err := os.MkdirAll(markdownDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(fullPath, []byte(initialContent), 0644); err != nil {
		t.Fatalf("Failed to write initial content: %v", err)
	}

	// Insert line with createSectionIfMissing=true
	err := InsertLine(markdownDir, markdownFile, section, line, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	// Verify content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 4 {
		t.Fatalf("Expected at least 4 lines, got %d", len(lines))
	}
	if lines[0] != "## Other Section" {
		t.Errorf("First line should be other section, got: %s", lines[0])
	}
	if lines[1] != "- Existing item" {
		t.Errorf("Second line should be existing item, got: %s", lines[1])
	}
	if lines[2] != "" {
		t.Errorf("Third line should be empty, got: %s", lines[2])
	}
	if lines[3] != section {
		t.Errorf("Fourth line should be section, got: %s", lines[3])
	}
	if !strings.Contains(lines[4], "Fleeting") {
		t.Errorf("Fifth line should contain 'Fleeting', got: %s", lines[4])
	}
}

func TestInsertLineCreateSectionIfMissingAtEnd(t *testing.T) {
	tmpDir := t.TempDir()
	markdownDir := filepath.Join(tmpDir, "markdown")
	markdownFile := "test.md"
	section := "## 💭 ✍️ ✨ Notes"
	line := "- Test item"
	position := "after-heading"

	// Create initial content without the section and with trailing newline
	initialContent := "## Other Section\n- Existing item\n\n"
	fullPath := filepath.Join(markdownDir, markdownFile)
	if err := os.MkdirAll(markdownDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(fullPath, []byte(initialContent), 0644); err != nil {
		t.Fatalf("Failed to write initial content: %v", err)
	}

	// Insert line with createSectionIfMissing=true
	err := InsertLine(markdownDir, markdownFile, section, line, position, true)
	if err != nil {
		t.Fatalf("InsertLine failed: %v", err)
	}

	// Verify content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 4 {
		t.Fatalf("Expected at least 4 lines, got %d", len(lines))
	}
	if lines[0] != "## Other Section" {
		t.Errorf("First line should be other section, got: %s", lines[0])
	}
	if lines[1] != "- Existing item" {
		t.Errorf("Second line should be existing item, got: %s", lines[1])
	}
	if lines[2] != "" {
		t.Errorf("Third line should be empty, got: %s", lines[2])
	}
	if lines[3] != section {
		t.Errorf("Fourth line should be section, got: %s", lines[3])
	}
	if !strings.Contains(lines[4], "Fleeting") {
		t.Errorf("Fifth line should contain 'Fleeting', got: %s", lines[4])
	}
}
