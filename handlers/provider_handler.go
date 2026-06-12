package handlers

import (
	"net/http"

	"github.com/bxbx1205/go-otp-framework/models"
	"github.com/bxbx1205/go-otp-framework/services"
	"github.com/bxbx1205/go-otp-framework/utils"

	"github.com/gin-gonic/gin"
)

func UpsertProvider(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.UpsertProviderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.UpsertProvider(userID, req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "provider saved successfully", nil)
}

func ListProviders(c *gin.Context) {
	userID := c.GetString("user_id")

	providers, err := services.ListProviders(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "providers fetched", providers)
}

func DeleteProvider(c *gin.Context) {
	userID := c.GetString("user_id")
	providerName := c.Param("provider")

	err := services.DeleteProvider(userID, providerName)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "provider deleted", nil)
}
