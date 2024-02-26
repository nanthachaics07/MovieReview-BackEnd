package middlewares

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/utility"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func MiddlewareDeserializeRout(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("jwt") != "" {
		tokenString = c.Cookies("jwt")
	}

	if tokenString == "" {
		return errs.NewUnauthorizedError("You are not logged in")
	}

	config, err := utility.GetConfig()
	if err != nil {
		fmt.Println("Error getting config: ", err)
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		fmt.Println("Error parsing token in Header: ", err)
		return errs.NewUnauthorizedError("invalid token")
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return errs.NewUnauthorizedError("Claims are not valid")

	}

	var user models.User
	database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID.String() != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
