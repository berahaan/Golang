package database

import (
	"GOLANG/internals/models"
)

func MigrateTables() {
	DB.AutoMigrate(&models.Album{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.OTP{})
	DB.AutoMigrate(&models.PasswordResetToken{})

}
