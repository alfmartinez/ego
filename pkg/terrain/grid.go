package terrain

import (
	"ego/pkg/renderer"
	"ego/pkg/utils"
)

type grid struct {
	tiles map[utils.Position]Tile
}

func CreateGrid(width int, height int) Terrain {
	tiles := make(map[utils.Position]Tile)
	tileType := CreateTileType("plain")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := utils.Position{X: x, Y: y}
			tiles[position] = CreateTile(position, tileType)
		}
	}
	return &grid{tiles}

}

func (g *grid) GetTile(position utils.Position) Tile {
	return g.tiles[position]
}

func (g *grid) Render(r renderer.Renderer) {
	for _, tile := range g.tiles {
		r.Render(tile)
	}
}
