package core

import (
	"bufio"
	"fmt"
	"os"
)

func Start() error {
	scanner := bufio.NewScanner(os.Stdin)

	// choicesFile := "path/to/usernames.txt"

	for scanner.Scan() {
		input := scanner.Text()

		// Vérifiez si $USERNAME est dans la chaîne entrée
		// if strings.Contains(input, "$USERNAME") {
		// 	// Lance fzf avec le fichier des choix comme entrée
		// 	cmd := exec.Command("fzf", "--height=40%", "--reverse")
		// 	cmd.Stdin, _ = os.Open(choicesFile)
		// 	cmd.Stderr = os.Stderr
		// 	output, err := cmd.Output()
		// 	if err != nil {
		// 		fmt.Fprintln(os.Stderr, "erreur de fzf:", err)
		// 		continue
		// 	}

		// 	// Remplace $USERNAME par le choix de fzf dans la chaîne entrée
		// 	replacement := strings.TrimSpace(string(output))
		// 	input = strings.ReplaceAll(input, "$USERNAME", replacement)
		// }

		fmt.Println(input)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading error: %v", err)
	}
	return nil
}
