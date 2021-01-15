package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var coverTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrCoverInvalidURL},
	{"foo.com/", ErrCoverOutOfRange}, // will be above 2000 char long
	{"a.c/", ErrCoverOutOfRange},
	{"aex.", ErrCoverInvalidURL},
	{"aex12com", ErrCoverInvalidURL},
	{"https://cdn.newton.neutrinocorp.org", ErrCoverInvalidURL},
	{"https://cdn.newton.neutrinocorp.org/books/123.gif", ErrCoverInvalidExtension},
	{"https://cdn.newton.neutrinocorp.org/books/123.pdf", ErrCoverInvalidExtension},
	{"https://cdn.newton.neutrinocorp.org/books/123.jpg", nil},
	{"foo.com/", nil}, // will be 2000 char long
	{"https://cdn.newton.neutrinocorp.org/books/123.png", nil},
	{"https://cdn.newton.neutrinocorp.org/books/123.jpeg", nil},
	{"https://cdn.newton.neutrinocorp.org/books/123.webp", nil},
}

func TestNewCover(t *testing.T) {
	for i, tt := range coverTestingSuite {
		if i == 1 {
			tt.in += populateString(1989)
			tt.in += ".jpg"
		} else if i == 9 {
			tt.in += populateString(1988)
			tt.in += ".jpg"
		}
		t.Run("New cover", func(t *testing.T) {
			c, err := NewCover(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, c.Value())
		})
	}
}

func BenchmarkNewCover(b *testing.B) {
	b.Run("Bench New cover", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = NewCover("https://cdn.newton.neutrinocorp.org/books/123.png")
		}
	})
}
