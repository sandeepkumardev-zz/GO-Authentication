package main

import (
	"auth/models"
	"auth/routes"
)

func main() {
	models.ConnectToDb()

	router := routes.RouterSetup()
	router.Run(":3000")
}
