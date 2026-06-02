package handlers

import (
	"net/http"

	"otp-service/metrics"

	"github.com/gin-gonic/gin"
)

func GetMetrics(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"otp_sent":     metrics.OTPSent,
			"otp_verified": metrics.OTPVerified,
			"sms_success":  metrics.SMSSuccess,
			"sms_failed":   metrics.SMSFailed,
		},
	)
}