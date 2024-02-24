package repositories

import (
	"MovieReviewAPIs/models"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	LoginUser(user *models.User) (string, error)
	RegisterUser(user *models.User) error
	LogoutUser(c *fiber.Ctx) error
}
