package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Displays the list of supported variables",
	Long:  `In a table, displays the variables compatible with fzf-creds`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compatible options")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Variable", "Description"})
		t.AppendRows([]table.Row{{Config.Variables_prefix + "DOMAIN", "The target domain"}})
		t.AppendRows([]table.Row{{Config.Variables_prefix + "USERNAME", "The target user"}})
		t.AppendRows([]table.Row{{Config.Variables_prefix + "PASSWORD", "The target password"}})
		t.AppendRows([]table.Row{{Config.Variables_prefix + "TARGET", "The target IP"}})
		t.AppendSeparator()
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(optionsCmd)
}
