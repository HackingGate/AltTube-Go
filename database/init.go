package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hackinggate/alttube-go/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	entsql "entgo.io/ent/dialect/sql" // Import the Ent SQL dialect package.
)

// Client is the global Ent client instance.
var Client *ent.Client

// DB is the global standard SQL database instance.
var DB *sql.DB

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

	// Open the standard SQL DB first.
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open standard SQL DB: %v", err)
	}

	// Attempt to ping the DB with retries.
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = DB.PingContext(ctx)
		cancel()
		if err != nil {
			log.Printf("Attempt %d: Failed to ping DB: %v. Retrying in 5 seconds...", i+1, err)
			time.Sleep(5 * time.Second)
			continue
		}
		// Successful ping.
		break
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// Create the Ent client using the standard SQL DB.
	drv := entsql.OpenDB("postgres", DB)
	Client = ent.NewClient(ent.Driver(drv))

	// Run database migrations using Atlas.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := Client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	// Optionally, query the PostgreSQL version using raw SQL.
	if err := getVersions(ctx); err != nil {
		log.Printf("Warning: failed to get PostgreSQL version: %v", err)
	}

	log.Println("Database connection established and migrations completed successfully.")
}

// getVersions queries the PostgreSQL version using raw SQL.
func getVersions(ctx context.Context) error {
	var version string
	err := DB.QueryRowContext(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		return err
	}
	log.Printf("PostgreSQL version: %s", version)
	return nil
}
