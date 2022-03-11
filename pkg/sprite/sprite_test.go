package sprite

import "testing"

func TestSprite(t *testing.T) {
	sprite := CreateSprite("test")
	if sprite.Path() != "test:0:0" {
		t.Errorf("Sprite should return expected path, got %s", sprite.Path())
	}

	if sprite.Size() != 1000 {
		t.Error("Sprite should have size 1000")
	}
}
