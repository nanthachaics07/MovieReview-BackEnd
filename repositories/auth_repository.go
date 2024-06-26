package repositories

import (
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/handler/errs"
	"MovieReviewAPIs/models"
	"MovieReviewAPIs/utility"
	"fmt"
	"log"
	"strconv"
	"strings"
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

func (r *userRepository) LoginUser(payload *models.SignInInput, c *fiber.Ctx) (string, error) {

	// result := r.db.Where("email =?", user.Email).First(payload)
	// if result.Error != nil {
	// 	return "", errs.NewBadRequestError(result.Error.Error())
	// }

	if err := c.BodyParser(&payload); err != nil {
		database.LogInfoErr("LoginUser", err.Error())
		return "", errs.NewBadRequestError(err.Error())
	}

	errors := models.ValidateStruct(payload)
	// if len(errors) > 0 {
	if errors != nil {
		database.LogInfoErr("LoginUser", errors[0].Value)
		return "", errs.NewBadRequestError(errors[0].Value) // set first error
	}

	var user models.User
	// result := r.db.Where("email =?", payload.Email).First(&user)
	result := r.db.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		database.LogInfoErr("LoginUser", result.Error.Error())
		return "", errs.NewBadRequestError(result.Error.Error())
	}

	if user.ID == 0 {
		database.LogInfoErr("LoginUser", "User not found")
		return "", errs.NewBadRequestError("User not found")
	}

	// compare hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(payload.Password))
	if err != nil {
		database.LogInfoErr("LoginUser", err.Error())
		log.Printf("Password does not match : %v", err)
		return "", errs.NewBadRequestError(err.Error())
	}
	// Load configuration
	config, err := utility.GetConfig()
	if err != nil {
		database.LogInfoErr("LoginUser", err.Error())
		log.Fatalf("Error loading configuration: %v", err)
	}
	// generate token
	jwtSecretKey := config.JwtSecret
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256,
	// 	jwt.MapClaims{
	// 		"sub": user.ID,
	// 		"exp": time.Now().Add(time.Hour * 72).Unix(),
	// 		"iat": time.Now().Unix(),
	// 		"nbf": time.Now().Unix(),
	// 	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // 1 hour
		// IssuedAt:  time.Now().Unix(),
		// NotBefore: time.Now().Unix(),
		// Subject:   user.Email,
		// Audience:  []string{"localhost"},
	})

	tokenStringVerify, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		database.LogInfoErr("LoginUser", err.Error())
		return "", errs.NewBadgatewayError(err.Error())
	}
	fmt.Println("tokenStringVerify: ", tokenStringVerify)

	// Set the token in the Authorization header
	// c.Set("Authorization", "Bearer "+tokenStringVerify)
	// fmt.Println("tokenStringVerify After c.Set: ", tokenStringVerify)

	// Set cookie  "Remember Me" check box
	expires := time.Hour * 1 // Default expiration time
	if rememberMe := c.FormValue("remember"); rememberMe == "true" {
		expires = time.Hour * 24 // Extend expiration time for "Remember Me"
	}

	// const setDoman = "localhost" //TODO: 	fix move to .env
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: tokenStringVerify,
		// Path:     "/",
		Expires:  time.Now().Add(expires),
		HTTPOnly: true,
		SameSite: "Lax", // localhost = None //FIX: if bad req delete it
		Secure:   false, // Set to true if using HTTPS
		// 	// Secure:   false, // Set to true if using HTTPS //TODO: เดี๋ยวจะมาทำแปบบบบบ
		// 	// Domain:   setDoman,
	})
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    tokenStringVerify,
	// 	Expires:  time.Now().Add(expires),
	// 	HTTPOnly: true,
	// 	Secure:   false, // Set to true if using HTTPS
	// 	Path:     "/",
	// 	SameSite: "Lax",
	// })

	database.UseTrackingLog(user.Email, "Login", 1)

	return tokenStringVerify, nil
}

func (r *userRepository) RegisterUser(payload *models.SignUpInput, c *fiber.Ctx) error {

	if err := c.BodyParser(&payload); err != nil {
		database.LogInfoErr("RegisterUser", err.Error())
		return errs.NewBadRequestError(err.Error())
	}

	if payload.Name == "" {
		database.LogInfoErr("RegisterUser", "Name cannot be empty")
		return errs.NewBadRequestError("Name cannot be empty payload.Name == ''")
	}

	errors := models.ValidateStruct(payload)
	// if len(errors) > 0 {
	if errors != nil {
		database.LogInfoErr("RegisterUser", errors[0].Value)
		return errs.NewBadRequestError(errors[0].Value) // set first error
	}

	if payload.Password != payload.PasswordConfirm {
		database.LogInfoErr("RegisterUser", "Password does not match")
		return errs.NewBadRequestError("Password does not match payload.Password != payload.PasswordConfirm")
	}

	// Hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		database.LogInfoErr("RegisterUser", err.Error())
		return err
	}

	// Create the user record
	newUser := &models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashPass),
		// Photo:    &payload.Photo,  //TODO: 	upload image [SOON]
	}
	result := r.db.Create(&newUser)

	// log check duplicate String
	if result.Error != nil && !strings.Contains(result.Error.Error(), "duplicate") {
		database.LogInfoErr("RegisterUser", result.Error.Error())
		return errs.NewConflictError(result.Error.Error())
	} else if result.Error != nil {
		database.LogInfoErr("RegisterUser", result.Error.Error())
		return errs.NewBadgatewayError(result.Error.Error())
	} else if result.RowsAffected == 0 {
		database.LogInfoErr("RegisterUser", "No rows affected")
		return errs.NewConflictError("User already exist")
	}

	database.UseTrackingLog(payload.Email, "Register", 2)

	return nil
}

func (r *userRepository) LogoutUser(c *fiber.Ctx) error {

	var expired = time.Now().Add(-time.Hour * 24)

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  expired,
		HTTPOnly: true,
		SameSite: "Lax", // localhost = None //FIX: if bad req delete it
		Secure:   false, // Set to true if using HTTPS
	})
	// // Remove the cookie
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    "",
	// 	Expires:  expired,
	// 	HTTPOnly: true,
	// 	Secure:   false,
	// 	Path:     "/",
	// 	SameSite: "Lax",
	// })

	database.UseTrackingLog(c.IP(), "Logout", 3)
	// Return a success response
	// return c.SendStatus(fiber.StatusOK)
	return nil
}
