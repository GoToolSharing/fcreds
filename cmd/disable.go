package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disable called")
		disableFcreds()
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}

func disableFcreds() {
	// Define the path to the zshrc or other profile file
	profilePath := "/root/.zshrc"

	// Read the contents of the file
	content, err := os.ReadFile(profilePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// Convert the content to a string and remove the lines added by fcreds_enable
	lines := strings.Split(string(content), "\n")
	var newLines []string
	for _, line := range lines {
		if !strings.Contains(line, "FCREDS_ENABLED=\"1\"") && !strings.Contains(line, "alias precmd") {
			newLines = append(newLines, line)
		}
	}

	// Write the new content back to the file
	err = os.WriteFile(profilePath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

	fmt.Println("fcreds disabled successfully.")
}

// unset -f preexec
