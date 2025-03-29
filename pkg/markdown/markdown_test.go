package markdown

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddLine(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test adding a line into an existing section
	content := `# Test File

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: Existing note

## Other Section
- Other note
`
	filePath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "- ⚡ *06:33:45 pm:* **Fleeting**:: New note", "after-heading", false, false)
	if err != nil {
		t.Errorf("Failed to add line: %v", err)
	}

	// Verify the content
	updatedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	expected := []byte(`# Test File

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: New note
- ⚡ *06:33:45 pm:* **Fleeting**:: Existing note

## Other Section
- Other note
`)
	if string(updatedContent) != string(expected) {
		t.Errorf("Content mismatch.\nExpected:\n%q\nGot:\n%q", expected, updatedContent)
	}
}

func TestAddLineEmptyLine(t *testing.T) {
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test adding an empty line
	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "", "after-heading", false, false)
	if err != nil {
		t.Errorf("Failed to add empty line: %v", err)
	}
}

func TestAddLineEmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test adding to an empty file
	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "- ⚡ *06:33:45 pm:* **Fleeting**:: New note", "after-heading", false, false)
	if err != nil {
		t.Errorf("Failed to add to empty file: %v", err)
	}

	// Verify the content
	filePath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expected := []byte(`## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: New note
`)
	if string(content) != string(expected) {
		t.Errorf("Content mismatch.\nExpected:\n%q\nGot:\n%q", expected, content)
	}
}

func TestAddLineEndOfSection(t *testing.T) {
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test adding at the end of a section
	content := `# Test File

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: Existing note

## Other Section
- Other note
`
	filePath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "- ⚡ *06:33:45 pm:* **Fleeting**:: New note", "before-end", false, false)
	if err != nil {
		t.Errorf("Failed to add line: %v", err)
	}

	// Verify the content
	updatedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	// Compare byte by byte
	expected := []byte(`# Test File

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: Existing note
- ⚡ *06:33:45 pm:* **Fleeting**:: New note

## Other Section
- Other note
`)

	if string(updatedContent) != string(expected) {
		t.Errorf("Content mismatch.\nExpected:\n%q\nGot:\n%q", expected, updatedContent)
	}
}

func TestAddLineEndOfFile(t *testing.T) {
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test adding at the end of the file
	content := `# Test File

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: Existing note
`
	filePath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "- ⚡ *06:33:45 pm:* **Fleeting**:: New note", "before-end", false, false)
	if err != nil {
		t.Errorf("Failed to add line: %v", err)
	}

	// Verify the content
	updatedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	// Compare byte by byte
	expected := []byte(`# Test File

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: Existing note
- ⚡ *06:33:45 pm:* **Fleeting**:: New note
`)

	if string(updatedContent) != string(expected) {
		t.Errorf("Content mismatch.\nExpected:\n%q\nGot:\n%q", expected, updatedContent)
	}
}

func TestAddLineCreateSectionIfMissing(t *testing.T) {
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test creating a new section when it doesn't exist
	content := `# Test File

## Other Section
- Other note
`
	filePath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "- ⚡ *06:33:45 pm:* **Fleeting**:: New note", "after-heading", true, false)
	if err != nil {
		t.Errorf("Failed to add line: %v", err)
	}

	// Verify the content
	updatedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	expected := []byte(`# Test File

## Other Section
- Other note

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: New note
`)
	if string(updatedContent) != string(expected) {
		t.Errorf("Content mismatch.\nExpected:\n%q\nGot:\n%q", expected, updatedContent)
	}
}

func TestAddLineCreateSectionIfMissingAtEnd(t *testing.T) {
	tmpDir := t.TempDir()
	projectDir := filepath.Join(tmpDir, "project")
	dailyNotePath := "notes"
	dailyNoteName := "test.md"
	section := "## 💡 🧠 🔥 Fleeting Ideas"

	// Test creating a new section at the end of the file
	content := `# Test File

## Other Section
- Other note
`
	filePath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	err := AddLine(projectDir, dailyNotePath, dailyNoteName, section, "- ⚡ *06:33:45 pm:* **Fleeting**:: New note", "after-heading", true, false)
	if err != nil {
		t.Errorf("Failed to add line: %v", err)
	}

	// Verify the content
	updatedContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	expected := []byte(`# Test File

## Other Section
- Other note

## 💡 🧠 🔥 Fleeting Ideas
- ⚡ *06:33:45 pm:* **Fleeting**:: New note
`)
	if string(updatedContent) != string(expected) {
		t.Errorf("Content mismatch.\nExpected:\n%q\nGot:\n%q", expected, updatedContent)
	}
}
