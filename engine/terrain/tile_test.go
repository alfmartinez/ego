package terrain

import (
	"ego/engine/movement"
	"image"
	"testing"
)

func TestTile(t *testing.T) {
	tileType := CreateTileType("plain")
	tile := CreateTile(GridCoord{0, 0}, tileType, 10)

	if tile.Position() != image.Pt(0, 0) {
		t.Errorf("Tile should be at 0,0")
	}

	rect := tile.Rect()
	if !rect.Eq(image.Rect(0, 0, 10, 10)) {
		t.Errorf("Tile should have 0-10 rect, got %+v", rect)
	}

	if !tile.IsAt(movement.Loc(image.Pt(5, 5))) {
		t.Errorf("5,5 should be inside tile")
	}

	if tile.IsAt(movement.Loc(image.Pt(10, 10))) {
		t.Errorf("10,10 should be outside tile")
	}

	if tile.Size() != 10 {
		t.Errorf("Tile should have size 10, got %+v", tile.Size())
	}

}
