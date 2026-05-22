package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client


func ConnectMongoDB() {
	mongoURI:=os.Getenv("MONGO_URI")

	ctx,cancel:= context.WithTimeout(
		context.Background(),
		time.Second*10,
	)

	defer cancel()

	client,err:= mongo.Connect(ctx,options.Client().ApplyURI(mongoURI),)

	if err != nil {
		log.Fatal(err)
	}
	
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected Successfully")


	MongoClient = client
}