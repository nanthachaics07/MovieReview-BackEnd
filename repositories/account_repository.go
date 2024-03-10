package repositories

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/utility"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	// "github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type accountUser struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountUser {
	return &accountUser{db: db}
}
func (u *accountUser) UserAccount(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	config, err := utility.GetConfig()
	if err != nil {
		database.LogInfoErr("User", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("user", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("user", "invalid token")
		return c.JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var userFromDB models.User
	result := u.db.Where("id = ?", claims.Issuer).First(&userFromDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "User not found",
			})
		}
		c.Status(fiber.StatusInternalServerError)
		database.LogInfoErr("user", "error retrieving user from database: "+result.Error.Error())
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	database.UseTrackingLog(c.IP(), "User", 3)

	return c.JSON(userFromDB)
}

// func (u *accountUser) UserAccount(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")

// 	config, err := utility.GetConfig()
// 	if err != nil {
// 		database.LogInfoErr("User", err.Error())
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Internal Server Error",
// 		})
// 	}

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(config.JwtSecret), nil
// 		})

// 	if err != nil {
// 		c.Status(fiber.StatusUnauthorized)
// 		database.LogInfoErr("user", "unauthenticated")
// 		return c.JSON(fiber.Map{
// 			"message": "unauthenticated",
// 		})
// 	}

// 	if !token.Valid {
// 		c.Status(fiber.StatusUnauthorized)
// 		database.LogInfoErr("user", "invalid token")
// 		return c.JSON(fiber.Map{
// 			"message": "invalid token",
// 		})
// 	}

// 	user, err := database.GetUserFromToken(token)
// 	if err != nil {
// 		c.Status(fiber.StatusInternalServerError)
// 		database.LogInfoErr("user", "error getting user from token: "+err.Error())
// 		return c.JSON(fiber.Map{
// 			"message": "Internal Server Error",
// 		})
// 	}

// 	database.UseTrackingLog(c.IP(), "User", 3)

// 	return c.JSON(user)
// }
