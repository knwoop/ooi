package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/knwoop/oooooi/internal/calendar"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Google Calendar",
	Long:  "Start OAuth 2.0 flow to authenticate with Google Calendar API.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		configDir, err := calendar.ConfigDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Config directory: %s\n", configDir)
		fmt.Println("Starting Google OAuth authentication...")

		token, err := calendar.Authenticate(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Authentication failed: %v\n", err)
			os.Exit(1)
		}

		if err := calendar.SaveToken(token); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to save token: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Authentication successful! Token saved.")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
