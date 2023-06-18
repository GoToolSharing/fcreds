package utils

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Workspace_path        string   `mapstructure:"WORKSPACE_PATH"`
	Cme_db_path           string   `mapstructure:"CME_DB_PATH"`
	Variables_prefix      string   `mapstructure:"VARIABLES_PREFIX"`
	Variables_custom_list []string `mapstructure:"VARIABLES_CUSTOM_LIST"`
	Aliases_file_path     string   `mapstructure:"ALIASES_FILE_PATH"`
	Local_database_name   string   `mapstructure:"LOCAL_DATABASE_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env.example")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func RemoveLineFromFile(filename string, lineToRemove string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	tempFilename := filename + ".tmp"
	tempFile, err := os.Create(tempFilename)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == lineToRemove {
			continue
		}
		_, err := tempFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	file.Close()
	tempFile.Close()

	err = os.Rename(tempFilename, filename)
	if err != nil {
		return err
	}

	return nil
}

func AppendToFile(filename string, command string) {
	f, err := os.OpenFile(filename,
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

func CheckExistingStringOnFile(filePath string, comparedString string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, comparedString) {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}
