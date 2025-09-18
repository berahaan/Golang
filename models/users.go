package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"` // don’t expose password in JSON response
}

type SignUpInput struct {
	Email    string `json:"email" `
	Password string `json:"password"`
}
type LoginInput struct {
	Email    string `json:"email" binding:"required" `
	Password string `json:"password" binding:"required"`
}
type JwtClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}
