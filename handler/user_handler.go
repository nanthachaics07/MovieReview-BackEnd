package handler

import (
	"MovieReviewAPIs/model"
	"MovieReviewAPIs/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DefineUserRoutes(app *fiber.App, userService *services.UserService) {
	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		err := userService.RegisterUser(user.Email, user.Password)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(user)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		token, err := userService.LoginUser(user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})
		return c.JSON(map[string]string{
			"message": "User logged in successfully",
			"token":   "Your token is: " + token,
		})
	})

	app.Post("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.JSON(map[string]string{
			"message": "User logged out successfully",
		})
	})
}
