package utils

import (
	"auth/config"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyByCookie(ctx *gin.Context) (token *jwt.Token, errr string) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, "First login then come back!"
		}
		return nil, "Something went wrong!"
	}

	token, errr = verifyToken(cookie)
	return token, errr
}

func VerifyByHeaders(ctx *gin.Context) (token *jwt.Token, err string) {
	if data := ctx.Request.Header["Authorization"][0]; len(data) > 0 {
		token, err = verifyToken(data)
		return token, err
	} else {
		return nil, "First login then come back!"
	}
}

func verifyToken(data string) (*jwt.Token, string) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JwtKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, "Invalid Token"
	}

	return token, ""
}
