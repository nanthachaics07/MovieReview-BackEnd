package middlewares

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/utility"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"MovieReviewAPIs/handler/errs"
)

type JwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get JWT token from the cookie
		cookie := c.Cookies("jwt")
		if cookie == "" {
			return errs.NewUnauthorizedError("Missing JWT token")
		}

		// Parse and validate the JWT token
		token, err := jwt.ParseWithClaims(cookie, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			config, err := utility.GetConfig()
			if err != nil {
				database.LogInfoErr("AuthenticationRequired", err.Error())
				fmt.Println("Error getting config: ", err)
			}
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errs.NewUnauthorizedError("Invalid signing method")
			}
			// Return the secret key used for signing the token
			return []byte(config.JwtSecret), nil
		})

		// Check if there's an error parsing the token
		if err != nil {
			return errs.NewUnauthorizedError("Invalid token")
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
			fmt.Println("claims: ", claims)
			// You can access claims.UserID to get the user ID, for example:
			// userID := claims.UserID
			// Proceed with the request
			return c.Next()
		} else {
			return errs.NewUnauthorizedError("Invalid token")
		}
	}
}

func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	config, err := utility.GetConfig()
	if err != nil {
		database.LogInfoErr("VerifyAuth", err.Error())
		log.Fatalf("Error getting config: %v", err)
	}

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})
}
