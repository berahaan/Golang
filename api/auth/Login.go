package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	emails "GOLANG/pkg/email"
	"GOLANG/pkg/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLoginAuth(c *gin.Context) {
	// now we need to take emails and password of a users from clients
	fmt.Println("Login Endpoint Hit")
	var LoginInput models.LoginInput

	if err := c.ShouldBindJSON(&LoginInput); err != nil {
		c.JSON(400, gin.H{
			"error": "unable to bind json" + err.Error(),
		})
		return
	}
	//  sanitize the email and password
	LoginInput.Email = utils.SanitizeEmail(LoginInput.Email)
	LoginInput.Password = utils.SanitizePassword(LoginInput.Password)

	if !utils.ValidateEmail(LoginInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Emails Please provide a valid emails",
		})
		return
	}
	// find user by emails and get all information to the user variables
	var user models.User
	result := database.DB.Where("email=?", LoginInput.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid email or password ",
		})
		return
	}
	// compare password with stored hash password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid email or password",
		})
		return
	}
	// Create a functions to generate and send OTP of 6 digit
	num := utils.GenerateOTP(6)
	otp_expires := time.Now().Add(5 * time.Minute)
	// Store OTP in database along with User ID
	err := services.StoreOTP(user.ID, num, otp_expires)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to store OTP",
		})
		return
	}
	subject := "Your Login OTP Code"
	body := "Hello, your OTP code is: " + num + "\nIt will expire in 5 minutes."
	// Send OTP Emails
	if err := emails.SendEmailsTo(user.Email, subject, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send OTP email" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "OTP sent to your emails please check your emails",
	})
	// Verify Tokens
	// generate JWT with claims

}
