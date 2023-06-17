package crackmapexec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/QU35T-code/fzf-creds/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID     int
	Name   string
	Domain string
}

func GetDomains() string {
	fmt.Println("Crackmapexec -> getDomains()")
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config :", err)
	}
	dbFiles, err := filepath.Glob(filepath.Join(config.Cme_db_path, "*.db"))
	if err != nil {
		log.Fatal(err)
	}
	if dbFiles == nil {
		fmt.Println("There is no data for crackmapexec")
		return ""
	}
	var domainsList []string

	for _, dbFile := range dbFiles {
		db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
		db.Logger = logger.Default.LogMode(logger.Silent)
		if err != nil {
			log.Printf("Error opening file %s : %s\n", dbFile, err)
			continue
		}

		var domains []string
		result := db.Model(&User{}).Select("DISTINCT domain").Find(&domains)
		if result.Error != nil {
			continue
		}

		if result.RowsAffected == 0 {
			continue
		}
		domainsList = append(domainsList, domains...)
	}
	fmt.Println("Domains List : ", domainsList)

	fzf_format := strings.Join(domainsList, "\n")

	cmd := exec.Command("fzf", "--prompt", "Domain > ")

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Stdin = strings.NewReader(fzf_format)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}

func GetUsernames() string {
	fmt.Println("Crackmapexec -> getUsernames()")
	return ""
}

func GetPasswords() string {
	fmt.Println("Crackmapexec -> getPasswords()")
	return ""
}

func GetTargets() string {
	fmt.Println("Crackmapexec -> getTargets()")
	return ""
}
