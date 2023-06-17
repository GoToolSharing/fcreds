package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a table of linked tools",
	Long:  `Displays the list of tools that work with fzf-creds`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tools linked to fzf-creds")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
