package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var Client *dynamodb.Client

func Init() {
	config, error := config.LoadDefaultConfig(context.TODO())
	if error != nil {
		panic("error :: " + error.Error())
	}
	Client = dynamodb.NewFromConfig(config)
}
