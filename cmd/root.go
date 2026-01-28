package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oooooi",
	Short: "Meeting reminder CLI tool",
	Long:  "oooooi is a macOS CLI tool that automatically opens Google Meet 1 minute before meetings.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Start daemon/scheduler
		fmt.Println("Starting oooooi daemon...")
		fmt.Println("Watching for upcoming meetings...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
