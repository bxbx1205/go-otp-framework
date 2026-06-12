package repositories

import (
	"context"

	"github.com/myusername/otp-framework/config"
	"github.com/myusername/otp-framework/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertProvider(provider models.Provider) error {
	filter := bson.M{
		"user_id":  provider.UserID,
		"provider": provider.Provider,
	}

	update := bson.M{
		"$set": provider,
	}

	opts := options.Update().SetUpsert(true)

	_, err := config.ProviderCollection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func GetProviderByUserID(userID string, providerName string) (models.Provider, error) {
	var provider models.Provider
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return provider, err
	}

	err = config.ProviderCollection.FindOne(context.Background(), bson.M{
		"user_id":  objID,
		"provider": providerName,
	}).Decode(&provider)

	return provider, err
}

func ListProvidersByUser(userID string) ([]models.Provider, error) {
	var providers []models.Provider
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	cursor, err := config.ProviderCollection.Find(context.Background(), bson.M{"user_id": objID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &providers)
	return providers, err
}

func DeleteProvider(userID string, providerName string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = config.ProviderCollection.DeleteOne(context.Background(), bson.M{
		"user_id":  objID,
		"provider": providerName,
	})
	return err
}
