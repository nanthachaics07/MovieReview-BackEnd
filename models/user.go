package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null" validate:"required"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	Password  string `json:"password" gorm:"not null" validate:"required"`
	AdminRole uint   `json:"adminRole"`
	// BusinessRole      uint      `json:"businessRole"`
	// ResidentialRole   uint      `json:"residentialRole"`
	// LastModified      time.Time `json:"lastModified"`
	// LastLoginAttempt  time.Time `json:"lastLoginAttempt"`
	// LoginAttemptCount uint      `json:"loginAttemptCount"`
	// Blocked           uint      `json:"blocked"`
	// PasswordPolicyId  uint      `json:"-" gorm:"default:1;foreignKey:id"`
	// LoginPolicyId     uint      `json:"-" gorm:"default:1;foreignKey:id"`
	// PasswordHistoryId uint      `json:"-" gorm:"default:1;foreignKey:id"`
}

// func (User) TableName() string {
// 	return "users"
// }
