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
	Use:                   "link [command]",
	Short:                 "Link a tool with fzf-creds",
	Long:                  `Link a tool in the fzf-creds database so that it can be used`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		var tools models.Tools
		result := database.DB.First(&tools, "name = ?", command)
		if result.Error != nil {
			if result.RowsAffected != 0 {
				log.Fatal(result.Error)
			}
		}

		template := utils.GetAliasTemplate(command)
		if result.RowsAffected != 0 {
			ret := utils.CheckExistingStringOnFile(config.Aliases_file_path, template)
			if ret {
				fmt.Println("The tool is already linked")
				return
			}
			utils.AppendToFile(config.Aliases_file_path, command)
			fmt.Println("The tool has been successfully added to the aliases file")
			return
		}

		result = database.DB.Create(&models.Tools{Name: command})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		utils.AppendToFile(config.Aliases_file_path, template)

		fmt.Println("The tool has been successfully linked to fzf-creds")
		fmt.Println("Don't forget to source the aliases file -> source " + config.Aliases_file_path)
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
