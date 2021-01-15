package valueobject

import (
	"errors"
	"time"
)

// PublishYear concrete year an aggregate.Book was published
type PublishYear struct {
	value int
}

const publishYearMinValue = 0

var (
	publishYearMaxValue = time.Now().UTC().Year()
	// ErrPublishYearOutOfRange the given publish year was out of range, max is current year
	ErrPublishYearOutOfRange = errors.New("publish year is out of range [0, current year)")
)

// NewPublishYear creates and validates a PublishYear
func NewPublishYear(v int) (*PublishYear, error) {
	if err := ensurePublishYearLength(v); err != nil {
		return nil, err
	}
	return &PublishYear{value: v}, nil
}

func ensurePublishYearLength(v int) error {
	if v < publishYearMinValue || v > publishYearMaxValue {
		return ErrPublishYearOutOfRange
	}
	return nil
}
