package sprite

import "testing"

func TestSprite(t *testing.T) {
	t.Run("CreateSprite init", func(t *testing.T) {
		sprite := CreateSprite("test")
		if sprite.Path() != "test:0:0" {
			t.Errorf("Sprite should return expected path, got %s", sprite.Path())
		}

		if sprite.Size() != 1000 {
			t.Error("Sprite should have size 1000")
		}
	})

	t.Run("CreateSprite Frame Set", func(t *testing.T) {
		sprite := CreateSprite("test")
		sprite.Frame(12, 3)
		if sprite.Path() != "test:12:3" {
			t.Errorf("Sprite should return expected path, got %s", sprite.Path())
		}

		if sprite.Size() != 1000 {
			t.Error("Sprite should have size 1000")
		}
	})
}
