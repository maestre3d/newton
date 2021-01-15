package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var usernameTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrUsernameOutOfRange},
	{"", ErrUsernameOutOfRange}, // will be above 128 char long
	{"a", nil},
	{"aruizea", nil},
	{"", nil}, // will be 128 char long
}

func TestNewUsername(t *testing.T) {
	for i, tt := range usernameTestingSuite {
		if i == 1 {
			tt.in = populateString(129)
		} else if i == 4 {
			tt.in = populateString(128)
		}
		t.Run("New username", func(t *testing.T) {
			u, err := NewUsername(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, u.Value())
		})
	}
}
