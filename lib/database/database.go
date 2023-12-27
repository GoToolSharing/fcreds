package database

import (
	"log"

	"github.com/QU35T-code/fzf-creds/config"
	"github.com/QU35T-code/fzf-creds/lib/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	database, err := gorm.Open(sqlite.Open(config.WorkspacePath+"/"+config.LocalDbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = database

	// DB.Logger = logger.Default.LogMode(logger.Info)
	DB.Logger = logger.Default.LogMode(logger.Silent)
	DB.AutoMigrate(&models.Tools{})
}
