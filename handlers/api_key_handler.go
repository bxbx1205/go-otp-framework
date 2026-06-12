package handlers

import (
	"net/http"

	"github.com/myusername/otp-framework/models"
	"github.com/myusername/otp-framework/services"
	"github.com/myusername/otp-framework/utils"

	"github.com/gin-gonic/gin"
)

func CreateAPIKey(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	apiKey, err := services.CreateAPIKey(userID, req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "api key created successfully", apiKey)
}

func ListAPIKeys(c *gin.Context) {
	userID := c.GetString("user_id")

	keys, err := services.ListAPIKeys(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "api keys fetched", keys)
}

func RevokeAPIKey(c *gin.Context) {
	userID := c.GetString("user_id")
	keyToRevoke := c.Param("key")

	err := services.RevokeAPIKey(keyToRevoke, userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "api key revoked", nil)
}
