package middlewares

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/utility"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"MovieReviewAPIs/handler/errs"
)

type JwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleware() fiber.Handler { //TODO: fix this cookie can't set in header on react 18 beta:
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
			fmt.Println("claims.UserID: ", claims.UserID)
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

func UserTokenMiddleware() fiber.Handler { //TODO: Update ROLE 'admin'
	return func(c *fiber.Ctx) error {
		// ดึงค่า token จาก header
		tokenAuth := c.Get("Authorization")
		println("tokenAUTH: ", tokenAuth)

		if !strings.HasPrefix(tokenAuth, "Bearer ") {
			database.LogInfoErr("MiddlewareDeserializeRout", "Missing or invalid token prefix")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid token prefix",
			})
		}

		token := strings.TrimPrefix(tokenAuth, "Bearer ")
		fmt.Println("token in header: ", token)

		if token == "" {
			database.LogInfoErr("MiddlewareDeserializeRout", "Missing token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token",
			})
		}

		config, err := utility.GetConfig()
		if err != nil {
			database.LogInfoErr("MiddlewareDeserializeRout", "Error getting config")
			fmt.Println("Error getting config: ", err)
			return err
		}

		// Parse and validate the token
		tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				database.LogInfoErr("MiddlewareDeserializeRout", "Unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}
			return []byte(config.JwtSecret), nil
		})

		if err != nil {
			database.LogInfoErr("MiddlewareDeserializeRout", "Error parsing token in Header")
			fmt.Println("Error parsing token in Header: ", err)
			return errs.NewUnauthorizedError("Invalid token")
		}

		// Validate the token claims
		_, ok := tokenByte.Claims.(jwt.MapClaims)
		if !ok || !tokenByte.Valid {
			database.LogInfoErr("MiddlewareDeserializeRout", "Claims are not valid")
			return errs.NewUnauthorizedError("Claims are not valid")
		}

		return c.Next()
	}
}

func CookieTokenMiddleware() fiber.Handler { //TODO: Update For Postman test
	return func(c *fiber.Ctx) error {
		var token string

		// Extract token from the Authorization headerrrrrr
		authorization := c.Get("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			token = strings.TrimPrefix(authorization, "Bearer ")
			fmt.Println("Token from Authorization header: ", token)
		} else {
			// Fallback to check the Set-Cookie header
			setCookie := c.Get("Cookie")
			fmt.Println("Cookies: ", setCookie)
			if setCookie != "" {
				// Parse the Set-Cookie header to extract the jwt token
				cookieParts := strings.Split(setCookie, ";")
				for _, part := range cookieParts {
					part = strings.TrimSpace(part)
					if strings.HasPrefix(part, "jwt=") {
						token = strings.TrimPrefix(part, "jwt=")
						fmt.Println("Token from Cookie: ", token)
						break
					}
				}
			}
		}

		if token == "" {
			database.LogInfoErr("UserTokenMiddleware", "Missing token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token",
			})
		}

		// Get configuration which includes the JWT secret
		config, err := utility.GetConfig()
		if err != nil {
			database.LogInfoErr("UserTokenMiddleware", "Error getting config")
			fmt.Println("Error getting config: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Internal Server Error",
				"message": "Error getting configuration",
			})
		}

		// Parse and validate the token
		tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				database.LogInfoErr("UserTokenMiddleware", "Unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}
			return []byte(config.JwtSecret), nil
		})

		if err != nil {
			database.LogInfoErr("UserTokenMiddleware", "Error parsing token")
			fmt.Println("Error parsing token: ", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
		}

		// Validate the token claims
		if claims, ok := tokenByte.Claims.(jwt.MapClaims); ok && tokenByte.Valid {
			// if ok {
			// 	claims := tokenByte.Claims.(jwt.MapClaims)
			// 	c.Locals("userID", claims["userID"])
			// 	c.Locals("userRole", claims["userRole"])
			// 	c.Locals("userEmail", claims["userEmail"])
			// }
			// Optional: Add claims to context for further use
			c.Locals("user", claims)
			return c.Next()
		} else {
			database.LogInfoErr("UserTokenMiddleware", "Invalid claims or token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "Invalid claims or token",
			})
		}
	}
}
