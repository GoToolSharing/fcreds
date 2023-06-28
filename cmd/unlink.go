package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/database"
	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:                   "unlink [command] [...command...] [...command...]",
	Short:                 "Unlink a tool of fzf-creds",
	Long:                  `Remove a linked tool from the local fzf-creds database and remove its alias`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var toolsList []string
		var tools models.Tools
		for _, arg := range args {
			ret := utils.CheckIsFile(arg)
			if ret {
				if verbose {
					fmt.Println("[DEBUG][FILE] - Detected :", arg)
				}
				file, err := os.Open(arg)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					tool := scanner.Text()
					if verbose {
						fmt.Println("[DEBUG][FILE] - Line :", tool)
					}
					_, err := exec.LookPath(tool)
					if err == nil {
						if verbose {
							fmt.Println("[DEBUG][FILE] - Remove :", tool)
						}
						toolsList = append(toolsList, tool)
					}
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
				continue
			}
			_, err := exec.LookPath(arg)
			if err == nil {
				if verbose {
					fmt.Println("[DEBUG][TOOL] - Remove :", arg)
				}
				toolsList = append(toolsList, arg)
			}
		}
		if toolsList == nil {
			fmt.Println("No valid tool was specified")
			return
		}
		for _, tool := range toolsList {
			result := database.DB.First(&tools, "name = ?", tool)
			if result.Error != nil {
				if result.RowsAffected != 0 {
					log.Fatal(result.Error)
				}
			}
			template := utils.GetAliasTemplate(tool)
			ret := utils.CheckExistingStringOnFile(config.Aliases_file_path, template)
			if ret {
				utils.RemoveLineFromFile(config.Aliases_file_path, template)
				fmt.Println(tool, "has been successfully unlinked")
				result = database.DB.Delete(&tools, "name = ?", tool)
				if result.Error != nil {
					log.Fatal(result.Error)
				}
				continue
			}
		}
		fmt.Println("Remember to use the following command or open a new terminal : unalias {toolName}")
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}
