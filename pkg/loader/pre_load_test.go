package loader

import (
	"testing"

	"github.com/spf13/viper"
)

func TestPreLoad(t *testing.T) {
	viper.Set("sheets", Sheets{
		"sheet": Sheet{
			Path: "sheet.png",
			Rect: struct {
				W int
				H int
			}{
				W: 100,
				H: 100,
			},
			Sizes: []uint{100, 50},
		}})
	t.Run("CreateLoader", func(t *testing.T) {
		l := CreateSpriteLoader("pre_load")
		if _, ok := l.(*preLoad); !ok {
			t.Errorf("Expected *preLoad, got %T", l)
		}
	})

	t.Run("Loader can return sprite", func(t *testing.T) {
		l := CreateSpriteLoader("pre_load")
		l.Init()
		sprite := l.GetSprite("sheet:0:0", 100)
		if sprite == nil {
			t.Errorf("sprite should be available")
		}
		if sprite.Bounds().Dx() != 100 {
			t.Errorf("GetSprite should return sprite in expected size")
		}
	})

	t.Run("Loader can return sprite, malformed 1", func(t *testing.T) {
		l := CreateSpriteLoader("pre_load")

		l.Init()

		defer func() {
			r := recover()
			if r == nil {
				t.Error("Should have panicked")
			}
		}()
		l.GetSprite("sheet:A:0", 100)

	})

	t.Run("Loader can return sprite, malformed 2", func(t *testing.T) {
		l := CreateSpriteLoader("pre_load")

		l.Init()

		defer func() {
			r := recover()
			if r == nil {
				t.Error("Should have panicked")
			}
		}()
		l.GetSprite("sheet:0:A", 100)

	})
}
