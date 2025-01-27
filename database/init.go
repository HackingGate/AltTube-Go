package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

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

	// Initialize *sql.DB with retries
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			// Test the connection
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Failed to connect to database: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbInstance = db
	log.Println("Database connection established successfully.")
}
