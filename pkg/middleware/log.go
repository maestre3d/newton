package middleware

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/maestre3d/newton/internal/event"
	"go.uber.org/zap"
)

// LogDomainEvent logs the given domain event
func LogDomainEvent(logger *zap.Logger, ev event.DomainEvent) {
	logger.Info("received domain event", zap.Any("domain_event", ev))
}

// LogSQSMessage logs the given Amazon Web Services Simple Queue Service (SQS) message
func LogSQSMessage(logger *zap.Logger, msg events.SQSMessage) {
	logger.Info("received sqs message",
		zap.String("message_id", msg.MessageId),
		zap.String("event_source", msg.EventSource),
		zap.String("body", msg.Body),
	)
}

// LogSNSEntity logs the given Amazon Web Services Simple Notification Service (SNS) entity
func LogSNSEntity(logger *zap.Logger, ent events.SNSEntity) {
	logger.Info("received sns entity",
		zap.String("message_id", ent.MessageID),
		zap.String("type", ent.Type),
		zap.String("message", ent.Message),
		zap.String("subject", ent.Subject),
		zap.Time("timestamp", ent.Timestamp),
	)
}
