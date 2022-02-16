package terrain

import (
	"ego/pkg/utils"
)

type Grid struct {
	tiles map[utils.Position]Tile
}

func CreateGrid(width int, height int) *Grid {
	tiles := make(map[utils.Position]Tile)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := utils.Position{x, y}
			tiles[position] = Tile{}
		}
	}
	return &Grid{tiles}

}

func (g *Grid) GetTile(position utils.Position) Tile {
	return g.tiles[position]
}
