package repositories

import (
	"MovieReviewAPIs/models"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	LoginUser(payload *models.SignInInput, c *fiber.Ctx) error
	RegisterUser(payload *models.SignUpInput, c *fiber.Ctx) error
	LogoutUser(c *fiber.Ctx) error
}
