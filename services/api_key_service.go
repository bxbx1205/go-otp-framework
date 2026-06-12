package services

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/myusername/otp-framework/models"
	"github.com/myusername/otp-framework/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateRandomKey() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return "api_live_" + hex.EncodeToString(bytes)
}

func CreateAPIKey(userID string, req models.CreateAPIKeyRequest) (models.APIKey, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.APIKey{}, err
	}

	apiKey := models.APIKey{
		UserID:    objID,
		Name:      req.Name,
		Key:       generateRandomKey(),
		Revoked:   false,
		CreatedAt: time.Now(),
	}

	err = repositories.CreateAPIKey(apiKey)
	return apiKey, err
}

func ListAPIKeys(userID string) ([]models.APIKey, error) {
	return repositories.ListAPIKeysByUser(userID)
}

func RevokeAPIKey(key string, userID string) error {
	return repositories.RevokeAPIKey(key, userID)
}
