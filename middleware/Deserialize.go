package authentication

import (
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/utility"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeRequiresAuth(c *fiber.Ctx) error {
	config, err := utility.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})
	if err != nil || !token.Valid {
		fmt.Println(err)
		return errs.NewUnauthorizedError("Unauthorized Token Invalid") // 401
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		fmt.Println(err)
		return errs.NewUnauthorizedError("Unauthorized Token Claims") // 401
	}

	fmt.Println(claims)
	return c.Next()
}
