package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link a tool with fzf-creds",
	Long:  `Link a tool in the fzf-creds database so that it can be used`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Bad number of arguments, got : ", len(args), ", expected : 1")
			cmd.Help()
			return
		}
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

		if result.RowsAffected != 0 {
			fmt.Println("The tool is already linked to fzf-creds")
			// TODO: check the aliases anyway to add it if it is not there
			return
		}

		result = database.DB.Create(&models.Tools{Name: command})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		utils.AppendToFile(Config.Aliases_file_path, command)

		fmt.Println("The tool has been successfully linked to fzf-creds")
		fmt.Println("Don't forget to source the aliases file -> source " + Config.Aliases_file_path)
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
