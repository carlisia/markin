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

// NewInsertCmd creates the insert command
func NewInsertCmd(cfg *config.Config) *cobra.Command {
	var line string

	cmd := &cobra.Command{
		Use:   "insert [line]",
		Short: "Insert a line into a specific section of a markdown file",
		Long: `Insert a line into a specific section of a markdown file.
If no line is provided, no insertion will be performed.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				line = args[0]
			}

			if line == "" {
				fmt.Println("No line provided. Skipping insertion.")
				return nil
			}

			// Format current time
			now := time.Now()
			timeStr := now.Format("03:04:05 pm")
			formattedLine := fmt.Sprintf("- ⚡ *%s:* **Fleeting**:: %s", timeStr, line)

			if err := markdown.InsertLine(
				cfg.MarkdownDir,
				cfg.MarkdownFile,
				cfg.Section,
				formattedLine,
				cfg.Position,
				cfg.CreateSectionIfMissing,
			); err != nil {
				return fmt.Errorf("failed to insert line: %w", err)
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
