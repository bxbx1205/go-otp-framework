package middleware

import (
	"net/http"
	"strings"

	"github.com/myusername/otp-framework/repositories"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyHeader := c.GetHeader("x-api-key")

		if apiKeyHeader == "" || !strings.HasPrefix(apiKeyHeader, "api_live_") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "valid api key required"})
			c.Abort()
			return
		}

		apiKey, err := repositories.FindAPIKey(apiKeyHeader)
		if err != nil || apiKey.Revoked {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or revoked api key"})
			c.Abort()
			return
		}

		c.Set("user_id", apiKey.UserID.Hex())
		c.Next()
	}
}
