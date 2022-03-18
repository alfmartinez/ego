package terrain

import (
	"ego/pkg/context"
	"fmt"

	"github.com/spf13/viper"
)

type TileType interface {
	Path() string
	MovementCost() int
	IsIgnored() bool
}

type tileType struct {
	Sprite   string
	Movement int
	Blocking bool
}

type TileTypes map[string]tileType

var tileTypes = make(map[string]TileType)

func RegisterTileTypes() {
	var typesData TileTypes
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	err := viper.UnmarshalKey("tile_types", &typesData)
	if err != nil {
		panic(err)
	}
	for name, _ := range typesData {
		value := typesData[name]
		RegisterTileType(name, &value)
	}
}

func RegisterTileType(name string, value TileType) {
	tileTypes[name] = value
}

func CreateTileType(name string) TileType {
	value, ok := tileTypes[name]
	if !ok {
		panic(fmt.Errorf("unknown tile type %s", name))
	}
	return value
}

func (t *tileType) Path() string {
	return t.Sprite
}

func (t *tileType) MovementCost() int {
	return t.Movement
}

func (t *tileType) IsIgnored() bool {
	return !t.Blocking
}
