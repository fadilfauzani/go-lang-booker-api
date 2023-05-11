package controllers

import (
	"fmt"
	"net/http"
	"sporesappapi/dtos"
	"sporesappapi/helper"
	"sporesappapi/models"
	"sporesappapi/repository"

	"github.com/gin-gonic/gin"
)

func GetGor(c *gin.Context) {
	query, _ := c.GetQuery("query")
	category, _ := c.GetQuery("category")

	fmt.Println(query, category)
	result, err := repository.SearchGor(query, category)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var gors []dtos.GorDto
	for _, val := range result {
		var valDto dtos.GorDto
		helper.MapGorToDto(*val, &valDto)
		gors = append(gors, valDto)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gors,
	})
}

func CreateGor(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var (
		gor    models.Gor
		gorDto dtos.GorDto
	)
	err := c.ShouldBind(&gorDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	helper.MapDtoToGor(gorDto, &gor)

	result, err := repository.CreateGor(gor)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	helper.MapGorToDto(*result, &gorDto)

	c.JSON(http.StatusOK, gin.H{
		"data": gorDto,
	})

}
