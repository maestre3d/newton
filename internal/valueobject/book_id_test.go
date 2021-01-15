package valueobject

import (
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/stretchr/testify/assert"
)

var bookIdTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrBookIDOutOfRange},
	{"", ErrBookIDOutOfRange}, // will be above 128 char long
	{"123456789012345", ErrBookIDOutOfRange},
	{"1234567890123456", nil},
	{"", nil}, // will be 128 char long
	{gonanoid.Must(16), nil},
}

func TestNewBookID(t *testing.T) {
	for i, tt := range bookIdTestingSuite {
		if i == 1 {
			tt.in = populateString(129)
		} else if i == 4 {
			tt.in = populateString(128)
		}

		t.Run("New book id", func(t *testing.T) {
			id, err := NewBookID(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, id.Value())
		})
	}
}

func BenchmarkNewBookID(b *testing.B) {
	b.Run("Bench New book id", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = NewBookID(gonanoid.Must(16))
		}
	})
}
