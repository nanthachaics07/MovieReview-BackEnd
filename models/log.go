package models

import (
	"gorm.io/gorm"
)

type Log_err struct {
	gorm.Model
	Funcname string `json:"userName"`
	Message  string `json:"message"`
}

type Log_tracking_user struct {
	gorm.Model
	Email      string `gorm:"varchar(50);not null"`
	Message    string `gorm:"varchar(100);not null"`
	TypeReqest uint   `gorm:"not null;default:0"`
}
