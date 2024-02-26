package handler

import (
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/services"
	"fmt"

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
		return errs.NewBadRequestError(err.Error())
	}
	_, err := u.UserService.LoginUser(userL)
	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	// Generate token
	// token, err := u.UserService.LoginUser(userL)
	// if err != nil {
	// 	return errs.NewBadRequestError(err.Error())
	// }

	// // Set cookie  "Remember Me" check box
	// expires := time.Hour * 1 // Default expiration time
	// if rememberMe := c.FormValue("remember"); rememberMe == "true" {
	// 	expires = time.Hour * 24 // Extend expiration time for "Remember Me"
	// }
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(expires),
	// 	HTTPOnly: true,
	// })

	// Set cookie
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 12),
	// 	HTTPOnly: true,
	// })

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  "Token created successfully",
	})
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	userR := new(models.User)
	if err := c.BodyParser(userR); err != nil {
		return errs.NewBadRequestError(err.Error())
	}
	userRegister := u.UserService.RegisterUser(userR)

	return c.JSON(fiber.Map{"status": "success", "user": userRegister})
}

func (u *UserHandler) LogoutUser(c *fiber.Ctx) error {

	err := u.UserService.LogoutUser(c)
	if err != nil {
		fmt.Println("Error logging out: ", err)
		return errs.NewBadRequestError(err.Error())
	}

	// // Set cookie
	// cookie := fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    "",
	// 	Expires:  time.Now().Add(-time.Hour),
	// 	HTTPOnly: true,
	// }
	// c.Cookie(&cookie)

	fmt.Println("User logged out successfully & Delete cookie")

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User logged out successfully",
	})
}