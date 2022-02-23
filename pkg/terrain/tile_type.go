package terrain

type TileType interface {
}

type plainType struct {
}

func CreateTileType(name string) TileType {
	return &plainType{}
}
