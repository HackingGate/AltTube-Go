package database

import (
	"fmt"
	"github.com/hackinggate/alttube-go/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct DSN from .env variables
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Initialize GORM with Postgres
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.AccessToken{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&models.RefreshToken{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&models.Video{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&models.LikeVideo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	dbInstance = db
}
