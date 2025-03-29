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

// NewFlCmd creates a command for adding a fleeting note
func NewFlCmd(cfg *config.Config, debug bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fl [note]",
		Short: "Add a fleeting note to your daily note",
		Long: `Add a fleeting note to your daily note file.
The note will be added under the configured section.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			note := args[0]
			timestamp := time.Now().Format("03:04:05 pm")
			formattedNote := fmt.Sprintf("- ⚡ *%s:* **Fleeting**:: %s", timestamp, note)
			if err := markdown.AddLine(
				cfg.ProjectDir,
				cfg.DailyNotePath,
				cfg.DailyNoteName,
				cfg.Section,
				formattedNote,
				cfg.Position,
				cfg.CreateSectionIfMissing,
				debug,
			); err != nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("failed to add fleeting note: %w", err)
			}
			return nil
		},
	}
	return cmd
}

// NewInitCmd creates a command for initializing the configuration
func NewInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize the configuration file",
		Long: `Initialize the configuration file with default settings.
This will create a sample configuration file in your home directory.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.GenerateSampleConfig(""); err != nil {
				return fmt.Errorf("failed to generate sample configuration: %w", err)
			}
			fmt.Println("Configuration file created successfully!")
			return nil
		},
	}
}
