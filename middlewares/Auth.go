package middleware

import (
	"net/http"
	"sporesappapi/helper"
	"sporesappapi/services"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	token, err := helper.ExtractBearerToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	_, ok := services.VerifyToken(token)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
