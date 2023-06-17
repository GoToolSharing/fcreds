package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resetting fzf-creds",
	Long:  `Removal of the local database and the working directory of the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reset called")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
