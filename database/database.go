package database

import (
	"log"

	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	db, err := gorm.Open(sqlite.Open(config.Workspace_path+"/"+config.Local_database_name), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Tools{})
}
