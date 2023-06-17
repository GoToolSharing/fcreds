package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "fzf-creds initialization",
	Long:  `Creation of the working directory and addition of the local database`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config :", err)
		}
		err = os.MkdirAll(config.Workspace_path, 0755)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		database.InitDB()
		fmt.Println("Local database initialized successfully (", config.Workspace_path+"/"+config.Local_database_name, ")")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
