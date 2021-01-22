package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/eventbus"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/persistence"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/service"
	"github.com/maestre3d/newton/pkg/controller"
	"github.com/maestre3d/newton/pkg/httputil"
	"go.uber.org/fx"
)

var infraModuleFx = fx.Options(
	fx.Provide(
		func() context.Context {
			return context.Background()
		},
		infrastructure.NewAWSConfig,
		infrastructure.NewAWSDynamoDB,
		infrastructure.NewConfiguration,
		infrastructure.NewZapLogger,
		infrastructure.NewAWSEventBridge,
		func(cfg infrastructure.Configuration, client *eventbridge.Client) event.Bus {
			return eventbus.NewEventBridge(cfg, client)
		},
		infrastructure.NewAWSSimpleStorage,
		func(cfg infrastructure.Configuration, c *s3.Client) service.FileBucket {
			return service.NewFileBucketS3(cfg, c)
		},
	),
)

func main() {
	app := fx.New(
		infraModuleFx,
		fx.Provide(
			func(cfg infrastructure.Configuration, db *dynamodb.Client) repository.Author {
				return persistence.NewAuthorDynamo(cfg, db)
			},
			application.NewAuthor,
			controller.NewAuthorHTTP,
			httputil.NewGorillaRouter,
		),
		fx.Invoke(
			httputil.StartServer,
		),
	)
	app.Run()
}
