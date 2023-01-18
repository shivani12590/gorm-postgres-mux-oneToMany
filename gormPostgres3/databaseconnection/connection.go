package databaseconnection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"onetomany/model"
	"os"
)

func Connection() *gorm.DB {
	url := os.Getenv("url")
	var err error
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("database connection failed")
	}
	if db != nil {
		log.Println("database connection successful")
	}

	db.AutoMigrate(model.Employee{}, model.Department{})
	return db

}
