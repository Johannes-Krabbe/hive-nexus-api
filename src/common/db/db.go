package db

import (
	"log"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

	// todo johannes: migrate all models
    db.AutoMigrate(&models.User{})

    return db
}