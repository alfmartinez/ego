package terrain

import (
	"fmt"

	"github.com/spf13/viper"
)

type TileType interface {
	Path() string
}

type tileType struct {
	Sprite   string
	Movement int
}

type TileTypes map[string]tileType

var types = make(map[string]TileType)

func RegisterTileTypes() {
	var typesData TileTypes
	err := viper.UnmarshalKey("tile_types", &typesData)
	if err != nil {
		panic(err)
	}
	for name, tileType := range typesData {
		RegisterTileType(name, &tileType)
	}
}

func RegisterTileType(name string, tileType TileType) {
	types[name] = tileType
}

func CreateTileType(name string) TileType {
	tileType, ok := types[name]
	if !ok {
		panic(fmt.Errorf("unknown tile type %s", name))
	}
	return tileType
}

func (t *tileType) Path() string {
	return t.Sprite
}
