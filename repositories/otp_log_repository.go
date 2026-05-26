package repositories

import (
	"context"
	"otp-service/config"
	"otp-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func getOTPLogCollection() *mongo.Collection {
	return config.MongoClient.
		Database("otp_service").
		Collection("otp_logs")
}

func CreateOTPLog(log models.OTPLog) error {

	collection := getOTPLogCollection()

	_, err := collection.InsertOne(
		context.Background(),
		log,
	)

	return err
}