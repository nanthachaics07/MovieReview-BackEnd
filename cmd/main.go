package main

import (
	"log"
	"time"

	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler"
	"MovieReviewAPIs/repositories"
	"MovieReviewAPIs/router"
	"MovieReviewAPIs/services"
	"MovieReviewAPIs/utility"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Initialize timezone
	initTimeZone()

	// Initialize database connection
	err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Load configuration
	config, err := utility.GetConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
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
	err = app.Listen(config.AppPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
