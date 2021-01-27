package module

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/maestre3d/newton/internal/eventbus"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// AWS Uber Fx's infrastructure Amazon Web Services (AWS) module
var AWS = fx.Options(
	fx.Provide(
		func() context.Context {
			return context.Background()
		},
		infrastructure.NewAWSConfig,
		infrastructure.NewAWSDynamoDB,
		infrastructure.NewConfiguration,
		infrastructure.NewZapLogger,
		infrastructure.NewAWSEventBridge,
		func(cfg infrastructure.Configuration, client *eventbridge.Client, logger *zap.Logger) eventbus.Bus {
			return eventbus.NewBus(eventbus.NewEventBridge(cfg, client), logger)
		},
		infrastructure.NewAWSSimpleStorage,
		func(cfg infrastructure.Configuration, c *s3.Client) service.FileBucket {
			return service.NewFileBucketS3(cfg, c)
		},
	),
)
