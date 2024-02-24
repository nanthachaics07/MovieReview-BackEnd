package main

import (
	"log"
	"os"
	"time"

	"MovieReviewAPIs/handler"
	"MovieReviewAPIs/model"
	"MovieReviewAPIs/repositories"
	"MovieReviewAPIs/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	db := initializeDB()

	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	handler.DefineRoutes(app, services.NewMovieService(repositories.NewMovieRepository(db)),
		services.NewUserService(repositories.NewUserRepository(db)))

	// Start Fiber server
	err = app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeDB() *gorm.DB {
	// Connect to PostgreSQL database
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&model.Movie{}, &model.User{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return db
}
