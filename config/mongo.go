package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var UserCollection *mongo.Collection
var OTPLogCollection *mongo.Collection
var APIKeyCollection *mongo.Collection
var ProviderCollection *mongo.Collection

func ConnectMongoDB(mongoURI string) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*10,
	)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected Successfully")

	MongoClient = client
	UserCollection = client.Database("otp_db").Collection("users")
	OTPLogCollection = client.Database("otp_db").Collection("otp_logs")
	APIKeyCollection = client.Database("otp_db").Collection("api_keys")
	ProviderCollection = client.Database("otp_db").Collection("providers")
}
