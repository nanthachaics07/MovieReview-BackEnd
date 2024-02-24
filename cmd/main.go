package main

import (
	"log"

	"MovieReviewAPIs/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Initialize database connection
	err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Define routes

	// Start Fiber server
	err = app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
