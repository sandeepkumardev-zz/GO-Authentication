package service

import (
	"auth/config"
	"auth/models"

	"golang.org/x/crypto/bcrypt"
)

var err error

func FindUser(user *models.User) (res string, msg string) {
	var pswd = user.Password

	if err = config.DB.Where("username = ?", user.Username).First(user).Error; err != nil {
		return "", "Invalid Username"
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pswd)); err != nil {
		return "", "Invalid Password"
	}

	return "Successfully signed in as " + user.Username, ""
}

func CreateUser(user *models.User) (res string, err error) {
	if err = config.DB.Where("username = ?", user.Username).First(user).Error; err != nil {
		// hash password
		hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if hashErr != nil {
			return "", hashErr
		}

		user.Password = string(hash)

		if err = config.DB.Create(user).Error; err != nil {
			return "", err
		}
		return "Successfully created user", nil
	}

	return "User already exists", nil
}
