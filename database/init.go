package database

import (
	"AltTube-Go/ent"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var Client *ent.Client

func Init() {
	// Construct DSN from .env variables
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Initialize Ent client with retries
	var client *ent.Client
	var err error

	for i := 0; i < 5; i++ {
		client, err = ent.Open("postgres", dsn)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil || client == nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migration
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	Client = client
	log.Println("Database connection established successfully.")
}
