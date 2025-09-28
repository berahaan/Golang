package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"log"
	"strconv"
	"time"
)

func StoreOTP(User_Id uint, otpNumber string, expTime time.Time) error {
	log.Printf("Store OTP for %d and code %s", User_Id, otpNumber)
	otp := models.OTP{
		UserID:      User_Id,
		Code:        otpNumber,
		ExpiresAt:   expTime,
		Used:        false,
		Attempts:    0,
		MaxAttempts: 3,
	}

	// now insert this information to db OTP
	return database.DB.Create(&otp).Error
}

func VerifyOTP(userId uint, code string) (bool, string) {
	log.Println("VerifyOTP Func", userId, code)
	ratelimitService := NewOTPService()
	// 1. First, find the active OTP record for this user
	var otpRecord, find_user models.OTP
	find_user_Result := database.DB.Where("user_id=?", userId).First(&find_user)
	if find_user_Result.Error != nil {
		return false, "User not exists"
	}
	// now check otp attempts
	isAttempt, messg := ratelimitService.CheckOtpAttempts(userId)
	if !isAttempt {
		log.Println("!isAttempt if conditions executed ....")
		return false, messg
	}
	result := database.DB.Where("user_id = ? AND used = ?",
		userId, false).First(&otpRecord)

	if result.Error != nil {
		return false, "No active OTP found. Please request a new one."
	}

	// 3. Now check if the entered code is CORRECT
	if otpRecord.Code == code {
		// 3. Check if OTP is expired
		log.Println(" otpRecord.Code == code ", otpRecord.Code, "Code ", code)
		if time.Now().After(otpRecord.ExpiresAt) {
			return false, "OTP has expired. Please request a new one."
		}
		// CORRECT CODE - Mark as used and return success
		result := database.DB.Model(&models.OTP{}).
			Where("user_id = ? AND code = ? AND used = ?",
				userId, code, false).
			Update("used", true)

		if result.RowsAffected == 1 {
			log.Println("✅ OTP verified successfully for user", userId)
			return true, ""
		}
		return false, "Failed to mark OTP as used"
	}

	// 5. WRONG CODE - Increment attempts
	log.Println("❌ Wrong OTP attempt for user", userId, "Expected:", otpRecord.Code, "Got:", code)

	// Increment attempts for this specific OTP
	incrementErr := ratelimitService.IncrementOtpAttempts(userId, otpRecord.Code)
	if incrementErr != nil {
		log.Println("Error incrementing attempts:", incrementErr)
	}

	// Check if this wrong attempt exceeds the limit
	if otpRecord.Attempts+1 >= otpRecord.MaxAttempts {
		return false, "Too many attempts. Please request a new OTP."
	}

	remainingAttempts := otpRecord.MaxAttempts - (otpRecord.Attempts + 1)
	return false, "Invalid security code. " + strconv.Itoa(remainingAttempts) + " attempts remaining."
}
