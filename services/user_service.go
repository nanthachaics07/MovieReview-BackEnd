package services

import (
	"MovieReviewAPIs/model"
	"MovieReviewAPIs/repositories"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (service *UserService) RegisterUser(email, password string) error {
	return service.UserRepository.RegisterUser(email, password)
}

func (service *UserService) LoginUser(user *model.User) (string, error) {
	selectedUser, err := service.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": selectedUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service *UserService) LogoutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.JSON(map[string]string{
		"message": "User logged out successfully",
	})
}
