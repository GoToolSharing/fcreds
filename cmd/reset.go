package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:                   "reset",
	Short:                 "Resetting fzf-creds",
	Long:                  `Removal of the local database and the working directory of the tool`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := os.RemoveAll(config.Workspace_path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The database has been deleted")
		aliasTemplate := "source " + config.Aliases_file_path
		utils.RemoveLineFromFile(config.GetRCFilePath(), aliasTemplate)
		fmt.Println("The link with the RC file has been destroyed")
		fmt.Println("fzf-creds successfully reset")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
