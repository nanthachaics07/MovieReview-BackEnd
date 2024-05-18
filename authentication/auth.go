package authentication

import (
	"MovieReviewAPIs/database"
	// "MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/utility"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt"
	"github.com/dgrijalva/jwt-go"
)

var FiberAuth *fiber.Ctx

func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	config, err := utility.GetConfig()
	if err != nil {
		database.LogInfoErr("VerifyAuth", err.Error())
		return nil, err
	}

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			// // check token signing method
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// 	return nil, errs.NewUnauthorizedError("Invalid token signing method")
			// }
			// if config.JwtSecret == "" {
			// 	return nil, errs.NewUnauthorizedError("Missing JWT secret")
			// }
			return []byte(config.JwtSecret), nil
		})
}
