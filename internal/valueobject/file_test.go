package valueobject

import (
	"io"
	"testing"

	"github.com/maestre3d/newton/internal/domain"
	"github.com/stretchr/testify/assert"
)

type fileStub struct{}

func (s fileStub) Read(p []byte) (int, error) {
	return 0, nil
}

var fileTestingSuite = []struct {
	filename     string
	size         int64
	file         io.Reader
	expName      string
	expSize      int64
	expExtension string
}{
	{"", 0, nil, "", 0, ""},
	{"foo", 0, nil, "foo", 0, ""},
	{"bar", domain.KibiByte * 100, nil, "bar", domain.KibiByte * 100, ""},
	{"baz", domain.KibiByte * 100, fileStub{}, "baz", domain.KibiByte * 100,
		""},
	{"foo.png", domain.MebiByte * 2, fileStub{}, "foo", domain.MebiByte * 2,
		"png"},
}

func TestNewFile(t *testing.T) {
	for _, tt := range fileTestingSuite {
		t.Run("New file", func(t *testing.T) {
			file := NewFile(tt.filename, tt.size, tt.file)
			assert.Equal(t, tt.expName, file.Name)
			assert.Equal(t, tt.expExtension, file.Extension)
			assert.Equal(t, tt.expSize, file.Size)
		})
	}
}
