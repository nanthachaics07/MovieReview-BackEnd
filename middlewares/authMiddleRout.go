package middlewares

import (
	// "MovieReviewAPIs/database"
	// "MovieReviewAPIs/handler/errs"

	"MovieReviewAPIs/utility"
	"fmt"

	// "strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/middleware"
	"github.com/golang-jwt/jwt"
	// jwtware "github.com/gofiber/jwt/v3"
)

func MiddlewareDeserializeRout(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	// tokenString := c.GetRespHeader("Set-Cookie")
	fmt.Println("Set-Cookie from Fiber cookie: ", tokenString)
	// if tokenString != "" {
	// 	// Parse the Set-Cookie header to extract the jwt token
	// 	cookieParts := strings.Split(tokenString, ";")
	// 	for _, part := range cookieParts {
	// 		if strings.HasPrefix(part, "jwt=") {
	// 			tokenString = strings.TrimPrefix(part, "jwt=")
	// 			fmt.Println("Token from Fiber cookie: ", tokenString)
	// 			break
	// 		}
	// 	}
	// }

	if tokenString == "" {
		fmt.Println("No token found in cookies")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not logged in"})
	}

	config, err := utility.GetConfig()
	if err != nil {
		fmt.Println("Error getting config: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Config error"})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil || !tokenByte.Valid {
		fmt.Println("Error parsing token: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	c.Locals("user", claims)
	return c.Next()
}
