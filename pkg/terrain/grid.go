package terrain

import (
	"ego/pkg/movement"
	"image"
)

type grid struct {
	tiles map[image.Point]Tile
}

const (
	tileSize = 100
)

func CreateGrid(width int, height int, register func(Tile)) Terrain {
	tiles := make(map[image.Point]Tile)
	tileType := CreateTileType("plain")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := image.Pt(x, y)
			tiles[position] = CreateTile(position, tileType, tileSize)
			if x == 5 && y == 5 {
				tiles[position].AddResource(Health, 5)
			}
			if x == 8 && y == 0 {
				tiles[position].AddResource(Health, 5)
			}
		}
	}

	surrounds := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for pos, t := range tiles {
		surrounding := make([]Tile, 0)
		for _, around := range surrounds {
			dP := image.Pt(around.dx, around.dy)
			pos := pos.Add(dP)
			if tile, ok := tiles[pos]; ok {
				surrounding = append(surrounding, tile)
			}
		}
		register(t)
		t.AddSurrounding(surrounding)
	}

	myGrid := &grid{tiles}

	terrainSingleton = myGrid
	return myGrid

}

func (g *grid) Tile(pt image.Point) Tile {
	return g.tiles[pt]
}

func (g *grid) GetTile(pos movement.Positionnable) Tile {
	return g.tiles[pos.Position().Div(tileSize)]
}

func (g *grid) AddSource(x int, y int, kind Resource, quantity uint) {
	tile := g.tiles[image.Pt(x, y)]
	tile.AddResource(kind, quantity)
}

func (g *grid) FindClosest(a movement.Positionnable, count int, validate func(Tile) bool) []Tile {
	var found = make([]Tile, 0, count)
	var added = make(map[Tile]bool)
	var order = make([]Tile, 0)

	var expand func(tile Tile)

	expand = func(tile Tile) {
		added[tile] = true
		order = append(order, tile)
		for _, x := range tile.Surrounding() {
			if _, ok := added[x]; !ok {
				expand(x)
			}
		}
	}

	expand(g.GetTile(a))

	for _, x := range order {
		if validate(x) {
			found = append(found, x)
		}
	}

	return found
}
