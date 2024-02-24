package handler

import (
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
	userL := new(models.User)
	if err := c.BodyParser(userL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	userLogin, err := u.UserService.LoginUser(userL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Generate token
	token, err := u.UserService.LoginUser(userL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"status": "success", "user": userLogin})
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	userR := new(models.User)
	if err := c.BodyParser(userR); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	userRegister := u.UserService.RegisterUser(userR)

	return c.JSON(fiber.Map{"status": "success", "user": userRegister})
}

func (u *UserHandler) LogoutUser(c *fiber.Ctx) error {

	err := u.UserService.LogoutUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Set cookie
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    "",
	// 	Expires:  time.Now().Add(-time.Hour),
	// 	HTTPOnly: true,
	// })

	return c.JSON(fiber.Map{"status": "success", "message": "User logged out successfully"})
}
