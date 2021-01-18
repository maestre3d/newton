package infrastructure

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

// NewAWSConfig allocates a new Amazon Web Services configuration
func NewAWSConfig(ctx context.Context) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("default"))
}

// NewAWSDynamoDB allocates a new Amazon Web Services DynamoDB client
func NewAWSDynamoDB(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg)
}

// NewAWSEventBridge allocates a new Amazon Web Services EventBridge client
func NewAWSEventBridge(cfg aws.Config) *eventbridge.Client {
	return eventbridge.NewFromConfig(cfg)
}
