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
	return u.UserRepository.LoginUser(payload, c)
}

func (u *userService) RegisterUser(payload *models.SignUpInput, c *fiber.Ctx) error {
	return u.UserRepository.RegisterUser(payload, c)
}

func (u *userService) LogoutUser(c *fiber.Ctx) error {
	return u.UserRepository.LogoutUser(c)
}
