package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var imageTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrImageInvalidURL},
	{"foo.com/", ErrImageOutOfRange}, // will be above 2000 char long
	{"a.c/", ErrImageOutOfRange},
	{"aex.", ErrImageInvalidURL},
	{"aex12com", ErrImageInvalidURL},
	{"https://cdn.newton.neutrinocorp.org", ErrImageInvalidURL},
	{"https://cdn.newton.neutrinocorp.org/books/123.gif", ErrImageInvalidExtension},
	{"https://cdn.newton.neutrinocorp.org/books/123.pdf", ErrImageInvalidExtension},
	{"https://cdn.newton.neutrinocorp.org/books/123.jpg", nil},
	{"foo.com/", nil}, // will be 2000 char long
	{"https://cdn.newton.neutrinocorp.org/books/123.png", nil},
	{"https://cdn.newton.neutrinocorp.org/books/123.jpeg", nil},
	{"https://cdn.newton.neutrinocorp.org/books/123.webp", nil},
}

func TestNewImage(t *testing.T) {
	for i, tt := range imageTestingSuite {
		if i == 1 {
			tt.in += populateString(1989)
			tt.in += ".jpg"
		} else if i == 9 {
			tt.in += populateString(1988)
			tt.in += ".jpg"
		}
		t.Run("New Image", func(t *testing.T) {
			c, err := NewImage(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, c.Value())
		})
	}
}

func BenchmarkNewImage(b *testing.B) {
	b.Run("Bench New Image", func(b *testing.B) {
		var v Image
		defer func() {
			if v != "" {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewImage("https://cdn.newton.neutrinocorp.org/books/123.png")
		}
	})
}
