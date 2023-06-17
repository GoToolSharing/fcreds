package cmd

import (
	"fmt"
	"log"

	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink a tool of fzf-creds",
	Long:  `Remove a linked tool from the local fzf-creds database and remove its alias`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Bad number of arguments, got : ", len(args), ", expected : 1")
			cmd.Help()
			return
		}
		command := args[0]
		var tools models.Tools
		result := database.DB.First(&tools, "name = ?", command)
		if result.Error != nil {
			if result.RowsAffected != 0 {
				log.Fatal(result.Error)
			}
		}

		if result.RowsAffected == 0 {
			fmt.Println("The tool is not linked with fzf-creds (not in the database)")
			// TODO: check anyway the aliases for the delete
			return
		}

		result = database.DB.Delete(&tools, "name = ?", command)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		fmt.Println("The " + command + " tool has been successfully unlinked from fzf-creds")

		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config :", err)
		}
		template := "alias " + command + "='fzf-creds smart " + command + "'"
		utils.RemoveLineFromFile(config.Aliases_file_path, template)
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}
