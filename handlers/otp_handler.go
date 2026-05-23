package handlers

import (
	"net/http"
	"otp-service/models"
	"otp-service/services"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context){

	var request models.SendOTPRequest

	err:= c.ShouldBindJSON(&request)

	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Request Body",
		})

		return

	}

	err= services.SendOTP(request.Phone)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send OTP",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully",
	})



}
func VerifyOTP(c *gin.Context) {

	var request models.VerifyOTPRequest


	err := c.ShouldBindJSON(&request)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}


	err = services.VerifyOTP(
		request.Phone,
		request.OTP,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "OTP verified successfully",
	})
}