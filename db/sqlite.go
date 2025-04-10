package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Istance *gorm.DB = nil

func InitDB() error {
	newIstance, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connected to DB")

	newIstance.AutoMigrate(&Item{})

	Istance = newIstance

	return nil
}
