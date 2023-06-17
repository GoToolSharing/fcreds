package database

import (
	"log"

	"github.com/QU35T-code/fzf-creds/models"
	"github.com/QU35T-code/fzf-creds/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	database, err := gorm.Open(sqlite.Open(config.Workspace_path+"/"+config.Local_database_name), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = database

	// DB.Logger = logger.Default.LogMode(logger.Info)
	DB.Logger = logger.Default.LogMode(logger.Silent)
	DB.AutoMigrate(&models.Tools{})
}
