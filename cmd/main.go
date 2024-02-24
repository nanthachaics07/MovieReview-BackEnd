package main

import (
	"log"
	"os"

	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler"
	"MovieReviewAPIs/repositories"
	"MovieReviewAPIs/router"
	"MovieReviewAPIs/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Initialize database connection
	err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Initialize Repository
	movieRepo := repositories.NewMovieRepository(database.DB)

	// Initialize movie service
	movieService := services.NewMovieService(movieRepo)

	// Initialize Handler
	movieHandler := handler.NewMovieHandler(movieService)

	// Initialize routes
	router.Router(app, movieHandler)

	// Start Fiber server
	godotenv.Load(".env")
	port := os.Getenv("APP_port")
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
