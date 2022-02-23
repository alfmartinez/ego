package terrain

import (
	"ego/pkg/renderer"
	"ego/pkg/utils"
)

type Terrain interface {
	GetTile(utils.Position) Tile
	Render(renderer.Renderer)
}
