package terrain

import (
	"testing"

	"github.com/spf13/viper"
)

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
		RegisterTileType("plain", &tileType{"foo", 10})
		viper.Set("grid", GridData{
			Size: 10,
			Types: map[string]string{
				"a": "plain",
			},
			Content: "AAA\nAAA\nAAA",
		})
		g := CreateGrid(func(t Tile) {})
		o := GetTerrain()
		if o != g {
			t.Errorf("GetTerrain should return created grid, got %+v", o)
		}
	})
}
