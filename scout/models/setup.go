package models

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("scout.sqlite"), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}
}

func MigrateTables() {
	DB.AutoMigrate(&Index{})
	DB.AutoMigrate(&Document{})
}

