package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link a tool with fzf-creds",
	Long:  `Link a tool in the fzf-creds database so that it can be used`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Bad number of arguments, got : ", len(args), ", expected : 1")
			cmd.Help()
			return
		}
		command := args[0]
		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		var tools models.Tools
		result := database.DB.First(&tools, "name = ?", command)
		if result.Error != nil {
			if result.RowsAffected != 0 {
				log.Fatal(result.Error)
			}
		}

		if result.RowsAffected != 0 {
			fmt.Println("The tool is already linked to fzf-creds")
			return
		}

		result = database.DB.Create(&models.Tools{Name: command})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config :", err)
		}
		appendToAliasesFile(config, command)

		fmt.Println("The tool has been successfully linked to fzf-creds")
		fmt.Println("Don't forget to source the aliases file -> source " + config.Aliases_file_path)
	},
}

func appendToAliasesFile(config utils.Config, command string) {
	f, err := os.OpenFile(config.Aliases_file_path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	template := "alias " + command + "='fzf-creds smart " + command + "'"
	if _, err := f.WriteString(template + "\n"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
