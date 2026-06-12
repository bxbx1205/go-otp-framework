package handlers

import (
	"net/http"
	"github.com/bxbx1205/go-otp-framework/models"
	"github.com/bxbx1205/go-otp-framework/services"
	"github.com/bxbx1205/go-otp-framework/utils"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context) {

	userID := c.GetString("user_id")

	var request models.SendOTPRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)

		return

	}

	err = services.SendOTP(userID, request.Phone)

	if err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}
	utils.SuccessResponse(
		c,
		http.StatusOK,
		"OTP sent successfully",
		nil,
	)

}
func VerifyOTP(c *gin.Context) {

	userID := c.GetString("user_id")

	var request models.VerifyOTPRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	token, err := services.VerifyOTP(
		userID,
		request.Phone,
		request.OTP,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "OTP verified",

			"token": token,
		},
	)
}
