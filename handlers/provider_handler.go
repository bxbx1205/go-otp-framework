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

func AddTwilioProvider(c *gin.Context) {
	userID := c.GetString("user_id")

	var input struct {
		SID   string `json:"sid" binding:"required"`
		Token string `json:"token" binding:"required"`
		Phone string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	req := models.UpsertProviderRequest{
		Provider:    "twilio",
		AccountSID:  input.SID,
		AuthToken:   input.Token,
		PhoneNumber: input.Phone,
	}

	if err := services.UpsertProvider(userID, req); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "twilio provider added successfully", nil)
}

func AddAWSProvider(c *gin.Context) {
	userID := c.GetString("user_id")

	var input struct {
		AccessKey string `json:"accessKey" binding:"required"`
		SecretKey string `json:"secretKey" binding:"required"`
		Region    string `json:"region" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	req := models.UpsertProviderRequest{
		Provider:  "aws",
		AccessKey: input.AccessKey,
		SecretKey: input.SecretKey,
		Region:    input.Region,
	}

	if err := services.UpsertProvider(userID, req); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "aws provider added successfully", nil)
}
