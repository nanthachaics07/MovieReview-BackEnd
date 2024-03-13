package services

import "github.com/gofiber/fiber/v2"

type AccountService interface {
	UserAccount(c *fiber.Ctx) error
	UsersAccountAll(c *fiber.Ctx) error
}
