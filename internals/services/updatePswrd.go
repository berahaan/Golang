package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePassword struct {
	UserID   uint   `json:"user_id"`
	Password string `json:"password"`
}

func HandleUpdatePassword(c *gin.Context) {
	var UpdatePswrd UpdatePassword
	if err := c.ShouldBindJSON(&UpdatePswrd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request ",
		})
		return
	}
	// Validate the Password strenght

	if validate, msg := utils.ValidatePaswordStrength(UpdatePswrd.Password); !validate {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": msg,
		})
		return
	}
	hashedPasswrd, err := bcrypt.GenerateFromPassword([]byte(UpdatePswrd.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process password",
		})
		return
	}

	result := database.DB.Model(&models.User{}).Where("user_id=?", UpdatePswrd.UserID).Update("password", hashedPasswrd)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Password Updated Successfully",
	})

}
