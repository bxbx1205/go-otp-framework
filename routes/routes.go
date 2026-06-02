package routes

import (
	"otp-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET(
		"/health",
		handlers.HealthCheck,
	)

	router.GET(
		"/dlq",
		handlers.GetDLQ,
	)

	router.GET(
	"/metrics",
	handlers.GetMetrics,
)

router.GET(
	"/queue-status",
	handlers.QueueStatus,
)

	otpRoutes := router.Group("api/v1/otp")
	{
		otpRoutes.POST("/send", handlers.SendOTP)
		otpRoutes.POST("/verify", handlers.VerifyOTP)
	}
}
