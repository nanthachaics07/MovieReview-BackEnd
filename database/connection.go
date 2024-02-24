package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"MovieReviewAPIs/model"
)

func InitializeDB() *gorm.DB {
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
	err = db.AutoMigrate(&model.Movie{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return db
}
