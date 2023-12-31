package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:                   "link [command] [...command...] [...command...]",
	Short:                 "Link a tool with fzf-creds",
	Long:                  `Link a tool in the fzf-creds database so that it can be used`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var toolsList []string
		var tools models.Tools
		for _, arg := range args {
			_, err := exec.LookPath(arg)
			if err == nil {
				toolsList = append(toolsList, arg)
			}
		}
		for _, tool := range toolsList {
			result := database.DB.First(&tools, "name = ?", tool)
			if result.Error != nil {
				if result.RowsAffected != 0 {
					log.Fatal(result.Error)
				}
			}

			template := utils.GetAliasTemplate(tool)
			ret := utils.CheckExistingStringOnFile(config.Aliases_file_path, template)
			if ret {
				fmt.Println(tool, "is already linked")
				continue
			}
			if result.RowsAffected != 0 {
				utils.AppendToFile(config.Aliases_file_path, tool)
				fmt.Println(tool, "has been successfully added to the aliases file")
				continue
			}

			result = database.DB.Create(&models.Tools{Name: tool})
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			utils.AppendToFile(config.Aliases_file_path, template)

			fmt.Println(tool, "has been successfully linked to fzf-creds")
		}
		fmt.Println("Don't forget to source the aliases file -> source " + config.Aliases_file_path)
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
