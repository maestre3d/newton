package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/eventbus"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/persistence"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/pkg/controller"
	"github.com/maestre3d/newton/pkg/httputil"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
		func(logger *zap.Logger) event.Bus {
			return eventbus.NewLocal(logger)
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
