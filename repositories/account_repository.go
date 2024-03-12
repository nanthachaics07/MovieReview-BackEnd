package repositories

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"errors"
	"fmt"

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

	token, err := authentication.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("UserAccount", "unauthenticated")
		return err
	}

	user, _ := database.GetUserFromToken(token)

	database.UseTrackingLog(c.IP(), "User", 3)

	return c.JSON(user)
}

func (u *accountUser) UsersAccountAll(c *fiber.Ctx) error {
	token, err := authentication.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("UsersAccountAll", "unauthenticated")
		return err
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}
	fmt.Println(user.Role)

	var userFromDB []models.User
	result := u.db.Find(&userFromDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusNotFound)
			return errs.NewNotFoundError("Users not found. Error retrieving user from database")
		}
		c.Status(fiber.StatusInternalServerError)
		database.LogInfoErr("user", "error retrieving user from database: "+result.Error.Error())
		return errs.NewUnauthorizedError("Error retrieving user from database")
	}

	if *user.Role != "admin" {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("UsersAccountAll", "unauthorized")
		return errs.NewUnauthorizedError("unauthorized user role!! WHO ARE U?")
	}

	database.UseTrackingLog(c.IP(), "User", 3)

	return c.JSON(userFromDB)
}
