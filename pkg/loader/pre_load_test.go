package loader

import "testing"

func TestPreLoad(t *testing.T) {
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
}
