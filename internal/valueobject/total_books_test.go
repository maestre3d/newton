package valueobject

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var totalBooksTestingSuite = []struct {
	in  int64
	exp error
}{
	{-1, ErrTotalBooksOutOfRange},
	{math.MaxInt64, nil},
	{0, nil},
	{1, nil},
}

func TestNewTotalBooks(t *testing.T) {
	for _, tt := range totalBooksTestingSuite {
		t.Run("New total books", func(t *testing.T) {
			total, err := NewTotalBooks(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, uint64(tt.in), total.Value())
		})
	}
}

func BenchmarkNewTotalBooks(b *testing.B) {
	b.Run("Bench New total books", func(b *testing.B) {
		var v TotalBooks
		defer func() {
			if v != 0 {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewTotalBooks(10)
		}
	})
}
