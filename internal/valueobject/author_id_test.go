package valueobject

import (
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

var authorIDTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrAuthorIDOutOfRange},
	{"", ErrAuthorIDOutOfRange}, // will be above 128 char long
	{"123456789012345", ErrAuthorIDOutOfRange},
	{"1234567890123456", nil},
	{"", nil}, // will be 128 char long
	{gonanoid.Must(16), nil},
}

func TestNewAuthorID(t *testing.T) {
	for i, tt := range authorIDTestingSuite {
		if i == 1 {
			tt.in = populateString(129)
		} else if i == 4 {
			tt.in = populateString(128)
		}

		t.Run("New author id", func(t *testing.T) {
			id, err := NewAuthorID(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, id.Value())
		})
	}
}

func BenchmarkNewAuthorID(b *testing.B) {
	id := gonanoid.Must(16)
	b.Run("Bench New author id", func(b *testing.B) {
		var v AuthorID
		defer func() {
			// avoids v non-used
			if v != "" {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewAuthorID(id)
		}
	})
}
