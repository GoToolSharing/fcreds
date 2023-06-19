package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/crackmapexec"
	"github.com/spf13/cobra"
)

var smartCmd = &cobra.Command{
	Use:                   "smart [command]",
	Short:                 "fzf-creds wrapper",
	Long:                  `The wrapper that allows you to dynamically select variable values`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		mapping := map[string]func(){
			config.Prefix + "DOMAIN":   crackmapexec.GetDomains,
			config.Prefix + "USERNAME": crackmapexec.GetUsernames,
			config.Prefix + "PASSWORD": crackmapexec.GetPasswords,
			config.Prefix + "TARGET":   crackmapexec.GetTargets,
			config.Prefix + "HASH":     crackmapexec.GetHashes,
		}
		templated_command := strings.Join(args[1:], " ")
		var commandsList []string
		templated_command = args[0] + " " + templated_command
		for _, variable := range config.Variables_custom_list {
			if strings.Contains(config.Prefix+templated_command, variable) {
				commandsList = append(commandsList, config.Prefix+variable)
			}
		}

		for _, v := range commandsList {
			if v == config.Prefix+"DOMAIN" || v == config.Prefix+"TARGET" {
				mapping[v]()
			}
		}

		for _, v := range commandsList {
			if v == config.Prefix+"DOMAIN" || v == config.Prefix+"TARGET" {
				continue
			}
			mapping[v]()
		}

		commandToExecute := strings.ReplaceAll(templated_command, config.Prefix+"DOMAIN", strings.ReplaceAll(crackmapexec.GetData().Domain, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"USERNAME", strings.ReplaceAll(crackmapexec.GetData().Username, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"PASSWORD", strings.ReplaceAll(crackmapexec.GetData().Password, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"TARGET", strings.ReplaceAll(crackmapexec.GetData().Target, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"HASH", strings.ReplaceAll(crackmapexec.GetData().Hash, "\n", ""))
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
	smartCmd.Flags().SetInterspersed(false)
}
