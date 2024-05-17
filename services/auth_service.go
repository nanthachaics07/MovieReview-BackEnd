package services

import (
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/repositories"

	"github.com/gofiber/fiber/v2"
)

type userService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{UserRepository: userRepository}
}

func (u *userService) LoginUser(payload *models.SignInInput, c *fiber.Ctx) error {
	loginUser := u.UserRepository.LoginUser(payload, c)
	return loginUser
}

func (u *userService) RegisterUser(payload *models.SignUpInput, c *fiber.Ctx) error {
	registerUser := u.UserRepository.RegisterUser(payload, c)
	return registerUser
}

func (u *userService) LogoutUser(c *fiber.Ctx) error {
	logoutUser := u.UserRepository.LogoutUser(c)
	return logoutUser
}
