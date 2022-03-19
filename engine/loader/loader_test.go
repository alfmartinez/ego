package loader

import (
	"image"
	"testing"
)

type mockLoader struct{}

func (l *mockLoader) Init() {

}
func (l *mockLoader) GetSprite(string, uint) image.Image {
	return image.NewNRGBA(image.Rect(0, 0, 100, 100))
}

func TestLoader(t *testing.T) {
	RegisterLoader("foo", func() Loader {
		return &mockLoader{}
	})
	l := CreateSpriteLoader("foo")
	if _, ok := l.(*mockLoader); !ok {
		t.Error("Expected mockLoader")
	}
}

func TestLoadSpriteSheet(t *testing.T) {
	t.Run("Load missing sprite", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("Should have panicked")
			}
		}()
		loadSpriteSheet("foo.png")
	})

	t.Run("Load missing sprite", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("Should have panicked")
			}
		}()
		loadSpriteSheet("gif-esempio.gif")
	})
}
