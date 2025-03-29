package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// expandPath expands environment variables in a path
func expandPath(path string) string {
	return os.ExpandEnv(path)
}

// debugPrint prints debug information only when debug is enabled
func debugPrint(debug bool, format string, args ...any) {
	if debug {
		fmt.Printf(format, args...)
	}
}

// AddLine adds a line into a specific section of a markdown file
func AddLine(projectDir, dailyNotePath, dailyNoteName, section, line, position string, createSectionIfMissing, debug bool) error {
	if line == "" {
		return nil
	}

	// Expand environment variables in paths
	projectDir = expandPath(projectDir)
	dailyNotePath = expandPath(dailyNotePath)
	dailyNoteName = expandPath(dailyNoteName)

	// Construct the full path to the daily note
	fullPath := filepath.Join(projectDir, dailyNotePath, dailyNoteName)

	debugPrint(debug, "Debug: Expanded paths:\n")
	debugPrint(debug, "  Project dir: %s\n", projectDir)
	debugPrint(debug, "  Daily note path: %s\n", dailyNotePath)
	debugPrint(debug, "  Daily note name: %s\n", dailyNoteName)
	debugPrint(debug, "  Full path: %s\n", fullPath)

	// Check if path still contains unexpanded environment variables
	if strings.Contains(fullPath, "$") {
		return fmt.Errorf("path contains unexpanded environment variables. Please use absolute paths or expand the variables:\n  Project dir: %s\n  Daily note path: %s\n  Daily note name: %s\n  Full path: %s",
			projectDir, dailyNotePath, dailyNoteName, fullPath)
	}

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		debugPrint(debug, "Debug: File does not exist, creating it\n")
		if err := createNewFile(fullPath, section, line, debug); err != nil {
			return err
		}
		return nil
	}

	// Read file content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	// Check if section exists
	if !strings.Contains(string(content), section) {
		debugPrint(debug, "Debug: Section not found, creating it\n")
		if createSectionIfMissing {
			if err := appendSection(fullPath, section, line); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("section '%s' not found in file at %s and create_section_if_missing is false", section, fullPath)
	}

	// Add line in the appropriate position
	if err := addLineInSection(fullPath, section, line, position); err != nil {
		return err
	}

	return nil
}

// createNewFile creates a new markdown file with the given section and line
func createNewFile(fullPath, section, line string, debug bool) error {
	debugPrint(debug, "Debug: Creating directory structure for: %s\n", filepath.Dir(fullPath))
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}

	content := fmt.Sprintf("%s\n%s\n", section, line)
	debugPrint(debug, "Debug: Writing content to file: %s\n", fullPath)
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}

// appendSection appends a new section with the given line to the end of the file
func appendSection(fullPath, section, line string) error {
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	// Add a newline if the file doesn't end with one
	contentStr := string(content)
	if len(contentStr) > 0 && !strings.HasSuffix(contentStr, "\n") {
		contentStr += "\n"
	}

	// Add an extra newline if the file doesn't end with two newlines
	if !strings.HasSuffix(contentStr, "\n\n") {
		contentStr += "\n"
	}

	// Append the new section and line
	contentStr += fmt.Sprintf("%s\n%s\n", section, line)

	if err := os.WriteFile(fullPath, []byte(contentStr), 0644); err != nil {
		return err
	}

	return nil
}

// addLineInSection adds a line into an existing section
func addLineInSection(fullPath, section, line, position string) error {
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	// Normalize line endings and split into lines
	contentStr := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(contentStr, "\n")

	// Remove trailing empty lines from input
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	var newLines []string
	sectionFound := false
	endOfSection := false
	lineAdded := false

	for i, currentLine := range lines {
		currentLine = strings.TrimRight(currentLine, "\r\n")
		if strings.HasPrefix(currentLine, "## ") {
			if sectionFound {
				endOfSection = true
				// Add a newline before the next section if needed
				if len(newLines) > 0 && len(strings.TrimSpace(newLines[len(newLines)-1])) > 0 {
					newLines = append(newLines, "")
				}
			}
			if strings.TrimSpace(currentLine) == strings.TrimSpace(section) {
				sectionFound = true
				newLines = append(newLines, currentLine)
				if position == "after-heading" {
					newLines = append(newLines, line)
					lineAdded = true
				}
				continue
			}
		}

		if sectionFound && !endOfSection && position == "before-end" && !lineAdded {
			// Check if we're at the end of the section
			isLastLine := i == len(lines)-1
			nextLineIsSection := !isLastLine && strings.HasPrefix(lines[i+1], "## ")

			if isLastLine || nextLineIsSection {
				if len(strings.TrimSpace(currentLine)) > 0 {
					newLines = append(newLines, currentLine)
					newLines = append(newLines, line)
				} else {
					newLines = append(newLines, line)
				}
				lineAdded = true
				continue
			}
		}

		newLines = append(newLines, currentLine)
	}

	if !lineAdded && sectionFound {
		if len(newLines) > 0 && len(strings.TrimSpace(newLines[len(newLines)-1])) > 0 {
			newLines = append(newLines, "")
		}
		newLines = append(newLines, line)
	}

	// Clean up multiple consecutive newlines and ensure proper spacing
	var cleanLines []string
	lastLineEmpty := false
	lastLineWasSection := false

	for _, l := range newLines {
		l = strings.TrimRight(l, "\r\n")
		isSection := strings.HasPrefix(l, "## ")

		if l == "" {
			if !lastLineEmpty && !lastLineWasSection {
				cleanLines = append(cleanLines, l)
			}
			lastLineEmpty = true
		} else {
			if isSection && len(cleanLines) > 0 && !lastLineEmpty {
				cleanLines = append(cleanLines, "")
			}
			cleanLines = append(cleanLines, l)
			lastLineEmpty = false
		}
		lastLineWasSection = isSection
	}

	// Remove trailing empty lines
	for len(cleanLines) > 0 && cleanLines[len(cleanLines)-1] == "" {
		cleanLines = cleanLines[:len(cleanLines)-1]
	}

	// Add a newline at the end if the file doesn't end with a section
	if len(cleanLines) > 0 {
		lastLine := cleanLines[len(cleanLines)-1]
		if strings.HasPrefix(lastLine, "- ") {
			// If the last line is a note, don't add a newline
		} else if !strings.HasPrefix(lastLine, "## ") {
			// If the last line is not a section header or a note, add a newline
			cleanLines = append(cleanLines, "")
		}
	}

	// Join lines with a newline and ensure the file ends with exactly one newline
	finalContent := strings.Join(cleanLines, "\n")
	if !strings.HasSuffix(finalContent, "\n") {
		finalContent += "\n"
	}

	// Normalize line endings one last time
	finalContent = strings.ReplaceAll(finalContent, "\r\n", "\n")
	finalContent = strings.TrimRight(finalContent, "\r\n") + "\n"

	return os.WriteFile(fullPath, []byte(finalContent), 0644)
}
