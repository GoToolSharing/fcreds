package cmd

import (
	"fmt"
	"log"

	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug command",
	Long:  `Debug command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debug called")
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}
		fmt.Println(config.Variables_custom_list[1])
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
