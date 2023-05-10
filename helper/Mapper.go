package helper

import (
	"sporesappapi/dtos"
	"sporesappapi/models"
)

func MapRegisterToUser(src dtos.RegisterDto, dst *models.User){
	dst.Email = src.Email
	dst.Nama = src.Nama
	dst.Password = src.Password
	dst.State = src.State
	dst.City = src.City
	dst.PinCode = src.PinCode
}

func MapUserToProfile(src models.User, dst *dtos.ProfileDto){
	dst.ID = src.ID
	dst.Email = src.Email
	dst.Nama = src.Nama
	dst.State = src.State
	dst.City = src.City
	dst.PinCode = src.PinCode
	dst.Balance = src.Balance
}

func MapProfileToUser(src dtos.ProfileDto, dst *models.User){
	dst.Nama = src.Nama
	dst.State = src.State
	dst.City = src.City
	dst.PinCode = src.PinCode
}

func MapGorToDto(src models.Gor, dst *dtos.GorDto){
	dst.ID = src.ID
	dst.Nama = src.Nama
	dst.Alamat = src.Alamat
	dst.Latitude = src.Latitude
	dst.Longitude = src.Longitude
	dst.PhotoUrl = src.PhotoUrl
	dst.Rating = src.Rating
	dst.Category = src.Category
}

func MapDtotoGor(src dtos.GorDto, dst *models.Gor){
	dst.Nama = src.Nama
	dst.Alamat = src.Alamat
	dst.Latitude = src.Latitude
	dst.Longitude = src.Longitude
	dst.PhotoUrl = src.PhotoUrl
	dst.Rating = src.Rating
	dst.Category = src.Category

}