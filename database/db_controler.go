package database

import (
	"MovieReviewAPIs/models"

	// "github.com/golang-jwt/jwt"
	"github.com/dgrijalva/jwt-go"
)

func GetUserFromToken(token *jwt.Token) (*models.User, error) {
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	DB.Where("id = ?", claims.Issuer).First(&user)

	return &user, nil
}
