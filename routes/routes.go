package routes

import (
	"github.com/bxbx1205/go-otp-framework/handlers"
	"github.com/bxbx1205/go-otp-framework/middleware"

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

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.GET(
			"/stats",
			handlers.AdminStats,
		)

		adminRoutes.GET(
			"/users",
			handlers.GetUsers,
		)

		adminRoutes.GET(
			"/otp-logs",
			handlers.GetOTPLogs,
		)

		adminRoutes.GET(
			"/providers",
			handlers.GetProviderStatus,
		)
	}

	protected := router.Group("/user")

	protected.Use(
		middleware.AuthMiddleware(),
	)

	{
		protected.GET(
			"/profile",
			handlers.Profile,
		)
	}

	authRoutes := router.Group("api/v1/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
	}

	apiKeyRoutes := router.Group("api/v1/api-keys")
	apiKeyRoutes.Use(middleware.AuthMiddleware())
	{
		apiKeyRoutes.POST("/", handlers.CreateAPIKey)
		apiKeyRoutes.GET("/", handlers.ListAPIKeys)
		apiKeyRoutes.DELETE("/:key", handlers.RevokeAPIKey)
	}

	providerRoutes := router.Group("api/v1/providers")
	providerRoutes.Use(middleware.AuthMiddleware())
	{
		providerRoutes.POST("/", handlers.UpsertProvider)
		providerRoutes.GET("/", handlers.ListProviders)
		providerRoutes.DELETE("/:provider", handlers.DeleteProvider)
	}

	otpRoutes := router.Group("api/v1/otp")
	otpRoutes.Use(middleware.APIKeyMiddleware())
	{
		otpRoutes.POST("/send", handlers.SendOTP)
		otpRoutes.POST("/verify", handlers.VerifyOTP)
	}
}
