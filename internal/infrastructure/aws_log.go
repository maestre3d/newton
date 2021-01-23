package infrastructure

import (
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

// LogSQSMessage logs the given Amazon Web Services Simple Queue Service (SQS) message
func LogSQSMessage(logger *zap.Logger, msg events.SQSMessage) {
	logger.Info("received sqs message",
		zap.String("message_id", msg.MessageId),
		zap.String("event_source", msg.EventSource),
		zap.String("body", msg.Body),
	)
}
