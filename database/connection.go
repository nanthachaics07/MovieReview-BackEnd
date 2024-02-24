package database

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "MovieReviewAPIs/models"
)

var DB *gorm.DB

func InitializeDB() error {
	// Connect to PostgreSQL database
	godotenv.Load(".env")
	dsn := os.Getenv("DB_prod")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	dbcon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = dbcon

	// Auto migrate models // TODO: add models here
	// err = dbcon.AutoMigrate(&models.Movies{}, &models.User{}, &models.Log{}, &models.Client{}, &models.LoginPolicy{}, &models.PasswordPolicy{}, &models.PasswordHistory{})
	// if err != nil {
	// 	log.Fatalf("Error migrating models: %v", err)
	// }

	return nil
}
