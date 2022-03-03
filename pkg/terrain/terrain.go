package terrain

import (
	"ego/pkg/movement"
	"ego/pkg/renderer"
	"image"
)

type Terrain interface {
	Tile(image.Point) Tile
	GetTile(movement.Positionnable) Tile
	Render(renderer.Renderer)
	SearchAround(movement.Positionnable, int, func(Tile) bool) []movement.Positionnable
	FindClosest(movement.Positionnable, func(Tile) bool) Tile
	AddSource(int, int, string, uint)
}
