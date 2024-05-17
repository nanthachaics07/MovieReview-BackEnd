package repositories

import (
	"MovieReviewAPIs/authentication"
	"MovieReviewAPIs/database"
	"MovieReviewAPIs/models"

	"github.com/gofiber/fiber/v2"

	// "github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type accountUser struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountUser {
	return &accountUser{db: db}
}
func (u *accountUser) UserAccount(c *fiber.Ctx) error {

	token, err := authentication.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		database.LogInfoErr("UserAccount", "unauthenticated")
		return err
	}

	user, err := database.GetUserFromToken(token)
	if err != nil {
		return err
	}

	database.UseTrackingLog(c.IP(), "User", 3)

	return c.JSON(user)
}

func (u *accountUser) UsersAccountAll(c *fiber.Ctx) ([]models.User, error) {

	var userFromDB []models.User
	if err := u.db.Find(&userFromDB).Error; err != nil {
		return nil, err
	}
	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		c.Status(fiber.StatusNotFound)
	// 		return errs.NewNotFoundError("Users not found. Error retrieving user from database")
	// 	}
	// 	c.Status(fiber.StatusInternalServerError)
	// 	database.LogInfoErr("user", "error retrieving user from database: "+result.Error.Error())
	// 	return errs.NewUnauthorizedError("Error retrieving user from database")
	// }
	database.UseTrackingLog(c.IP(), "User", 3)

	return userFromDB, nil
}

func (u *accountUser) GetuserByID(c *fiber.Ctx, id uint) (*models.User, error) {
	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
