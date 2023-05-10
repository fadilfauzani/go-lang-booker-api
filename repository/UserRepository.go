package repository

import (
	"sporesappapi/data"
	"sporesappapi/helper"
	"sporesappapi/models"
)

func CreateUser(user models.User) (*models.User, error) {
	db := data.GetDB()

	hashedPassword := helper.HashPassword(user.Password)
	user.Password = hashedPassword

	err := db.Model(&models.User{}).Create(&user).Error

	return &user, err
}

func UpdateUser(user models.User) (*models.User, error) {
	db := data.GetDB()

	err := db.Model(&models.User{}).Where("email = ?", user.Email).Updates(&user).Error

	return &user, err
}
func ChangePasswordUser(user models.User) (*models.User, error){
	db := data.GetDB()

	hashedPassword := helper.HashPassword(user.Password)
	user.Password = hashedPassword

	err := db.Model(&models.User{}).Where("email = ?", user.Email).Updates(&user).Error
	return &user, err
}
func SeachUserByEmail(email string) (*models.User, error) {
	db := data.GetDB()
	var user models.User
	err := db.Model(&models.User{}).Where("email = ?", email).First(&user).Error
	return &user, err
}

func LoginByEmailPassword(email string, password string) (*models.User, bool) {
	user, err := SeachUserByEmail(email)
	if err != nil {
		return nil, false
	}
	if helper.CheckPasswordHash(password, user.Password) {
		return user, true
	}
	return nil, false
}
