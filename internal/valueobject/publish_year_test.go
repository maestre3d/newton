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

func BenchmarkNewPublishYear(b *testing.B) {
	b.Run("Bench New publish year", func(b *testing.B) {
		var y PublishYear
		defer func() {
			if y != 0 {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			y, _ = NewPublishYear(1984)
		}
	})
}
