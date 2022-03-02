package loader

import (
	"image"
	"testing"
	"time"
)

func ExecutionTime(f func()) time.Duration {
	now := time.Now()
	f()
	return time.Since(now)
}

func TestCreateSpriteLoaderOnDemand(t *testing.T) {
	loader := CreateSpriteLoader("on_demand")
	if _, ok := loader.(*onDemandLoader); !ok {
		t.Error("Create Loader should return on demand loader")
	}
}

func TestGetSpriteExisting(t *testing.T) {
	loader := CreateSpriteLoader("on_demand")
	if _, ok := loader.(*onDemandLoader); !ok {
		t.Error("Create Loader should return on demand loader")
	}
	var i image.Image
	origDuration := ExecutionTime(func() {
		i = loader.GetSprite("mario.png", 1)
	})

	if !i.Bounds().Eq(image.Rect(0, 0, 1, 1)) {
		t.Error("GetSprite should returned resized sprite")
	}

	nextDuration := ExecutionTime(func() {
		i = loader.GetSprite("mario.png", 1)
	})

	if nextDuration > origDuration {
		t.Errorf("Second call should be fast: %+v versus %+v", origDuration, nextDuration)
	}
}

func TestGetSpriteMissingFile(t *testing.T) {
	loader := CreateSpriteLoader("on_demand")
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Should have panicked")
		}
	}()
	loader.GetSprite("missing.png", 1)
}

func TestGetSpriteUnknownFormat(t *testing.T) {
	loader := CreateSpriteLoader("on_demand")
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Should have panicked")
		}
	}()
	loader.GetSprite("Transparent.gif", 1)
}
