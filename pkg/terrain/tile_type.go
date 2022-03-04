package terrain

type TileType interface {
	Path() string
}

var types = make(map[string]TileType)

func CreateTileType(name string) TileType {
	tileType, ok := types[name]
	if !ok {
		tileType = &defaultType{"mario.png"}
		types[name] = tileType
	}
	return tileType
}

type defaultType struct {
	path string
}

func (t *defaultType) Path() string {
	return t.path
}
