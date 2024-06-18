package database

import (
	"MovieReviewAPIs/utility"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var DB DBinstance

func InitializeDB() error {
	// Connect to PostgreSQL database
	config, err := utility.GetConfig()
	if err != nil {
		LogInfoErr("InitializeDB", err.Error())
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
		LogInfoErr("InitializeDB", err.Error())
		log.Fatalf("Error connecting to database: %v", err)
	}

	reSQLdb, err := dbcon.DB()
	if err != nil {
		LogInfoErr("InitializeDB", err.Error())
		fmt.Println("Connected to database Because: ", err)
		reSQLdb.Close()
		return fmt.Errorf("error getting sql.DB: %#v", err)
	}

	DB.Db = dbcon

	// Create UUID extension in PG
	// dbcon.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	// Auto migrate models // TODO: add models here

	// err = dbcon.AutoMigrate(
	// 	&models.User{},
	// 	&models.Log_err{},
	// 	&models.Log_tracking_user{},
	// )
	// if err != nil {
	// 	LogInfoErr("InitializeDB", err.Error())
	// 	log.Fatalf("Error migrating models: %v", err)
	// }

	settingDBConnection()

	return nil
}

func settingDBConnection() {
	// Get the underlying *sql.DB instance
	sqlDB, err := DB.Db.DB()
	if err != nil {
		LogInfoErr("settingDB", err.Error())
		log.Fatalf("Error getting underlying *sql.DB: %v", err)
	}

	// Set max open connections
	sqlDB.SetMaxOpenConns(20)

	// Set max idle connections
	sqlDB.SetMaxIdleConns(20)

	// Set max lifetime
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	// Set max idle time
	// sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// Ping the database
	err = sqlDB.Ping()
	if err != nil {
		LogInfoErr("settingDB", err.Error())
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to database successfully")
}
