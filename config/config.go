package config

import (
	"log"
	"os"

	"github.com/QU35T-code/fzf-creds/utils"
)

var Workspace_path = os.ExpandEnv("$HOME/.fzf-creds")
var Local_db_name = "fzf-creds.db"
var Cme_db_path = os.ExpandEnv("$HOME/.cme/workspaces/default")
var Prefix = "%"
var Variables_custom_list = []string{"DOMAIN", "USERNAME", "PASSWORD", "TARGET"}
var Aliases_file_path = os.ExpandEnv("$HOME/.fzf-creds/aliases")

func createDefaultDirectory() {
	err := os.MkdirAll(Workspace_path, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func createDefaultAliasesFile() {
	if _, err := os.Stat(Aliases_file_path); os.IsNotExist(err) {
		if file, err := os.Create(Aliases_file_path); err == nil {
			defer file.Close()
		} else {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
}

func GetRCFilePath() string {
	shellMappings := map[string]string{
		"/bin/bash": os.ExpandEnv("$HOME/.bashrc"),
		"/bin/zsh":  os.ExpandEnv("$HOME/.zshrc"),
		"/bin/fish": os.ExpandEnv("$HOME/.fishrc"),
	}

	shell := os.Getenv("SHELL")
	return shellMappings[shell]
}

func linkToRcFile() {
	RCFile := GetRCFilePath()
	aliasTemplate := "source " + Aliases_file_path
	ret := utils.CheckExistingStringOnFile(RCFile, aliasTemplate)
	if !ret {
		utils.AppendToFile(RCFile, aliasTemplate)
	}
}

func Init() {
	createDefaultDirectory()
	createDefaultAliasesFile()
	linkToRcFile()
}
