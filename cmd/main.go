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
	_ "MovieReviewAPIs/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Set timezone BKK
func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	} else if ict.String() != "Asia/Bangkok" {
		log.Fatal("Timezone is not UTC+7 BKK")
	}
	fmt.Println("Timezone: ", ict)

	time.Local = ict
}

func main() { // Do not delete any line in this func if U won't check each file

	//TODO: Delete this config if U Not Funny
	// Tell Me Who Handsome
	errs := godotenv.Load("startup.env")
	if errs != nil {
		log.Fatalf("Error loading .env file: %v", errs)
	}
	os.Getenv("Secret_Load")

	// Initialize timezone
	initTimeZone()

	// Initialize database connection
	err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// // Initialize Fiber app
	app := fiber.New()

	// Initialize router
	router.InitRouterHeaderConfig(app)

	// Initialize Repository
	movieRepo := repositories.NewMovieRepository(database.DB)
	userRepo := repositories.NewUserRepository(database.DB)
	accountRepo := repositories.NewAccountRepository(database.DB)

	// Initialize movie service
	movieService := services.NewMovieService(movieRepo)
	userService := services.NewUserService(userRepo)
	accountService := services.NewAccountService(accountRepo)

	// Initialize Handler
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAccountHandler(accountService)

	// Initialize routes
	router.RouterControl(app, movieHandler, userHandler, accountHandler)

	// Start Fiber server (port)
	router.RouterPortListener(app)

}
