package cmd

import (
	"fmt"
	"os"

	"github.com/knwoop/ooi/internal/launchd"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall launchd service",
	Long:  "Remove ooi from launchd and stop auto-start on login.",
	Run: func(cmd *cobra.Command, args []string) {
		if !launchd.IsInstalled() {
			fmt.Println("Service is not installed.")
			return
		}

		fmt.Println("Uninstalling launchd service...")

		if err := launchd.Uninstall(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to uninstall: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Service uninstalled successfully!")
		fmt.Println("ooi will no longer start automatically.")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
