package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var bookResourceTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrBookResourceInvalidURL},
	{"foo.com/", ErrBookResourceOutOfRange}, // will be above 2000 char long
	{"a.c/", ErrBookResourceOutOfRange},
	{"aex.", ErrBookResourceInvalidURL},
	{"aex12com", ErrBookResourceInvalidURL},
	{"https://cdn.newton.neutrinocorp.org", ErrBookResourceInvalidURL},
	{"https://cdn.newton.neutrinocorp.org/books/123.pptx", ErrBookResourceInvalidExtension},
	{"https://cdn.newton.neutrinocorp.org/books/123.docx", ErrBookResourceInvalidExtension},
	{"https://cdn.newton.neutrinocorp.org/books/123.pdf", nil},
	{"foo.com/", nil}, // will be 2000 char long
}

func TestNewBookResource(t *testing.T) {
	for i, tt := range bookResourceTestingSuite {
		if i == 1 {
			tt.in += populateString(1989)
			tt.in += ".pdf"
		} else if i == 9 {
			tt.in += populateString(1988)
			tt.in += ".pdf"
		}
		t.Run("New book resource", func(t *testing.T) {
			r, err := NewBookResource(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, r.Value())
		})
	}
}

func BenchmarkNewBookResource(b *testing.B) {
	b.Run("Bench New book resource", func(b *testing.B) {
		var v *BookResource
		defer func() {
			if v != nil {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewBookResource("https://cdn.newton.neutrinocorp.org/books/123.pdf")
		}
	})
}
