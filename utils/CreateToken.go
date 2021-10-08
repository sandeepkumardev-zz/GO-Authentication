package utils

import (
	"auth/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string, expireTokenTime time.Time) (string, string) {
	atClaims := jwt.MapClaims{}
	atClaims["expiresAt"] = expireTokenTime.Unix()
	atClaims["username"] = username
	atClaims["authorized"] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenSting, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", "Something went wrong!"

	}
	return tokenSting, ""
}
