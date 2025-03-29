package commands

import (
	"fmt"
	"time"

	"github.com/carlisia/markin/internal/config"
	"github.com/carlisia/markin/pkg/markdown"
	"github.com/spf13/cobra"
)

const (
	cyan   = "\033[36m"
	yellow = "\033[33m"
	white  = "\033[37m"
	reset  = "\033[0m"
)

// NewFlCmd creates the fleeting note command
func NewFlCmd(cfg *config.Config) *cobra.Command {
	var note string

	cmd := &cobra.Command{
		Use:   "fl [note]",
		Short: "Add a fleeting note entry to a markdown file",
		Long: `Add a fleeting note entry to a markdown file.
A fleeting note is a quick thought or idea that you want to capture.
The note will be added as a bullet item with a timestamp.
If no note is provided, no entry will be added.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				note = args[0]
			}

			if note == "" {
				fmt.Println("No note provided. Skipping entry.")
				return nil
			}

			// Format current time
			now := time.Now()
			timeStr := now.Format("03:04:05 pm")
			bulletItem := fmt.Sprintf("- ⚡ *%s:* **Fleeting**:: %s", timeStr, note)

			if err := markdown.InsertLine(
				cfg.MarkdownDir,
				cfg.MarkdownFile,
				cfg.Section,
				bulletItem,
				cfg.Position,
				cfg.CreateSectionIfMissing,
			); err != nil {
				return fmt.Errorf("failed to add fleeting note entry: %w", err)
			}
			return nil
		},
	}

	return cmd
}

// NewInitCmd creates the init command
func NewInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Generate a sample config file in $HOME/.config/markin/.markin.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.GenerateSampleConfig(); err != nil {
				return fmt.Errorf("failed to generate sample config: %w", err)
			}
			return nil
		},
	}
}
