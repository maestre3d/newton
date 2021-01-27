package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/module"
	"github.com/maestre3d/newton/internal/persistence"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/subscriber"
	"github.com/maestre3d/newton/pkg/controller"
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
			func(app *application.Author, logger *zap.Logger) subscriber.Subscriber {
				return subscriber.NewSubscriber(subscriber.NewUpdateAuthorOnImageUploaded(app), logger)
			},
			controller.NewEventAWS,
		),
		fx.Invoke(
			func(ctr controller.EventAWS) {
				lambda.Start(ctr.Handle)
			},
		),
	)
	app.Run()
}
