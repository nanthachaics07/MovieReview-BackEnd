package authentication

import (
	"MovieReviewAPIs/database"
	"fmt"
	"strings"

	// "MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/utility"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt"
	"github.com/dgrijalva/jwt-go"
)

var FiberAuth *fiber.Ctx

// func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
// 	cookie := c.Cookies("jwt")
// 	fmt.Println("VerifyAuth : ", cookie)

// 	config, err := utility.GetConfig()
// 	if err != nil {
// 		database.LogInfoErr("VerifyAuth", err.Error())
// 		return nil, err
// 	}

// 	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
// 		func(token *jwt.Token) (interface{}, error) {

// 			// // check token signing method
// 			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			// 	return nil, errs.NewUnauthorizedError("Invalid token signing method")
// 			// }
// 			// if config.JwtSecret == "" {
// 			// 	return nil, errs.NewUnauthorizedError("Missing JWT secret")
// 			// }
// 			return []byte(config.JwtSecret), nil
// 		})
// }

func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
	authorization := c.Get("Authorization")
	var tokenString string

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else {
		cookie := c.Cookies("jwt")
		fmt.Println("VerifyAuth : ", cookie)

		if cookie == "" {
			return nil, fmt.Errorf("missing token")
		}

		tokenString = cookie
	}

	config, err := utility.GetConfig()
	if err != nil {
		database.LogInfoErr("VerifyAuth", err.Error())
		return nil, err
	}

	return jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Check token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token signing method")
			}
			if config.JwtSecret == "" {
				return nil, fmt.Errorf("missing JWT secret")
			}
			return []byte(config.JwtSecret), nil
		})
}
