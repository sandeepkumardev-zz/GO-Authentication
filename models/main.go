package models

import (
	"auth/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func ConnectToDb() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Printf("Status: %v\n", err)
	}

	config.DB.AutoMigrate(&User{})

	fmt.Println("Database connected!")
}
