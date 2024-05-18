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

// func MiddlewareDeserializeRout(c *fiber.Ctx) error {
// 	var tokenString string

// 	// Try to get the token from the Authorization header
// 	authorization := c.Get("Authorization")
// 	if strings.HasPrefix(authorization, "Bearer ") {
// 		tokenString = strings.TrimPrefix(authorization, "Bearer ")
// 		fmt.Println("Token from Authorization header: ", tokenString)
// 	} else {
// 		// Fallback to check the jwt cookie
// 		// Try to get the token from Set-Cookie header
// 		setCookie := c.GetRespHeader("Set-Cookie")
// 		fmt.Println("Set-Cookie from Fiber cookie: ", setCookie)
// 		if setCookie != "" {
// 			// Parse the Set-Cookie header to extract the jwt token
// 			cookieParts := strings.Split(setCookie, ";")
// 			for _, part := range cookieParts {
// 				if strings.HasPrefix(part, "jwt=") {
// 					tokenString = strings.TrimPrefix(part, "jwt=")
// 					fmt.Println("Token from Fiber cookie: ", tokenString)
// 					break
// 				}
// 			}
// 		}

// 		if tokenString == "" {
// 			tokenString = c.Cookies("jwt")
// 		}
// 		fmt.Println("Token from cookie: ", tokenString)
// 	}

// 	// If no token is found, return an unauthorized error
// 	if tokenString == "" {
// 		fmt.Println("No token found")
// 		database.LogInfoErr("MiddlewareDeserializeRout", "You are not logged in")
// 		return errs.NewUnauthorizedError("You are not logged in")
// 	}

// 	// Retrieve the configuration
// 	config, err := utility.GetConfig()
// 	if err != nil {
// 		database.LogInfoErr("MiddlewareDeserializeRout", "Error getting config")
// 		fmt.Println("Error getting config: ", err)
// 		return err
// 	}

// 	// Parse and validate the token
// 	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
// 		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
// 			database.LogInfoErr("MiddlewareDeserializeRout", "Unexpected signing method")
// 			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
// 		}
// 		return []byte(config.JwtSecret), nil
// 	})

// 	if err != nil {
// 		database.LogInfoErr("MiddlewareDeserializeRout", "Error parsing token in Header")
// 		fmt.Println("Error parsing token in Header: ", err)
// 		return errs.NewUnauthorizedError("Invalid token")
// 	}

// 	// Validate the token claims
// 	_, ok := tokenByte.Claims.(jwt.MapClaims)
// 	if !ok || !tokenByte.Valid {
// 		database.LogInfoErr("MiddlewareDeserializeRout", "Claims are not valid")
// 		return errs.NewUnauthorizedError("Claims are not valid")
// 	}

// 	// ย้ายลง DB :  Implement user retrieval and validation if necessary
// 	// var user models.User
// 	// err = database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"])).Error
// 	// if err != nil {
// 	// 	database.LogInfoErr("MiddlewareDeserializeRout", "The user belonging to this token no longer exists")
// 	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "The user belonging to this token no longer exists"})
// 	// }

// 	// Attach user information to the context if needed
// 	// c.Locals("user", models.FilterUserRecord(&user))

// 	return c.Next()
// }

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
