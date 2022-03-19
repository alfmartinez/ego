package sprite

import "testing"

func TestSprite(t *testing.T) {
	t.Run("CreateSprite init", func(t *testing.T) {
		sprite := CreateSprite("test", 512)
		if sprite.Path() != "test:0:0" {
			t.Errorf("Sprite should return expected path, got %s", sprite.Path())
		}

		if sprite.Size() != 512 {
			t.Error("Sprite should have size 512")
		}
	})

	t.Run("CreateSprite Frame Set", func(t *testing.T) {
		sprite := CreateSprite("test", 50)
		sprite.Frame(12, 3)
		if sprite.Path() != "test:12:3" {
			t.Errorf("Sprite should return expected path, got %s", sprite.Path())
		}

		if sprite.Size() != 50 {
			t.Error("Sprite should have size 50")
		}
	})
}
