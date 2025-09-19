package services

import (
	"GOLANG/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTtoken(email string, userId int) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(24 * time.Minute)
	//  Create a Claims
	claims := &models.JwtClaims{
		UserId: uint(userId),
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Backend by Golang",
		},
	}
	//  Create a token with the Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
