package models

import (
	"time"

	"gorm.io/gorm"
)

type PasswordResetToken struct {
	gorm.Model
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"size:64;not null;uniqueIndex" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Used      bool      `gorm:"default:false" json:"used"`
}

// TableName specifies the table name
func (PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}
