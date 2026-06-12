package config

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var SNSClient *sns.Client

func ConnectAWS(accessKey string, secretKey string, region string) {

	cfg := aws.Config{
		Region:      region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}

	SNSClient = sns.NewFromConfig(cfg)

	log.Println("AWS SNS Connected")

}
