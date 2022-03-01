package terrain

type TileType interface {
	Path() string
}

var types = make(map[string]TileType)

func CreateTileType(name string) TileType {
	tileType, ok := types[name]
	if !ok {
		tileType = &plainType{"mario.png"}
		types[name] = tileType
	}
	return tileType
}

type plainType struct {
	path string
}

func (t *plainType) Path() string {
	return t.path
}
