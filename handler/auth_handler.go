package handler

import (
	// "MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
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

func (u *UserHandler) LoginUserHandler(c *fiber.Ctx) error {
	payload := new(models.SignInInput)
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Error parsing body: ", err)
		database.LogInfoErr("LoginUserHandler", err.Error())
		return errs.NewBadRequestError(err.Error())
	}
	database.LogInfoErr("LoginUserHandler Useremail", payload.Email)

	userLogin := u.UserService.LoginUser(payload, c)
	if err := userLogin; err != nil {
		fmt.Println("Error logging in: ", err)
		database.LogInfoErr("LoginUserHandler", err.Error())
		return err
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  "token is generated",
	})
}

func (u *UserHandler) RegisterUserHandler(c *fiber.Ctx) error {
	payload := new(models.SignUpInput)
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Error parsing body: ", err)
		database.LogInfoErr("RegisterUserHandler", err.Error())
		return errs.NewBadRequestError(err.Error())
	}
	userRegister := u.UserService.RegisterUser(payload, c)

	return c.JSON(fiber.Map{"status": "success", "user": userRegister})
}

func (u *UserHandler) LogoutUserHandler(c *fiber.Ctx) error {
	// _, errorl := authentication.VerifyAuth(c)
	// if errorl != nil {
	// 	database.LogInfoErr("LogoutUserHandler", errorl.Error())
	// 	return errs.NewUnexpectedError(errorl.Error())
	// }

	// user, err := database.GetUserFromToken(token)

	err := u.UserService.LogoutUser(c)
	if err != nil {
		fmt.Println("Error logging out: ", err)
		database.LogInfoErr("LogoutUserHandler", err.Error())
		return errs.NewBadRequestError(err.Error())
	}

	fmt.Println("User logged out successfully & Delete cookie")
	database.UseTrackingLog(c.IP(), "Logout", 3)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User logged out successfully",
	})
}
