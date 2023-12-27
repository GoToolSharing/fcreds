package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/core"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:  "fzf-creds",
	Long: `Interactive execution of bash commands`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			fmt.Println("Can't execute this tool on a Windows machine for the moment")
			return
		}
		err := config.ConfigureLogger()
		if err != nil {
			config.GlobalConfig.Logger.Error("", zap.Error(err))
			os.Exit(1)
		}
		config.GlobalConfig.Logger.Debug(fmt.Sprintf("Verbosity level : %v", config.GlobalConfig.Verbose))
		err = config.Init()
		if err != nil {
			config.GlobalConfig.Logger.Error("", zap.Error(err))
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			core.Start()
		}
	},
}

func Execute() {
	// database.InitDB() // TODO : check if necessary
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().CountVarP(&config.GlobalConfig.Verbose, "verbose", "v", "Verbose level")
	// rootCmd.CompletionOptions.DisableDefaultCmd = true
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
}

// scanner := bufio.NewScanner(os.Stdin)

// 	// choicesFile := "path/to/usernames.txt"

// 	for scanner.Scan() {
// 		input := scanner.Text()

// 		// Vérifiez si $USERNAME est dans la chaîne entrée
// 		// if strings.Contains(input, "$USERNAME") {
// 		// 	// Lance fzf avec le fichier des choix comme entrée
// 		// 	cmd := exec.Command("fzf", "--height=40%", "--reverse")
// 		// 	cmd.Stdin, _ = os.Open(choicesFile)
// 		// 	cmd.Stderr = os.Stderr
// 		// 	output, err := cmd.Output()
// 		// 	if err != nil {
// 		// 		fmt.Fprintln(os.Stderr, "erreur de fzf:", err)
// 		// 		continue
// 		// 	}

// 		// 	// Remplace $USERNAME par le choix de fzf dans la chaîne entrée
// 		// 	replacement := strings.TrimSpace(string(output))
// 		// 	input = strings.ReplaceAll(input, "$USERNAME", replacement)
// 		// }

// 		fmt.Println(input)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "erreur de lecture:", err)
// 	}

// # Fonction pour préfixer les commandes avec "go run . wrapper"
// fzf_creds_wrapper() {
//     # Vérifier si la commande est déjà préfixée
//     if [[ $2 != "go run .*" ]]; then
//         # Échapper les guillemets simples dans la commande
//         local escaped_command="${2//\'/\'\\\'\'}"
//         # Encapsuler la commande échappée entre guillemets simples
//         local cmd_quoted="'$escaped_command'"
//         # Empêcher la commande originale de s'exécuter
//         eval "go run . wrapper $cmd_quoted"
//         return 1
//     fi
// }

// # Ajoutez la fonction à la liste des fonctions de pré-execution
// preexec_functions+=(fzf_creds_wrapper)
