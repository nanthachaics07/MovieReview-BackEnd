package database

import (
	"MovieReviewAPIs/models"
)

func LogInfoErr(funcname string, infoLog string) {
	logErr := models.Log_err{
		Funcname: funcname,
		Message:  infoLog,
	}

	DB.Create(&logErr)
}

func UseTrackingLog(email string, message string, reqType uint) {
	logLogin := models.Log_tracking_user{
		Email:      email,
		Message:    message,
		TypeReqest: reqType,
	}
	DB.Create(&logLogin)

}
