package loader

import "testing"

func TestPreLoad(t *testing.T) {
	t.Run("CreateLoader", func(t *testing.T) {
		l := CreateSpriteLoader("pre_load")
		if _, ok := l.(*preLoad); !ok {
			t.Errorf("Expected *preLoad, got %T", l)
		}
	})

	t.Run("CreateLoader", func(t *testing.T) {
		l := CreateSpriteLoader("pre_load")
		l.Init()
	})
}
