package valueobject

import (
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

var categoryIDTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrCategoryIDOutOfRange},
	{"", ErrCategoryIDOutOfRange}, // will be above 128 char long
	{"123456789012345", ErrCategoryIDOutOfRange},
	{"1234567890123456", nil},
	{"", nil}, // will be 128 char long
	{gonanoid.Must(16), nil},
}

func TestNewCategoryID(t *testing.T) {
	for i, tt := range categoryIDTestingSuite {
		if i == 1 {
			tt.in = populateString(129)
		} else if i == 4 {
			tt.in = populateString(128)
		}

		t.Run("New category id", func(t *testing.T) {
			id, err := NewCategoryID(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, id.Value())
		})
	}
}

func BenchmarkNewCategoryID(b *testing.B) {
	id := gonanoid.Must(16)
	b.Run("Bench New category id", func(b *testing.B) {
		var v CategoryID
		defer func() {
			if v != "" {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewCategoryID(id)
		}
	})
}
