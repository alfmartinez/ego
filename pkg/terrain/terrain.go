package terrain

import (
	"ego/pkg/movement"
	"image"
)

type Terrain interface {
	Tiles() map[image.Point]Tile
	Tile(image.Point) Tile
	GetTile(movement.Positionnable) Tile
	FindClosest(movement.Positionnable, int, func(Tile) bool) []Tile
	AddSource(int, int, Resource, uint)
}
