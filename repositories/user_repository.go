package repositories

import (
	"MovieReviewAPIs/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) RegisterUser(email, password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &model.User{
		Email:    email,
		Password: string(hashPass),
	}
	result := repo.DB.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
