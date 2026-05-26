package repositories

import (
	"context"
	"otp-service/config"
	"otp-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func getUserCollection() *mongo.Collection {
	return config.MongoClient.
		Database("otp_services").
		Collection("users")
}

func CreateUser(user models.User) error {

	collection := getUserCollection()

	_, err := collection.InsertOne(
		context.Background(),
		user,
	)

	return err
}