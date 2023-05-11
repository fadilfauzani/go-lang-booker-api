package controllers

import (
	"net/http"
	"sporesappapi/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

//just to simulate payment system

func PayBooking(c *gin.Context) {
	bookingId, err := strconv.Atoi(c.Param("booking"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid Booking Parameter"})
		return
	}

	err = repository.PayBooking(bookingId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
