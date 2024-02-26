package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler"
	"MovieReviewAPIs/repositories"
	"MovieReviewAPIs/router"
	"MovieReviewAPIs/services"
	"MovieReviewAPIs/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	} else if ict.String() != "Asia/Bangkok" {
		log.Fatal("Timezone is not UTC+7")
	}
	fmt.Println("Timezone: ", ict)

	time.Local = ict
}

func main() {

	// Tell Me Who Handsome
	//TODO: Delete this config if U Not Funny
	godotenv.Load("startup.env")
	os.Getenv("Secret_Load")

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

	// Enable CORS
	router.InitRouterConfig(app)

	// Initialize Repository
	movieRepo := repositories.NewMovieRepository(database.DB)
	userRepo := repositories.NewUserRepository(database.DB)

	// Initialize movie service
	movieService := services.NewMovieService(movieRepo)
	userService := services.NewUserService(userRepo)

	// Initialize Handler
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(userService)

	// Initialize routes
	router.Router(app, movieHandler, userHandler)

	// Start Fiber server
	err = app.Listen(config.AppPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
