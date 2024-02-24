package handler

import (
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/services"

	"github.com/gofiber/fiber"
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
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	user, err := u.UserService.LoginUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "user": user})
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	user, err := u.UserService.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "user": user})
}

func (u *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.UserService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "data": users})
}
