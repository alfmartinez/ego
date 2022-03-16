package terrain

import "fmt"

type TileType interface {
	Path() string
}

func init() {
	types["plain"] = &defaultType{"sheet:0:0"}
	types["swamp"] = &defaultType{"sheet:1:0"}
	types["mountain"] = &defaultType{"sheet:2:0"}
}

var types = make(map[string]TileType)

func CreateTileType(name string) TileType {
	tileType, ok := types[name]
	if !ok {
		panic(fmt.Errorf("unknown tile type %s", name))
	}
	return tileType
}

type defaultType struct {
	path string
}

func (t *defaultType) Path() string {
	return t.path
}
