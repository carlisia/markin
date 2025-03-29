package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// InsertLine inserts a line into a specific section of a markdown file
func InsertLine(markdownDir, markdownFile, section, line, position string, createSectionIfMissing bool) error {
	if line == "" {
		return nil
	}

	fullPath := filepath.Join(markdownDir, markdownFile)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if err := createNewFile(fullPath, markdownDir, section, line); err != nil {
			return fmt.Errorf("error creating new file: %w", err)
		}
		return nil
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return fmt.Errorf("error reading markdown file: %w", err)
	}

	// Handle empty file
	if len(strings.TrimSpace(string(content))) == 0 {
		if err := createNewFile(fullPath, markdownDir, section, line); err != nil {
			return fmt.Errorf("error creating new file: %w", err)
		}
		return nil
	}

	lines := strings.Split(string(content), "\n")
	newLines, err := processLines(lines, section, line, position, createSectionIfMissing)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fullPath, []byte(strings.Join(newLines, "\n")), 0644); err != nil {
		return fmt.Errorf("error writing markdown file: %w", err)
	}

	return nil
}

func createNewFile(fullPath, markdownDir, section, line string) error {
	if err := os.MkdirAll(markdownDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directories: %w", err)
	}

	emptyContent := fmt.Sprintf("%s\n%s\n", section, line)
	if err := os.WriteFile(fullPath, []byte(emptyContent), 0644); err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	return nil
}

func processLines(lines []string, section, line, position string, createSectionIfMissing bool) ([]string, error) {
	var newLines []string
	inserted := false
	inSection := false

	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		newLines = append(newLines, currentLine)

		if strings.TrimSpace(currentLine) == section {
			inSection = true
			if position == "after-heading" && !inserted {
				newLines = append(newLines, line)
				inserted = true
			}
			continue
		}

		if inSection && strings.HasPrefix(currentLine, "##") && strings.TrimSpace(currentLine) != section {
			if position == "end-of-section" && !inserted {
				newLines = append(newLines[:len(newLines)-1], line, newLines[len(newLines)-1])
				inserted = true
			}
			inSection = false
		}
	}

	// Handle end of file case
	if inSection && position == "end-of-section" && !inserted {
		newLines = append(newLines, line)
		inserted = true
	}

	// Handle missing section case
	if !inserted && createSectionIfMissing {
		// Add a newline before the new section if the file isn't empty
		if len(newLines) > 0 && newLines[len(newLines)-1] != "" {
			newLines = append(newLines, "")
		}
		newLines = append(newLines, section, line)
	}

	return newLines, nil
}
