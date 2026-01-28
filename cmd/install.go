package cmd

import (
	"fmt"
	"os"

	"github.com/knwoop/ooi/internal/launchd"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install launchd service for auto-start",
	Long:  "Register ooi as a launchd service to start automatically on login.",
	Run: func(cmd *cobra.Command, args []string) {
		if launchd.IsInstalled() {
			fmt.Println("Service is already installed. Run 'ooi reinstall' to update.")
			return
		}

		fmt.Println("Installing launchd service...")

		if err := launchd.Install(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to install: %v\n", err)
			os.Exit(1)
		}

		plistPath, _ := launchd.PlistPath()
		fmt.Println("Service installed successfully!")
		fmt.Printf("Plist: %s\n", plistPath)
		fmt.Println("ooi will now start automatically on login.")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
