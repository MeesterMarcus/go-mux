package config

import (
	"fmt"
	"log"
	"os"

	"github.com/MeesterMarcus/go-mux/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func ConnectToDB() (*gorm.DB, error) {
	DB_USER := os.Getenv("FOOTBALL_DB_USERNAME")
	DB_PASSWORD := os.Getenv("FOOTBALL_DB_PASSWORD")
	DB_NAME := os.Getenv("FOOTBALL_DB_NAME")
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=5432 sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&models.Booking{})
	return db, nil
}
