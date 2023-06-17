package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a table of linked tools",
	Long:  `Displays the list of tools that work with fzf-creds`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tools linked to fzf-creds")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Tool"})
		t.AppendRows([]table.Row{
			{"smbclient.py"},
			{"rpcclient.py"},
		})
		t.AppendSeparator()
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
