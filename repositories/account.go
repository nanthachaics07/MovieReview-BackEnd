package repositories

import "github.com/gofiber/fiber/v2"

type AccountRepository interface {
	UserAccount(c *fiber.Ctx) error
	UsersAccountAll(c *fiber.Ctx) error
}
