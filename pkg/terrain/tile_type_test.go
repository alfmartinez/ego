package terrain

import "testing"

func TestCreateTileType(t *testing.T) {
	tileType := CreateTileType("")
	if tileType.Path() != "mario.png" {
		t.Errorf("Default should have mario as path")
	}
}
