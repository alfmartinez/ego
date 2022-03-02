package terrain

import (
	"ego/pkg/movement"
	"ego/pkg/renderer"
)

type Terrain interface {
	GetTile(movement.Positionnable) Tile
	Render(renderer.Renderer)
	SearchAround(movement.Positionnable, int, func(Tile) bool) []movement.Positionnable
}
