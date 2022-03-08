package terrain

import "testing"

func TestTerrain(t *testing.T) {
	t.Run("GetTerrain should return nil if terrain has not been created", func(t *testing.T) {
		terrainSingleton = nil
		o := GetTerrain()
		if o != nil {
			t.Errorf("GetTerrain should return nil, got %+v", o)
		}
	})

	t.Run("GetTerrain should return Terrain if terrain has been created", func(t *testing.T) {
		terrainSingleton = nil
		g := CreateGrid(1, 1, func(t Tile) {})
		o := GetTerrain()
		if o != g {
			t.Errorf("GetTerrain should return created grid, got %+v", o)
		}
	})
}
