package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var titleTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrTitleOutOfRange},
	{"", ErrTitleOutOfRange}, // will be above 256 char long
	{"I", nil},
	{"The Principles of Quantum Mechanics", nil},
	{"", nil}, // will be 256 char long
}

func TestNewTitle(t *testing.T) {
	for i, tt := range titleTestingSuite {
		if i == 1 {
			tt.in = populateString(257)
		} else if i == 4 {
			tt.in = populateString(256)
		}
		t.Run("New title", func(t *testing.T) {
			title, err := NewTitle(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, title.Value())
		})
	}
}

func BenchmarkNewTitle(b *testing.B) {
	b.Run("Bench New title", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = NewTitle("Statistical Thermodynamics")
		}
	})
}
