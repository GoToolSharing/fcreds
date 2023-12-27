package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
	if _, err := f.WriteString(command + "\n"); err != nil {
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

// func GetAliasTemplate(command string) string {
// 	return "alias " + command + "='fzf-creds smart " + command + "'"
// }

// func CheckIsFile(filename string) bool {
// 	if _, err := os.Stat(filename); err == nil {
// 		return true
// 	}
// 	return false
// }
