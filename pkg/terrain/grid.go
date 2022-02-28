package terrain

import (
	"ego/pkg/renderer"
	"image"
)

type grid struct {
	tiles map[image.Point]Tile
}

func CreateGrid(width int, height int) Terrain {
	tiles := make(map[image.Point]Tile)
	tileType := CreateTileType("plain")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := image.Pt(x, y)
			tiles[position] = CreateTile(position, tileType)
		}
	}
	return &grid{tiles}

}

func (g *grid) GetTile(position image.Point) Tile {
	return g.tiles[position]
}

func (g *grid) Render(r renderer.Renderer) {
	for _, tile := range g.tiles {
		r.Render(tile)
	}
}
