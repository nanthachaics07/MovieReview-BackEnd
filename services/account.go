package services

import (
	"MovieReviewAPIs/models"

	"github.com/gofiber/fiber/v2"
)

type AccountService interface {
	UserAccount(c *fiber.Ctx, user *models.User) (*models.User, error)
	UsersAccountAll(c *fiber.Ctx, user *models.User) ([]models.User, error)
	GetUserByID(c *fiber.Ctx, user *models.User, id uint) (*models.User, error)
}
