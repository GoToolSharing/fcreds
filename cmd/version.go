package cmd

import (
	"fmt"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Displays the version of the fzf-creds tool`,
	Run: func(cmd *cobra.Command, args []string) {
		config.GlobalConfig.Logger.Info("Version command executed")
		config.GlobalConfig.Logger.Debug(fmt.Sprintf("config.Version: %s", config.Version))
		var message string
		if len(config.Version) == 40 {
			message = fmt.Sprintf("Development version (dev branch): %s", config.Version)
		} else {
			message = fmt.Sprintf("Stable version (main branch): %s", config.Version)
		}

		fmt.Println(message)
		config.GlobalConfig.Logger.Info("Exit version command correctly")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
