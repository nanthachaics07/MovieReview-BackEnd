package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID       *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string     `gorm:"type:varchar(100);not null"`
	Email    string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	Role     *string    `gorm:"type:varchar(50);default:'user';not null"`
	// Photo     *string    `gorm:"not null;default:'default.png'"`
	Verified  *bool      `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

// LoginInput
type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=6"`
	// Photo           string `json:"photo"`
}

// LogoutInput
type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

// UserResponse
type UserResponse struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
	Role  string    `json:"role,omitempty"`
	// Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func used to filter user record
func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:    *user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  *user.Role,
		// Photo:     *user.Photo,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
