package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resetting fzf-creds",
	Long:  `Removal of the local database and the working directory of the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config :", err)
		}
		err = os.RemoveAll(config.Workspace_path)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("The database has been deleted")
		fmt.Println("fzf-creds successfully reset")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
