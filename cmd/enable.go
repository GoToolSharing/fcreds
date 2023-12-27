package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enable called")
		enableFcreds()
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}

func enableFcreds() {
	// The script that should be appended to zshrc or equivalent
	script := `
export FCREDS_ENABLED="1"
alias precmd="fcreds_prefix"
`
	// Open the ~/.zshrc or other shell profile file
	file, err := os.OpenFile("/root/.zshrc", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	// Append the script to the file
	if _, err = file.WriteString(script); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	// Reload the shell profile
	cmd := exec.Command("source", "/root/.zshrc")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error reloading shell profile: ", err)
		return
	}

	fmt.Println("fcreds enabled successfully.")
}
