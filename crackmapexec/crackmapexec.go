package crackmapexec

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/QU35T-code/fzf-creds/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CrackmapexecInterface interface {
}

type Users struct {
}

type Computers struct {
}

var computedData ComputedData

type ComputedData struct {
	Domain   string
	Username string
	Password string
	Target   string
}

func getDataFromDatabases(search string, cmeInterface CrackmapexecInterface) []string {
	dbFiles, err := filepath.Glob(filepath.Join(config.Cme_db_path, "*.db"))
	if err != nil {
		log.Fatal(err)
	}
	if dbFiles == nil {
		return nil
	}
	var dataList []string

	for _, dbFile := range dbFiles {
		db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
		db.Logger = logger.Default.LogMode(logger.Silent)
		if err != nil {
			log.Printf("Error opening file %s : %s\n", dbFile, err)
			continue
		}

		var data []string
		result := db.Model(cmeInterface).Select("DISTINCT " + search).Find(&data)
		if result.Error != nil {
			continue
		}

		if result.RowsAffected == 0 {
			continue
		}
		dataList = append(dataList, data...)
	}

	return dataList
}

func getHeader() string {
	computedData.Domain = strings.ReplaceAll(computedData.Domain, "\n", "")
	computedData.Username = strings.ReplaceAll(computedData.Username, "\n", "")
	computedData.Password = strings.ReplaceAll(computedData.Password, "\n", "")
	computedData.Target = strings.ReplaceAll(computedData.Target, "\n", "")

	template := "Domain : '" + computedData.Domain + "' - Username : '" + computedData.Username + "' - Password : '" + computedData.Password + "' - Target : '" + computedData.Target + "'"
	return template
}

func askToFZF(dataList []string, message string) string {
	fzf_format := strings.Join(dataList, "\n")
	if len(dataList) == 1 {
		return dataList[0]
	}
	cmd := exec.Command("fzf", "--prompt", message+" > ", "--header", getHeader())

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Stdin = strings.NewReader(fzf_format)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}

func GetData() ComputedData {
	return computedData
}

func GetDomains() {
	users := &Users{}
	domainsList := getDataFromDatabases("domain", users)
	if domainsList == nil {
		return
	}
	domain_selection := askToFZF(domainsList, "Domain")
	computedData.Domain = domain_selection
}

func GetUsernames() {
	usernames := &Users{}
	usernamesList := getDataFromDatabases("username", usernames)
	if usernamesList == nil {
		return
	}
	username_selection := askToFZF(usernamesList, "Usernames")
	computedData.Username = username_selection
}

func GetPasswords() {
	passwords := &Users{}
	passwordsList := getDataFromDatabases("password", passwords)
	if passwordsList == nil {
		return
	}
	password_selection := askToFZF(passwordsList, "Passwords")
	computedData.Password = password_selection
}

func GetTargets() {
	targets := &Computers{}
	targetsList := getDataFromDatabases("ip", targets)
	if targetsList == nil {
		return
	}
	target_selection := askToFZF(targetsList, "Targets")
	computedData.Target = target_selection
}
