package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/QU35T-code/fzf-creds/crackmapexec"
	"github.com/spf13/cobra"
)

var smartCmd = &cobra.Command{
	Use:   "smart",
	Short: "fzf-creds wrapper",
	Long:  `The wrapper that allows you to dynamically select variable values`,
	Run: func(cmd *cobra.Command, args []string) {
		mapping := map[string]func(){
			Config.Variables_prefix + "DOMAIN":   crackmapexec.GetDomains,
			Config.Variables_prefix + "USERNAME": crackmapexec.GetUsernames,
			Config.Variables_prefix + "PASSWORD": crackmapexec.GetPasswords,
			Config.Variables_prefix + "TARGET":   crackmapexec.GetTargets,
		}
		templated_command := strings.Join(args[1:], "")
		templated_command = args[0] + " " + templated_command
		for _, variable := range Config.Variables_custom_list {
			if strings.Contains(Config.Variables_prefix+templated_command, variable) {
				mapping[Config.Variables_prefix+variable]()
			}
		}
		commandToExecute := strings.ReplaceAll(templated_command, Config.Variables_prefix+"DOMAIN", strings.ReplaceAll(crackmapexec.GetData().Domain, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, Config.Variables_prefix+"USERNAME", strings.ReplaceAll(crackmapexec.GetData().Username, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, Config.Variables_prefix+"PASSWORD", strings.ReplaceAll(crackmapexec.GetData().Password, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, Config.Variables_prefix+"TARGET", strings.ReplaceAll(crackmapexec.GetData().Target, "\n", ""))
		fmt.Println(commandToExecute)
		splited_command := strings.Split(commandToExecute, " ")
		execMe := exec.Command(splited_command[0], splited_command[1:]...)

		execMe.Stdin = os.Stdin
		execMe.Stdout = os.Stdout
		execMe.Stderr = os.Stderr

		err := execMe.Start()
		if err != nil {
			log.Fatal(err)
		}
		err = execMe.Wait()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(smartCmd)
}
