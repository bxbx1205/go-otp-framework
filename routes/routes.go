package routes

import (
	"otp-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){

	router.GET("/health",func(c *gin.Context) {

		c.JSON(200,gin.H{
			"Message":"Server running",
		})
	})

	otpRoutes:= router.Group("api/v1/otp")
	{
		otpRoutes.POST("/send",handlers.SendOTP)
		otpRoutes.POST("/verify", handlers.VerifyOTP)
	}
}