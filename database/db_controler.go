package database

import (
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"errors"
	"fmt"
	"log"
	"time"

	// "github.com/golang-jwt/jwt"
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// func GetUserFromToken(token *jwt.Token) (*models.User, error) {
// 	claims := token.Claims.(*jwt.StandardClaims)
// 	var user models.User
// 	if err := DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

func GetUserFromToken(token *jwt.Token) (*models.User, error) {
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errs.NewUnauthorizedError("Invalid token")
	}

	var userFromDB models.User
	result := DB.Db.Where("id = ?", claims.Issuer).First(&userFromDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			LogInfoErr("GetUserFromToken", "user not found in database")
			return nil, errs.NewUnauthorizedError("Invalid token")
		}
		LogInfoErr("GetUserFromToken", "error retrieving user from database: "+result.Error.Error())
		return nil, result.Error
	}

	return &userFromDB, nil
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
