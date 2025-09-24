package models

import (
	"time"

	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	Code      string    `gorm:"size:6;not null" json:"code"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Used      bool      `gorm:"default:false" json:"used"`
}
type VerifyOTPInput struct {
	UserId uint   `json:"user_id" `
	Code   string `json:"code"`
}
