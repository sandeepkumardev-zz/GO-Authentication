package controller

import (
	"auth/models"
	service "auth/services"
	"auth/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LogIn(ctx *gin.Context) {
	var user models.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, msg := service.FindUser(&user)
	if msg != "" {
		ctx.JSON(http.StatusUnauthorized, msg)
		return
	}

	// genrate token
	expireTokenTime := time.Now().Add(time.Minute * 10)

	tokenSting, err := utils.CreateToken(user.Username, expireTokenTime)
	if err != "" {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	//set Cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenSting,
		Expires: expireTokenTime,
	})

	type Response struct {
		Message string
		Token   string
	}

	ctx.JSON(http.StatusOK, &Response{Message: res, Token: tokenSting})
}

func SignUp(ctx *gin.Context) {
	var user models.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, err := service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Something went wrong!")
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func Secure(ctx *gin.Context) {
	//verify by cookie
	// token, err := utils.VerifyByCookie(ctx)

	//verify by headers
	token, err := utils.VerifyByHeaders(ctx)

	if err != "" {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func Refresh(ctx *gin.Context) {
	token, err := utils.VerifyByHeaders(ctx)

	if err != "" {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	// if time.Unix(token.Claims[expiresAt], 0).Sub(time.Now()) > 30*time.Second {

	// }

	// expireTokenTime := time.Now().Add(time.Minute * 10)

	// atClaims := jwt.MapClaims{}
	// atClaims["expiresAt"] = expireTokenTime.Unix()
	// atClaims["username"] = username
	// atClaims["authorized"] = true

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// tokenSting, err := token.SignedString(config.JwtKey)
	// if err != nil {
	// 	return "", "Something went wrong!"

	// }

	ctx.JSON(http.StatusOK, token.Claims)
}
