package cmd

import (
	"fmt"
	"strings"

	"github.com/QU35T-code/fzf-creds/crackmapexec"
	"github.com/spf13/cobra"
)

var smartCmd = &cobra.Command{
	Use:   "smart",
	Short: "fzf-creds wrapper",
	Long:  `The wrapper that allows you to dynamically select variable values`,
	Run: func(cmd *cobra.Command, args []string) {
		mapping := map[string]func() string{
			Config.Variables_prefix + "DOMAIN":   crackmapexec.GetDomains,
			Config.Variables_prefix + "USERNAME": crackmapexec.GetUsernames,
			Config.Variables_prefix + "PASSWORD": crackmapexec.GetPasswords,
			Config.Variables_prefix + "TARGET":   crackmapexec.GetTargets,
		}
		fmt.Println("Command : ", args)
		for _, variable := range Config.Variables_custom_list {
			command := strings.Join(args[0:], "")
			if strings.Contains(Config.Variables_prefix+command, variable) {
				fmt.Printf("Variable : %s\n", Config.Variables_prefix+variable)
				fmt.Println("------------------------")
				data := mapping[Config.Variables_prefix+variable]()
				fmt.Println("Data returned : ", data)
				fmt.Println("------------------------")
				return
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(smartCmd)
}
