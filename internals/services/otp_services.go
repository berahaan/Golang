package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"time"
)

func StoreOTP(User_Id uint, otpNumber string, expTime time.Time) error {
	otp := models.OTP{
		UserID:    User_Id,
		Code:      otpNumber,
		ExpiresAt: expTime,
		Used:      false,
	}

	//  now inser this information to db OTP
	return database.DB.Create(&otp).Error

}

func VerifyOTP(userId uint, code string) bool {
	var otp models.OTP

	// 1. Check if OTP exists, not expired, and unused
	err := database.DB.Where("user_id = ? AND code = ? AND used = ? AND expires_at > ?",
		userId, code, false, time.Now()).First(&otp).Error

	if err != nil {
		return false
	}

	// 2. Mark OTP as used
	otp.Used = true
	if err := database.DB.Save(&otp).Error; err != nil {
		return false
	}

	return true
}
