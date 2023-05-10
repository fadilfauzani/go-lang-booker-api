package data

import (
	"fmt"
	"log"
	"os"
	"sporesappapi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	err      error
)

func StartDB() {
	fmt.Printf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	if err = db.AutoMigrate(models.Booking{}, models.Gor{}, models.User{}); err != nil {
		log.Fatal("Error run database automigration", err)
	}
	fmt.Println("Database connection success.")
}

func GetDB() *gorm.DB {
	return db
}
