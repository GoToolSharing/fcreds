package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var homeDir = os.Getenv("HOME")
var WorkspacePath = os.ExpandEnv("$HOME/.local/fzf-creds")
var LocalDbName = "fzf-creds.db"
var NetexecDBPath = homeDir + "/.nxc/workspaces/default"
var Prefix = "$"
var VariablesCustomList = []string{"DOMAIN", "USERNAME", "PASSWORD", "TARGET", "HASH"}
var GlobalConfig Settings
var BaseDirectory = homeDir + "/.local/fzf-creds"

const Version = "NotAGoodVersion"

var CompletionFilePath = WorkspacePath + "/autocompletion"

func GetRCFilePath() string {
	shellMappings := map[string]string{
		"/bin/bash": os.ExpandEnv("$HOME/.bashrc"),
		"/bin/zsh":  os.ExpandEnv("$HOME/.zshrc"),
		"/bin/fish": os.ExpandEnv("$HOME/.fishrc"),
	}

	shell := os.Getenv("SHELL")
	return shellMappings[shell]
}

func ConfigureLogger() error {
	var logLevel zapcore.Level

	switch GlobalConfig.Verbose {
	case 0:
		logLevel = zap.ErrorLevel
	case 1:
		logLevel = zap.InfoLevel
	case 2:
		logLevel = zap.DebugLevel
	default:
		logLevel = zap.DebugLevel
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Ajoute des couleurs pour les niveaux de log

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	GlobalConfig.Logger, err = cfg.Build()
	if err != nil {
		return fmt.Errorf("logger configuration error: %v", err)
	}
	zap.ReplaceGlobals(GlobalConfig.Logger)
	return nil
}

// Init initializes the application by setting up necessary directories, creating a default configuration file if it doesn't exist, and loading the configuration.
func Init() error {
	if _, err := os.Stat(BaseDirectory); os.IsNotExist(err) {
		GlobalConfig.Logger.Info(fmt.Sprintf("The \"%s\" folder does not exist, creation in progress...\n", BaseDirectory))
		err := os.MkdirAll(BaseDirectory, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error folder creation: %s", err)
		}

		GlobalConfig.Logger.Info(fmt.Sprintf("\"%s\" folder created successfully\n\n", BaseDirectory))
	}
	return nil
}
