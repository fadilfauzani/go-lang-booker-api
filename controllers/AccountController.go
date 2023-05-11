package controllers

import (
	"fmt"
	"net/http"
	"sporesappapi/dtos"
	"sporesappapi/helper"
	"sporesappapi/models"
	"sporesappapi/repository"
	"sporesappapi/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var (
		user = models.User{
			Balance: 10000,
		}
	)
	var registerDto = dtos.RegisterDto{}
	err := c.ShouldBind(&registerDto)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !helper.ValidateEmailPassword(registerDto.Email, registerDto.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Email is Not valid or Password length is less then 6"})
		return
	}
	if _, err := repository.SeachUserByEmail(registerDto.Email); err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Email is already used"})
		return
	}

	helper.MapRegisterToUser(registerDto, &user)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := repository.CreateUser(user)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	token, err := services.CreateToken(result.ID, result.Email)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dtos.AuthDto{Token: token})
}

func LoginUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var loginDto = dtos.LoginDto{}
	err := c.ShouldBind(&loginDto)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !helper.ValidateEmailPassword(loginDto.Email, loginDto.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Email is not valid or Password length is less then 6"})
		return
	}
	user, valid := repository.LoginByEmailPassword(loginDto.Email, loginDto.Password)
	if !valid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Email or Password is not valid"})
		return
	}
	token, err := services.CreateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dtos.AuthDto{Token: token})
}
