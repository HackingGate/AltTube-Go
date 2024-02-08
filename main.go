package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Set up the database connection string
	dsn := "host=localhost dbname=AltTube user=AltTube password=AltTube port=5432 sslmode=disable"

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

// Example model
type Video struct {
	gorm.Model
	v string
}
