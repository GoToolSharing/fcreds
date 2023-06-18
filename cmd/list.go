package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a table of linked tools",
	Long:  `Displays the list of tools that work with fzf-creds`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tools linked to fzf-creds")
		var tools []models.Tools
		result := database.DB.Find(&tools)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Tool", "Alias"})
		for _, tool := range tools {
			template := utils.GetAliasTemplate(tool.Name)
			ret := utils.CheckExistingStringOnFile(config.Aliases_file_path, template)
			if ret {
				t.AppendRows([]table.Row{{tool.Name, template}})
				continue
			}
		}
		t.AppendSeparator()
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
