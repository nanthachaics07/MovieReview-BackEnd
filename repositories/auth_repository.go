package repositories

import (
	"MovieReviewAPIs/handler/errs"
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

// type fiberS struct {
// 	c *fiber.Ctx
// }

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) LoginUser(payload *models.SignInInput, c *fiber.Ctx) error {

	// result := r.db.Where("email =?", user.Email).First(payload)
	// if result.Error != nil {
	// 	return "", errs.NewBadRequestError(result.Error.Error())
	// }

	if err := c.BodyParser(&payload); err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	errors := models.ValidateStruct(payload)
	// if len(errors) > 0 {
	if errors != nil {
		return errs.NewBadRequestError(errors[0].Value) // set first error
	}

	var user models.User
	// result := r.db.Where("email =?", payload.Email).First(&user)
	result := r.db.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		return errs.NewBadRequestError(result.Error.Error())
	}

	// compare hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(payload.Password))
	if err != nil {
		log.Printf("Password does not match : %v", err)
		return errs.NewBadRequestError(err.Error())
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
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"iat": time.Now().Unix(),
			"nbf": time.Now().Unix(),
		})
	tokenStringVerify, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return errs.NewBadgatewayError(err.Error())
	}

	// Set cookie  "Remember Me" check box
	expires := time.Hour * 1 // Default expiration time
	if rememberMe := c.FormValue("remember"); rememberMe == "true" {
		expires = time.Hour * 24 // Extend expiration time for "Remember Me"
	}

	const setDoman = "localhost" //TODO: 	fix move to .env
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokenStringVerify,
		Path:     "/",
		Expires:  time.Now().Add(expires),
		HTTPOnly: true,
		Secure:   false, // Set to true if using HTTPS //TODO: เดี๋ยวจะมาทำแปบบบบบ
		Domain:   setDoman,
	})

	// fiberS.c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    tokenString,
	// 	Expires:  time.Now().Add(time.Hour * 72),
	// 	HTTPOnly: true,
	// })

	return nil
}

func (r *userRepository) RegisterUser(user *models.User, c *fiber.Ctx) error {

	// Hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user record
	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashPass),
	}

	result := r.db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) LogoutUser(c *fiber.Ctx) error {
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
