package database

import (
	"AltTube-Go/ent"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"entgo.io/ent/migrate"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Client is the global Ent client instance.
var Client *ent.Client

// loadEnv loads the .env file.
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}

// Init initializes the database connection and runs migrations.
func Init() {
	loadEnv()

	// Construct DSN from .env variables.
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var client *ent.Client
	var err error

	// Attempt to connect with retries.
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		client, err = ent.Open("postgres", dsn)
		if err != nil {
			log.Printf("Attempt %d: Failed to open connection: %v. Retrying in 5 seconds...", i+1, err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Ping to ensure the connection is ready.
		if err = client.DB().Ping(); err != nil {
			log.Printf("Attempt %d: Connection ping failed: %v. Retrying in 5 seconds...", i+1, err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Connection is established.
		break
	}

	if err != nil || client == nil {
		log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// Run database migrations using Atlas.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	Client = client
	log.Println("Database connection established and migrations completed successfully.")
}
