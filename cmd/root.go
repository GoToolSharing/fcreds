package cmd

import (
	"fmt"
	"os"
	"runtime"

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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
