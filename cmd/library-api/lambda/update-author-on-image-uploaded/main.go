package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/module"
	"github.com/maestre3d/newton/internal/persistence"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/subscriber"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	authorApp *application.Author
	logger    *zap.Logger
)

func handler(ctx context.Context, ev events.SQSEvent) {
	for _, message := range ev.Records {
		entity := events.SNSEntity{}
		if err := infrastructure.UnmarshalSQSToSNS(message.Body, &entity); err != nil {
			panic(err)
		}

		domainEv := event.AuthorImageUploaded{}
		if err := infrastructure.UnmarshalSNSToEvent(entity.Message, &domainEv); err != nil {
			panic(err)
		}
		infrastructure.LogSQSMessage(logger, message)
		logger.Info("marshal domain event", zap.Any("event", domainEv))
		if err := subscriber.UpdateAuthorOnImageUploaded(authorApp, ctx, domainEv); err != nil {
			panic(err)
		}
	}
}

func main() {
	app := fx.New(
		module.AWS,
		fx.Provide(
			func(cfg infrastructure.Configuration, db *dynamodb.Client) repository.Author {
				return persistence.NewAuthorDynamo(cfg, db)
			},
			application.NewAuthor,
		),
		fx.Invoke(
			func(l *zap.Logger) {
				logger = l
			},
			func(a *application.Author) {
				authorApp = a
			},
			func() {
				lambda.Start(handler)
			},
		),
	)
	app.Run()
}
