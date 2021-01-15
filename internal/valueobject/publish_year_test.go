package valueobject

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var publishYearTestingSuite = []struct {
	in  int
	exp error
}{
	{-1, ErrPublishYearOutOfRange},
	{time.Now().UTC().Year() + 1, ErrPublishYearOutOfRange},
	{0, nil},
	{1, nil},
	{1984, nil},
	{time.Now().UTC().Year(), nil},
}

func TestNewPublishYear(t *testing.T) {
	for _, tt := range publishYearTestingSuite {
		t.Run("New publish year", func(t *testing.T) {
			y, err := NewPublishYear(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, y.Value())
		})
	}
}
