package repositories

import (
	"MovieReviewAPIs/models"

	"github.com/gofiber/fiber/v2"
)

type AccountRepository interface {
	UserAccount(c *fiber.Ctx) error
	UsersAccountAll(c *fiber.Ctx) ([]models.User, error)
	GetuserByID(c *fiber.Ctx, id uint) (*models.User, error)
}
