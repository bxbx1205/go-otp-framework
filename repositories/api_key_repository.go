package repositories

import (
	"context"

	"github.com/myusername/otp-framework/config"
	"github.com/myusername/otp-framework/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAPIKey(apiKey models.APIKey) error {
	_, err := config.APIKeyCollection.InsertOne(context.Background(), apiKey)
	return err
}

func FindAPIKey(keyString string) (models.APIKey, error) {
	var apiKey models.APIKey
	err := config.APIKeyCollection.FindOne(context.Background(), bson.M{"api_key": keyString}).Decode(&apiKey)
	return apiKey, err
}

func ListAPIKeysByUser(userID string) ([]models.APIKey, error) {
	var keys []models.APIKey
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	cursor, err := config.APIKeyCollection.Find(context.Background(), bson.M{"user_id": objID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &keys)
	return keys, err
}

func RevokeAPIKey(keyString string, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = config.APIKeyCollection.UpdateOne(
		context.Background(),
		bson.M{"api_key": keyString, "user_id": objID},
		bson.M{"$set": bson.M{"revoked": true}},
	)
	return err
}
