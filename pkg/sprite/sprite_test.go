package sprite

import "testing"

func TestSprite(t *testing.T) {
	sprite := CreateSprite("test")
	if sprite.Path() != "test" {
		t.Error("Sprite should return given path")
	}

	if sprite.Size() != 50 {
		t.Error("Sprite should have size 50")
	}
}
