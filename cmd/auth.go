package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Google Calendar",
	Long:  "Start OAuth 2.0 flow to authenticate with Google Calendar API.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement OAuth flow
		fmt.Println("Starting Google OAuth authentication...")
		fmt.Println("Please check your browser to complete authentication.")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
