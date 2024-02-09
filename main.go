package main

import (
	"AltTube-Go/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

var r *gin.Engine

func main() {
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
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	err := db.AutoMigrate(&Video{})
	if err != nil {
		panic("Failed to migrate the schema")
		return
	}

	startApi()
}

func startApi() {
	r = gin.Default()

	r.GET("/ping", handlers.Ping)

	if err := r.Run(); err != nil {
		log.Fatalf("API failed to start: %v", err)
	}
}

type Video struct {
	gorm.Model
	v string
}