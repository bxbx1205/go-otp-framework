package services

import (
	"errors"

	"github.com/myusername/otp-framework/models"
	"github.com/myusername/otp-framework/repositories"
	"github.com/myusername/otp-framework/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpsertProvider(userID string, req models.UpsertProviderRequest) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	provider := models.Provider{
		UserID:   objID,
		Provider: req.Provider,
	}

	if req.Provider == "twilio" {
		encToken, err := utils.Encrypt(req.AuthToken)
		if err != nil {
			return err
		}
		provider.AccountSID = req.AccountSID
		provider.PhoneNumber = req.PhoneNumber
		provider.EncryptedAuthToken = encToken
	} else if req.Provider == "aws" {
		encSecret, err := utils.Encrypt(req.SecretKey)
		if err != nil {
			return err
		}
		provider.AccessKey = req.AccessKey
		provider.Region = req.Region
		provider.EncryptedSecretKey = encSecret
	} else {
		return errors.New("unsupported provider")
	}

	return repositories.UpsertProvider(provider)
}

func ListProviders(userID string) ([]models.Provider, error) {
	return repositories.ListProvidersByUser(userID)
}

func DeleteProvider(userID string, providerName string) error {
	return repositories.DeleteProvider(userID, providerName)
}
