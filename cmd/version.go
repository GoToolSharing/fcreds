package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Displays the version of the fzf-creds tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fzf-creds version dev")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
