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
	database.UseTrackingLog(c.IP(), "UserAccount", 3)

	return &user, nil
}

func (u *accountUser) UsersAccountAll(c *fiber.Ctx) ([]models.User, error) {

	var userFromDB []models.User
	if err := u.db.Find(&userFromDB).Error; err != nil {
		return nil, err
	}

	database.UseTrackingLog(c.IP(), "UserAccountAll", 3)
	return userFromDB, nil
}

func (u *accountUser) GetuserByID(c *fiber.Ctx, id uint) (*models.User, error) {
	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	database.UseTrackingLog(c.IP(), "GetuserByID", 3)
	return &user, nil
}

func (u *accountUser) UpdateUserByID(c *fiber.Ctx, payload *models.UserUpdate, id uint) error {

	if err := u.db.Model(&models.User{}).Where("id = ?", id).Updates(payload).Error; err != nil {
		return err
	}

	database.UseTrackingLog(c.IP(), "UpdateUserByID", 3)
	return nil
}

func (u *accountUser) DeleteUserByID(c *fiber.Ctx, id uint) error {
	var user models.User

	if err := u.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	database.UseTrackingLog(c.IP(), "DeleteUserByID", 3)
	return nil
}
