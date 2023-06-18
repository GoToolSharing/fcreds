package cmd

import (
	"fmt"
	"log"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:                   "unlink [command]",
	Short:                 "Unlink a tool of fzf-creds",
	Long:                  `Remove a linked tool from the local fzf-creds database and remove its alias`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		var tools models.Tools
		result := database.DB.First(&tools, "name = ?", command)
		if result.Error != nil {
			if result.RowsAffected != 0 {
				log.Fatal(result.Error)
			}
		}

		if result.RowsAffected == 0 {
			template := utils.GetAliasTemplate(command)
			ret := utils.CheckExistingStringOnFile(config.Aliases_file_path, template)
			if ret {
				utils.RemoveLineFromFile(config.Aliases_file_path, template)
				fmt.Println("The alias has been successfully removed from the file")
				return
			}
			return
		}

		result = database.DB.Delete(&tools, "name = ?", command)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		fmt.Println("The " + command + " tool has been successfully unlinked from fzf-creds")

		template := utils.GetAliasTemplate(command)
		utils.RemoveLineFromFile(config.Aliases_file_path, template)
		fmt.Println("Remember to use the following command or open a new terminal : unalias", command)
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}
