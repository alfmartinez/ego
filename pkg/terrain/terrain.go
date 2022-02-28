package terrain

import (
	"ego/pkg/renderer"
	"image"
)

type Terrain interface {
	GetTile(image.Point) Tile
	Render(renderer.Renderer)
}
