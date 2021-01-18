package valueobject

import (
	"time"

	"github.com/maestre3d/newton/internal/domain"
)

// PublishYear concrete year an aggregate.Book was published
type PublishYear int

const publishYearMinValue = 0

var (
	publishYearMaxValue = time.Now().UTC().Year()
	// ErrPublishYearOutOfRange the given publish year was out of range, max is current year
	ErrPublishYearOutOfRange = domain.NewOutOfRange("publish_year", publishYearMinValue, publishYearMaxValue)
)

// NewPublishYear creates and validates a PublishYear
func NewPublishYear(v int) (PublishYear, error) {
	if err := ensurePublishYearLength(v); err != nil {
		return 0, err
	}
	return PublishYear(v), nil
}

func ensurePublishYearLength(v int) error {
	if v < publishYearMinValue || v > publishYearMaxValue {
		return ErrPublishYearOutOfRange
	}
	return nil
}

// Value get the current value
func (y PublishYear) Value() int {
	return int(y)
}
