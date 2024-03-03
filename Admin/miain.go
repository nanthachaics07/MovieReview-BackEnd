package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"MovieReviewAPIs/Admin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Define Movie struct here

func main() {
	DBconfig := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		DBconfig.Database.Host, DBconfig.Database.Port, DBconfig.Database.Username,
		DBconfig.Database.Password, DBconfig.Database.Name)

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

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// AutoMigrate the Movie model
	err = db.AutoMigrate(&Movie{})
	if err != nil {
		log.Fatalf("Error migrating model: %v", err)
	}

	// Open JSON file
	filePath := "/Users/nanthachai/CS07_PROJECT/E-Commerce-Movie_Store/newBackEndGo/MovieReviewAPIs/Admin/etc/json/movies.json"
	// filePath := "/Users/nanthachai/CS07_PROJECT/E-Commerce-Movie_Store/newBackEndGo/MovieReviewAPIs/Admin/etc/json/movie.json"

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}
	fmt.Println("Import JSON file successful")

	var movies []Movie
	if err := json.Unmarshal(jsonData, &movies); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Insert data into database
	if err := insertMovies(db, movies); err != nil {
		log.Fatalf("Error inserting movies: %v", err)
	}

	fmt.Println("Movies inserted successfully!")
}

func insertMovies(db *gorm.DB, movies []Movie) error {
	for _, movie := range movies {
		if err := db.Create(&movie).Error; err != nil {
			return err
		}
	}
	return nil
}

func forceDeleteAll(db *gorm.DB, model interface{}) error {
	if err := db.Where("1 = 1").Delete(model).Error; err != nil {
		return err
	}
	return nil
}
