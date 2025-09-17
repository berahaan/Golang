package database

import (
	"GOLANG/models"
)

func MigrateTables() {
	DB.AutoMigrate(&models.Album{})
	DB.AutoMigrate(&models.User{})
}
