package repository

import (
	"sporesappapi/data"
	"sporesappapi/models"
	"strings"
)

func CreateGor(gor models.Gor) (*models.Gor, error){
	db := data.GetDB()
	err := db.Model(&models.Gor{}).Create(&gor).Error

	return &gor, err
}


func SearchGor(query string, category string) ([]*models.Gor, error){
	db := data.GetDB()
	var result []*models.Gor
	err := db.Model(&models.Gor{}).Where("LOWER(nama) LIKE ? AND LOWER(category) LIKE ?", "%" + strings.ToLower(query) + "%", "%" + strings.ToLower(category) + "%").Find(&result).Error

	return result, err
}