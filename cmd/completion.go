package cmd

import (
	"fmt"
	"os"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:                   "completion [enable|disable]",
	Short:                 "Manage autocompletion command",
	Long:                  `Autocompletion allows you to type TAB and see the options available for your commands`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"enable", "disable"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch os.Getenv("SHELL") {
		case "/bin/bash":
			cmd.Root().GenZshCompletionFile(config.Completion_file_path)
		case "/bin/zsh":
			cmd.Root().GenZshCompletionFile(config.Completion_file_path)
		case "/bin/fish":
			cmd.Root().GenZshCompletionFile(config.Completion_file_path)
		}
		template := "source " + config.Completion_file_path
		ret := utils.CheckExistingStringOnFile(config.GetRCFilePath(), template)
		if !ret && args[0] == "enable" {
			utils.AppendToFile(config.GetRCFilePath(), template)
			fmt.Println("Completion has been successfully activated")
			fmt.Println("Don't forget to source the completion file -> source " + config.Completion_file_path)
			return
		} else if ret && args[0] == "disable" {
			utils.RemoveLineFromFile(config.GetRCFilePath(), template)
			fmt.Println("Completion has been disabled")
			return
		} else if ret && args[0] == "enable" {
			fmt.Println("Completion is already enable")
			return
		} else {
			fmt.Println("Completion is already disable")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
