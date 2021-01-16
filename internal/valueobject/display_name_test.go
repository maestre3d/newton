package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var displayNameTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrDisplayNameOutOfRange},
	{"", ErrDisplayNameOutOfRange}, // will be above 256 char long
	{"I", ErrDisplayNameOutOfRange},
	{"Io", nil},
	{"Erwin Schrödinger", nil},
	{"", nil}, // will be 256 char long
}

func TestNewDisplayName(t *testing.T) {
	for i, tt := range displayNameTestingSuite {
		if i == 1 {
			tt.in = populateString(257)
		} else if i == 5 {
			tt.in = populateString(256)
		}
		t.Run("New Display name", func(t *testing.T) {
			name, err := NewDisplayName(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, name.Value())
		})
	}
}

func BenchmarkNewDisplayName(b *testing.B) {
	b.Run("Bench New Display name", func(b *testing.B) {
		var v *DisplayName
		defer func() {
			if v != nil {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewDisplayName("Max Planck")
		}
	})
}
