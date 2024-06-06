package repositories

import (
	"MovieReviewAPIs/models"

	"github.com/gofiber/fiber/v2"
)

type AccountRepository interface {
	UserAccount(c *fiber.Ctx, uid uint) (*models.User, error)
	UsersAccountAll(c *fiber.Ctx) ([]models.User, error)
	GetuserByID(c *fiber.Ctx, id uint) (*models.User, error)
	UpdateUserByID(c *fiber.Ctx, payload *models.UserUpdate, id uint) error
	DeleteUserByID(c *fiber.Ctx, id uint) error
}
