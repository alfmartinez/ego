package terrain

import (
	"ego/pkg/context"
	"ego/pkg/movement"
)

func FromContext() Terrain {
	return context.GetContext().Get("terrain").(Terrain)
}

type Terrain interface {
	Tiles() map[GridCoord]Tile
	Tile(GridCoord) Tile
	GetTile(movement.Positionnable) Tile
	FindClosest(movement.Positionnable, int, func(Tile) bool) []Tile
	AddSource(int, int, Resource, uint)
}
