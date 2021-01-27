package main

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/module"
	"github.com/maestre3d/newton/internal/persistence"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/pkg/controller"
	"github.com/maestre3d/newton/pkg/httputil"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		module.AWS,
		fx.Provide(
			func(cfg infrastructure.Configuration, db *dynamodb.Client, logger *zap.Logger) repository.Author {
				return repository.NewAuthor(persistence.NewAuthorDynamo(cfg, db), logger)
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
