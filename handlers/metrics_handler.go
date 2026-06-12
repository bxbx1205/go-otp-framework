package handlers

import (
	"net/http"

	"github.com/bxbx1205/go-otp-framework/metrics"

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
