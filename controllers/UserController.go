package controllers

import (
	"net/http"
	"sporesappapi/dtos"
	"sporesappapi/helper"
	"sporesappapi/repository"
	"sporesappapi/services"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	token, _ := helper.ExtractBearerToken(c.GetHeader("Authorization"))
	claim, _ := services.VerifyToken(token)

	user, err := repository.SeachUserByEmail(claim.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": "Account Not Found"})
		return
	}
	var profile dtos.ProfileDto
	helper.MapUserToProfile(*user, &profile)
	c.JSON(http.StatusOK, gin.H{
		"data": profile,
	})
}

func EditProfile(c *gin.Context){
	token, _ := helper.ExtractBearerToken(c.GetHeader("Authorization"))
	claim, _ := services.VerifyToken(token)

	user, err := repository.SeachUserByEmail(claim.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": "Account Not Found"})
		return
	}
	var profileDto = dtos.ProfileDto{}
	err = c.ShouldBind(&profileDto)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	helper.MapProfileToUser(profileDto, user)
	
	result, err := repository.UpdateUser(*user)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var newprofile dtos.ProfileDto
	helper.MapUserToProfile(*result, &newprofile)
	c.JSON(http.StatusOK, gin.H{
		"data": newprofile,
	})
}

func ChangePassword(c *gin.Context){
	token, _ := helper.ExtractBearerToken(c.GetHeader("Authorization"))
	claim, _ := services.VerifyToken(token)

	user, err := repository.SeachUserByEmail(claim.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": "Account Not Found"})
		return
	}
	var passwordDto = dtos.PasswordDto{}
	err = c.ShouldBind(&passwordDto)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = passwordDto.Password

	_, err = repository.ChangePasswordUser(*user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}