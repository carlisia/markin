package main

import (
	"fmt"
	"os"

	"github.com/carlisia/markin/internal/commands"
	"github.com/carlisia/markin/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

func main() {
	rootCmd := &cobra.Command{
		Use:   "markin",
		Short: "Markdown CLI to insert lines into specific sections",
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/markin/.markin.yaml)")

	// Initialize config
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Add commands
	rootCmd.AddCommand(commands.NewInsertCmd(cfg))
	rootCmd.AddCommand(commands.NewInitCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
