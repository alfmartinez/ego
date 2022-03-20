package terrain

import (
	"ego/engine/movement"
	"image"
	"reflect"
	"testing"
)

func TestTile(t *testing.T) {
	tileType := CreateTileType("plain")
	tile := CreateTile(GridCoord{0, 0}, tileType, 10)

	if !reflect.DeepEqual(tile.Position(), movement.Location{}) {
		t.Errorf("Tile should be at 0,0")
	}

	rect := tile.Rect()
	if !rect.Eq(image.Rect(0, 0, 10, 10)) {
		t.Errorf("Tile should have 0-10 rect, got %+v", rect)
	}

	if tile.Size() != 10 {
		t.Errorf("Tile should have size 10, got %+v", tile.Size())
	}

}
