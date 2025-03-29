package main

import (
	"fmt"
	"os"

	"github.com/carlisia/markin/internal/commands"
	"github.com/carlisia/markin/internal/config"
	"github.com/spf13/cobra"
)

var debug bool

func main() {
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "markin",
		Short: "A CLI tool for managing markdown notes",
		Long: `Markin is a CLI tool for managing markdown notes.
It provides commands for adding different types of notes to markdown files.`,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")

	rootCmd.AddCommand(commands.NewFlCmd(cfg, debug))
	rootCmd.AddCommand(commands.NewInitCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
