package main

import (
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

	// Initialize Gin
	r := gin.Default()

	// Define a simple route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// Run the server
	err = r.Run()
	if err != nil {
		panic("Failed to start the server")
		return
	}
}

type Video struct {
	gorm.Model
	v string
}
