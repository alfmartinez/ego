package terrain

import "testing"

func TestCreateTileType(t *testing.T) {
	tileType := CreateTileType("")
	if tileType.Path() != "sheet:0:0" {
		t.Errorf("Default should have mario as path")
	}
}
