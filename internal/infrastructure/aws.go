package infrastructure

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewAWSConfig allocates a new Amazon Web Services configuration
func NewAWSConfig(ctx context.Context) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx)
}

// NewAWSDynamoDB allocates a new Amazon Web Services DynamoDB client
func NewAWSDynamoDB(localCfg Configuration, cfg aws.Config) *dynamodb.Client {
	if localCfg.IsDev() {
		cfg.EndpointResolver = aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			if service == dynamodb.ServiceID {
				return aws.Endpoint{
					PartitionID:       "aws",
					HostnameImmutable: true,
					URL:               "http://localhost:8001",
					SigningRegion:     "localhost",
				}, nil
			}
			// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})
	}
	return dynamodb.NewFromConfig(cfg)
}

// NewAWSEventBridge allocates a new Amazon Web Services EventBridge client
func NewAWSEventBridge(cfg aws.Config) *eventbridge.Client {
	return eventbridge.NewFromConfig(cfg)
}

// NewAWSSimpleStorage allocates a new Amazon Web Services S3 (Simple Storage Service) client
func NewAWSSimpleStorage(newtonCfg Configuration, cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = newtonCfg.BucketRegion
	})
}
