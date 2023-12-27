/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/lib/crackmapexec"
	"github.com/spf13/cobra"
)

// wrapperCmd represents the wrapper command
var wrapperCmd = &cobra.Command{
	Use:   "wrapper",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		mapping := map[string]func(){
			config.Prefix + "DOMAIN":   crackmapexec.GetDomains,
			config.Prefix + "USERNAME": crackmapexec.GetUsernames,
			config.Prefix + "PASSWORD": crackmapexec.GetPasswords,
			config.Prefix + "TARGET":   crackmapexec.GetTargets,
			config.Prefix + "HASH":     crackmapexec.GetHashes,
		}
		command := strings.Join(args[0:], " ")
		fmt.Println("Start command")
		// fmt.Println(command)
		var commandsList []string
		// command = args[0] + " " + command
		for _, variable := range config.VariablesCustomList {
			if strings.Contains(config.Prefix+command, variable) {
				commandsList = append(commandsList, config.Prefix+variable)
			}
		}

		for _, v := range commandsList {
			mapping[v]()
		}

		// for _, v := range commandsList {
		// 	if v == config.Prefix+"DOMAIN" || v == config.Prefix+"TARGET" {
		// 		continue
		// 	}
		// 	mapping[v]()
		// }
		commandToExecute := strings.ReplaceAll(command, config.Prefix+"DOMAIN", strings.ReplaceAll(crackmapexec.GetData().Domain, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"USERNAME", strings.ReplaceAll(crackmapexec.GetData().Username, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"PASSWORD", strings.ReplaceAll(crackmapexec.GetData().Password, "\n", ""))
		commandToExecute = strings.ReplaceAll(commandToExecute, config.Prefix+"TARGET", strings.ReplaceAll(crackmapexec.GetData().Target, "\n", ""))
		// fmt.Println(commandToExecute)
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
	rootCmd.AddCommand(wrapperCmd)
}
