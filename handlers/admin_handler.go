package handlers

import (
	"net/http"
	"os"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/metrics"
	"github.com/bxbx1205/go-otp-framework/repositories"
	"github.com/bxbx1205/go-otp-framework/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	users, err :=
		repositories.GetAllUsers()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		users,
	)
}

func GetOTPLogs(c *gin.Context) {

	logs, err :=
		repositories.GetAllOTPLogs()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		logs,
	)
}

func GetProviderStatus(
	c *gin.Context,
) {

	provider :=
		os.Getenv(
			"SMS_PROVIDER",
		)

	c.JSON(
		http.StatusOK,
		gin.H{
			"provider": provider,
			"status":   "active",
		},
	)
}

func AdminStats(
	c *gin.Context,
) {

	queueSize, _ :=
		config.RedisClient.LLen(
			config.Ctx,
			services.SMSQueue,
		).Result()

	dlqSize, _ :=
		config.RedisClient.LLen(
			config.Ctx,
			services.SMSDLQ,
		).Result()

	users, _ :=
		repositories.GetAllUsers()

	c.JSON(
		http.StatusOK,
		gin.H{
			"users": len(users),

			"otp_sent": metrics.OTPSent,

			"otp_verified": metrics.OTPVerified,

			"sms_success": metrics.SMSSuccess,

			"sms_failed": metrics.SMSFailed,

			"queue_size": queueSize,

			"dlq_size": dlqSize,
		},
	)
}
