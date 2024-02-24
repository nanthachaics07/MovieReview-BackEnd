package authentication

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/wpcodevo/golang-fiber-jwt/initializers"
	"github.com/wpcodevo/golang-fiber-jwt/models"
)

func DeserializeRequiresAuth(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	var user models.User
	initializers.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID.String() != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// )

// func AuthenticationRequired(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")
// 	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
// 		})
// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}
// 	claims, ok := token.Claims.(*jwt.MapClaims)
// 	if !ok {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}

// 	fmt.Println(claims)
// 	return c.Next()
// }
