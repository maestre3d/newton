package aggregate

import (
	"testing"

	"github.com/maestre3d/newton/internal/valueobject"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/stretchr/testify/assert"
)

var authorStubDTO = struct {
	ID       valueobject.AuthorID
	Name     valueobject.DisplayName
	CreateBy valueobject.Username
	Image    valueobject.Image
}{}

func init() {
	authorStubDTO.ID, _ = valueobject.NewAuthorID(gonanoid.MustID(16))
	authorStubDTO.Image, _ = valueobject.NewImage("https://foo.com/picture.jpg")
	authorStubDTO.CreateBy, _ = valueobject.NewUsername("aruiz")
	authorStubDTO.Name, _ = valueobject.NewDisplayName("Ludwig Boltzmann")
}

func TestNewAuthor(t *testing.T) {
	t.Run("New author", func(t *testing.T) {
		author := NewAuthor(authorStubDTO.ID, authorStubDTO.Name, authorStubDTO.CreateBy,
			authorStubDTO.Image)
		assert.Equal(t, authorStubDTO.ID.Value(), author.ID.Value())
		assert.Equal(t, authorStubDTO.CreateBy.Value(), author.CreateBy.Value())
		assert.Equal(t, authorStubDTO.Image.Value(), author.Image.Value())
		assert.Equal(t, authorStubDTO.Name.Value(), author.DisplayName.Value())
		assert.Equal(t, 1, len(author.PullEvents()))
		assert.NotEmpty(t, author.Metadata.CreateTime)
		assert.NotEmpty(t, author.Metadata.UpdateTime)
		assert.True(t, author.Metadata.State)
		assert.False(t, author.Metadata.MarkAsRemoval)
	})
}

func BenchmarkNewAuthor(b *testing.B) {
	b.Run("Bench New author", func(b *testing.B) {
		var v *Author
		defer func() {
			// avoids v non-used
			if v != nil {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v = NewAuthor(authorStubDTO.ID, authorStubDTO.Name, authorStubDTO.CreateBy,
				authorStubDTO.Image)
		}
	})
}

var authorChangeStateTestingSuite = []struct {
	in bool
}{
	{false},
	{true},
}

func TestAuthor_ChangeState(t *testing.T) {
	for _, tt := range authorChangeStateTestingSuite {
		t.Run("Author change state", func(t *testing.T) {
			author := NewAuthor(authorStubDTO.ID, authorStubDTO.Name, authorStubDTO.CreateBy,
				authorStubDTO.Image)
			author.ChangeState(tt.in)
			assert.Equal(t, tt.in, author.Metadata.State)
			assert.Equal(t, 2, len(author.PullEvents())) // created, and state change
		})
	}
}

func TestAuthor_Update(t *testing.T) {
	t.Run("Author update", func(t *testing.T) {
		author := NewAuthor(authorStubDTO.ID, authorStubDTO.Name, authorStubDTO.CreateBy,
			authorStubDTO.Image)
		oldUpdateTime := author.Metadata.UpdateTime
		image, _ := valueobject.NewImage("https://foo.com/bar.jpg")
		createdBy, _ := valueobject.NewUsername("br1")
		name, _ := valueobject.NewDisplayName("Max Born")
		author.Update(name, createdBy, image)
		assert.NotEqual(t, authorStubDTO.CreateBy.Value(), author.CreateBy.Value())
		assert.NotEqual(t, authorStubDTO.Image.Value(), author.Image.Value())
		assert.NotEqual(t, authorStubDTO.Name.Value(), author.DisplayName.Value())
		assert.NotEqualValues(t, oldUpdateTime, author.Metadata.UpdateTime)
		assert.Equal(t, 2, len(author.PullEvents()))
	})
}

func TestAuthor_Remove(t *testing.T) {
	t.Run("Author remove", func(t *testing.T) {
		author := NewAuthor(authorStubDTO.ID, authorStubDTO.Name, authorStubDTO.CreateBy,
			authorStubDTO.Image)
		author.Remove()
		assert.True(t, author.Metadata.MarkAsRemoval)
		assert.Equal(t, 2, len(author.PullEvents()))
	})
}
