package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "fzf-creds",
	Long: `Interactive execution of bash commands`,
}

func Execute() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this tool on a Windows machine")
		return
	}
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	err = os.MkdirAll(config.Workspace_path, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	database.InitDB()
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
