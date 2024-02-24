package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "MovieReviewAPIs/models"
	"MovieReviewAPIs/utility"
)

var DB *gorm.DB

func InitializeDB() error {
	// Connect to PostgreSQL database
	config, err := utility.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	dsn := config.DBPass

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
