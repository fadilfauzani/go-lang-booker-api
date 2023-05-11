package main

import (
	"fmt"
	"sporesappapi/controllers"
	"sporesappapi/data"
	middleware "sporesappapi/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// fmt.Print(os.Getenv("DB_HOST"))
	data.StartDB()

	r := gin.Default()
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)

	authGroup := r.Group("")
	authGroup.Use(middleware.Authorize)

	userGroup := authGroup.Group("/user")
	userGroup.GET("", controllers.GetProfile)
	userGroup.PATCH("", controllers.EditProfile)
	userGroup.PATCH("/changePassword", controllers.ChangePassword)

	gorGroup := authGroup.Group("/gor")
	gorGroup.GET("", controllers.GetGor)
	gorGroup.POST("", controllers.CreateGor)

	bookingGroup := authGroup.Group("/booking")
	bookingGroup.POST("/:gor", controllers.CreateBookingRequest)
	bookingGroup.GET("/history", controllers.GetUserBookingHistory)

	paymentGroup := authGroup.Group("/pay")
	paymentGroup.PATCH("/:booking", controllers.PayBooking)
	r.Run(":8080")

}
