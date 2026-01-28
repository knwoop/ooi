package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install launchd service for auto-start",
	Long:  "Register oooooi as a launchd service to start automatically on login.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Create and load launchd plist
		fmt.Println("Installing launchd service...")
		fmt.Println("oooooi will now start automatically on login.")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
