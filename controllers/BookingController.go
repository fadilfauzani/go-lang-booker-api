package controllers

import (
	"fmt"
	"net/http"
	"sporesappapi/dtos"
	"sporesappapi/helper"
	"sporesappapi/models"
	"sporesappapi/repository"
	"sporesappapi/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBookingRequest(c *gin.Context) {
	gor, err := strconv.Atoi(c.Param("gor"))
	fmt.Println(gor)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid Gor Parameter"})
		return
	}

	token, _ := helper.ExtractBearerToken(c.GetHeader("Authorization"))
	claim, _ := services.VerifyToken(token)

	var bookingDto dtos.BookingDto
	err = c.ShouldBind(&bookingDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !helper.ValidateBookingTime(bookingDto.BookedTimeStart, bookingDto.BookedTimeEnd) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid Time Parameter"})
		return
	}

	var booking = models.Booking{
		UserId:   claim.ID,
		GorId:    gor,
		BookedAt: time.Now(),
	}
	//api call only provide booking start and booking end
	helper.MapDtoToBooking(bookingDto, &booking)

	result, err := repository.CreateBookingRequest(booking)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	helper.MapBookingToDto(*result, &bookingDto)

	c.JSON(http.StatusOK, gin.H{
		"data": bookingDto,
	})

}

func GetUserBookingHistory(c *gin.Context) {
	token, _ := helper.ExtractBearerToken(c.GetHeader("Authorization"))
	claim, _ := services.VerifyToken(token)

	result, err := repository.SearchBookingHistory(claim.ID)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var bookings []dtos.BookingDto
	for _, val := range result {
		var valDto dtos.BookingDto
		helper.MapBookingToDto(*val, &valDto)
		bookings = append(bookings, valDto)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}
