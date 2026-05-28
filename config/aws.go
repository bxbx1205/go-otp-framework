package config

import (
	"context"
	"log"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var SNSClient *sns.Client

func ConnectAWS() {

	cfg,err:= awsConfig.LoadDefaultConfig(context.Background(),)

	if err!=nil {
		log.Fatal("failed to load AWS config : ",err)
	}

	SNSClient=sns.NewFromConfig(cfg)

	log.Println("AWS SNS Connected")


}