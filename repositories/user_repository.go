package repositories

import (
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/utility"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) LoginUser(user *models.User) (string, error) {
	selectedUser := new(models.User)
	result := r.db.Where("email =?", user.Email).First(selectedUser)
	if result.Error != nil {
		return "", result.Error
	}
	// compare hashed password
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password),
		[]byte(user.Password))
	if err != nil {
		log.Printf("Password does not match : %v", err)
		return "", err
	}
	// Load configuration
	config, err := utility.GetConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	// generate token
	jwtSecretKey := config.JwtSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": selectedUser.ID,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (r *userRepository) RegisterUser(user *models.User) error {
	// Hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user record
	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: []byte(hashPass),
	}

	result := r.db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) LogoutUser(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	// Return success message // TODO: Memory mapping
	return c.JSON(fiber.Map{
		"message": "User logged out successfully",
	})
}
