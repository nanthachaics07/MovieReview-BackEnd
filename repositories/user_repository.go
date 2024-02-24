package repositories

import (
	"MovieReviewAPIs/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(user *models.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindUserByToken(token string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUserToken(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUserToken(user *models.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUserPassword(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUserPassword(user *models.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUserEmail(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUserEmail(user *models.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
