package infrastructure

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/maestre3d/newton/internal/event"
)

// UnmarshalSQSToSNS parses an Amazon Web Services Simple Queue Service (SQS) body field into the given
// events.SNSEntity
func UnmarshalSQSToSNS(b string, entity *events.SNSEntity) error {
	return json.Unmarshal([]byte(b), entity)
}

// UnmarshalSNSToEvent parses an Amazon Web Services Simple Notification Service (SNS) message field into the given
// event.DomainEvent
func UnmarshalSNSToEvent(b string, ev event.DomainEvent) error {
	return json.Unmarshal([]byte(b), ev)
}
