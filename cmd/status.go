package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show next upcoming meeting",
	Long:  "Display the next scheduled meeting with Google Meet link.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Fetch and display next meeting
		fmt.Println("Fetching next meeting...")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
