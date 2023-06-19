package cmd

import (
	"fmt"
	"os"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:                   "options",
	Short:                 "Displays the list of supported variables",
	Long:                  `In a table, displays the variables compatible with fzf-creds`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compatible options")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Variable", "Description"})
		t.AppendRows([]table.Row{{config.Prefix + "DOMAIN", "The target domain"}})
		t.AppendRows([]table.Row{{config.Prefix + "USERNAME", "The target user"}})
		t.AppendRows([]table.Row{{config.Prefix + "PASSWORD", "The target password"}})
		t.AppendRows([]table.Row{{config.Prefix + "TARGET", "The target IP"}})
		t.AppendRows([]table.Row{{config.Prefix + "HASH", "The target NT HASH"}})
		t.AppendSeparator()
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(optionsCmd)
}
