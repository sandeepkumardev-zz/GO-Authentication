package routes

import (
	controller "auth/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/")
	{
		auth.POST("signin", controller.LogIn)
		auth.POST("signup", controller.SignUp)
	}

	pvt := router.Group("/")
	{
		pvt.GET("secure", controller.Secure)
		pvt.POST("refresh", controller.Refresh)
	}

	return router
}
