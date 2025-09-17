package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"` // don’t expose password in JSON response
}

type UserInput struct {
	Email    string `json:"email" `
	Password string `json:"password"`
}
