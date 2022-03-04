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
	FindClosest(movement.Positionnable, int, func(Tile) bool) []Tile
	AddSource(int, int, string, uint)
}
