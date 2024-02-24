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

func (u *userService) LoginUser(user *models.User) (string, error) {
	return u.UserRepository.LoginUser(user)
}

func (u *userService) RegisterUser(user *models.User) error {
	return u.UserRepository.RegisterUser(user)
}

func (u *userService) LogoutUser(c *fiber.Ctx) error {
	return u.UserRepository.LogoutUser(c)
}
