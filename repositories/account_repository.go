package repositories

import (
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
func (u *accountUser) UserAccount(c *fiber.Ctx, uid uint) (*models.User, error) {

	var user models.User
	if err := u.db.Where("id = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	database.UseTrackingLog(c.IP(), "User", 3)

	return &user, nil
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

func (u *accountUser) UpdateUserByID(c *fiber.Ctx, id uint) (*models.User, error) {
	return nil, nil
}

func (u *accountUser) DeleteUserByID(c *fiber.Ctx, id uint) error {
	return nil
}
