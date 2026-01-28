package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall launchd service",
	Long:  "Remove oooooi from launchd and stop auto-start on login.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Unload and remove launchd plist
		fmt.Println("Uninstalling launchd service...")
		fmt.Println("oooooi will no longer start automatically.")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
