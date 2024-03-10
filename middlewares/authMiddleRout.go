package middlewares

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/utility"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/middleware"
	"github.com/golang-jwt/jwt"

	jwtware "github.com/gofiber/jwt/v3"
)

func TokenValidator(c *fiber.Ctx) error {
	// Extract token from response headers
	tokenCookie := c.Cookies("jwt")
	fmt.Println("tokenCookie: ", tokenCookie)
	if tokenCookie == "" {
		return errs.NewUnauthorizedError("Missing token")
	}

	// Parse the token
	token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
		configJ, err := utility.GetConfig()
		fmt.Println("configJ: ", configJ)

		if err != nil {
			database.LogInfoErr("AuthenticationRequired", err.Error())
			fmt.Println("Error getting config: ", err)
		}
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthorizedError("Invalid token")
		}
		// Provide your JWT secret here
		return []byte(configJ.JwtSecret), nil
	})
	if err != nil {
		return errs.NewUnauthorizedError("Invalid token")
	}

	// Check if the token is valid
	if !token.Valid {
		return errs.NewUnauthorizedError("Invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errs.NewUnauthorizedError("Invalid token claims")
	}

	// You can perform additional checks on claims if needed

	// Set user ID from claims to context for later use
	c.Locals("userID", claims["sub"])

	// Token is valid, proceed to next handler
	return c.Next()
}

func AuthRequ() func(*fiber.Ctx) error {
	configJ, err := utility.GetConfig()
	fmt.Println("configJ: ", configJ)

	if err != nil {
		database.LogInfoErr("AuthenticationRequired", err.Error())
		fmt.Println("Error getting config: ", err)
	}
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(configJ.JwtSecret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	})
}

func AuthenticationRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	fmt.Println("cookie: ", cookie)

	configJ, err := utility.GetConfig()
	fmt.Println("configJ: ", configJ)

	if err != nil {
		database.LogInfoErr("AuthenticationRequired", err.Error())
		fmt.Println("Error getting config: ", err)
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configJ.JwtSecret), nil
		})

	fmt.Println("token: ", token)

	if err != nil || !token.Valid {
		fmt.Println("Error validating token: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized token Validation",
		})
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		fmt.Println("Error parsing claims: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized token claims",
		})
	}
	fmt.Println(claims)
	return c.Next()
}

func MiddlewareDeserializeRout(c *fiber.Ctx) error {
	var tokenString string
	// authorization := c.GetRespHeader("Set-Cookie")
	authorization := c.Get("Authorization")
	fmt.Println("authorization c.Get: ", authorization)

	if strings.HasPrefix(authorization, "jwt=") {
		tokenString = strings.TrimPrefix(authorization, "JWT=")
		fmt.Println("tokenString (-Bearer): ", tokenString)
	} else if c.Cookies("jwt") != "" {
		tokenString = c.Cookies("jwt")
		fmt.Println("tokenString: ", tokenString)
	}

	if tokenString == "" {
		fmt.Println("tokenString: ", tokenString)
		database.LogInfoErr("MiddlewareDeserializeRout", "You are not logged in")
		return errs.NewUnauthorizedError("#.Empty token in UR Browser.# You are not logged in")
	}

	config, err := utility.GetConfig()
	if err != nil {
		database.LogInfoErr("MiddlewareDeserializeRout", "Error getting config")
		fmt.Println("Error getting config: ", err)
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if JWM, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {

			fmt.Println("JWM: ", JWM)
			database.LogInfoErr("MiddlewareDeserializeRout", "unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})
	fmt.Printf("tokenByte: %v", tokenByte)

	if err != nil {
		database.LogInfoErr("MiddlewareDeserializeRout", "Error parsing token in Header")
		fmt.Println("Error parsing token in Header: ", err)
		return errs.NewUnauthorizedError("invalid token")
	}

	_, ok := tokenByte.Claims.(jwt.MapClaims)
	// claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		database.LogInfoErr("MiddlewareDeserializeRout", "Claims are not valid")
		return errs.NewUnauthorizedError("Claims are not valid")

	}

	// var user models.User
	// database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	// if user.ID.String() != claims["sub"] {
	// 	database.LogInfoErr("MiddlewareDeserializeRout", "the user belonging to this token no logger exists")
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	// }

	// c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
