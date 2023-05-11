package repository

import (
	"sporesappapi/data"
	"sporesappapi/models"
)

func CreateBookingRequest(booking models.Booking) (*models.Booking, error) {
	db := data.GetDB()

	err := db.Model(&models.Booking{}).Create(&booking).Error

	return &booking, err
}

func SearchBookingHistory(userid int) ([]*models.Booking, error) {
	db := data.GetDB()
	var result []*models.Booking
	err := db.Model(&models.Booking{}).Where("user_id = ?", userid).Find(&result).Error

	return result, err
}

func PayBooking(bookingid int) error {
	db := data.GetDB()
	err := db.Model(&models.Booking{}).Where("id = ?", bookingid).Update("is_paid", true).Error

	return err
}
