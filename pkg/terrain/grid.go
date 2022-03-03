package terrain

import (
	"ego/pkg/movement"
	"ego/pkg/renderer"
	"image"
)

type grid struct {
	tiles map[image.Point]Tile
}

const (
	tileSize = 100
)

func CreateGrid(width int, height int) Terrain {
	tiles := make(map[image.Point]Tile)
	tileType := CreateTileType("plain")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := image.Pt(x, y)
			tiles[position] = CreateTile(position, tileType, tileSize)
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
		t.AddSurrounding(surrounding)
	}

	return &grid{tiles}

}

func (g *grid) Tile(pt image.Point) Tile {
	return g.tiles[pt]
}

func (g *grid) GetTile(pos movement.Positionnable) Tile {
	return g.tiles[pos.Position().Div(tileSize)]
}

func (g *grid) Render(r renderer.Renderer) {
	for _, tile := range g.tiles {
		r.Render(tile)
	}
}

func (g *grid) SearchAround(a movement.Positionnable, r int, validate func(t Tile) bool) []movement.Positionnable {
	tile := g.GetTile(a)
	current := tile.Surrounding()
	found := make(map[movement.Positionnable]movement.Positionnable)
	if validate(tile) {
		pos := movement.Loc(tile.Position().Div(tileSize))
		found[pos] = pos
	}
	for r > 0 {
		var next []Tile
		for _, t := range current {
			_, ok := found[t]
			if validate(t) && !ok {
				found[t] = t
				next = append(next, t.Surrounding()...)
			}
		}
		r--
		current = next
	}

	results := make([]movement.Positionnable, 0, len(found))
	for _, f := range found {
		results = append(results, f)
	}

	return results
}

func (g *grid) AddSource(x int, y int, kind string, quantity uint) {
	tile := g.tiles[image.Pt(x, y)]
	tile.AddResource(kind, quantity)
}

func (g *grid) FindClosest(a movement.Positionnable, validate func(Tile) bool) Tile {

	var added = make(map[Tile]bool)
	var order = make([]Tile, 0)

	var expand func(tile Tile)

	expand = func() func(Tile) {
		return func(tile Tile) {
			added[tile] = true
			order = append(order, tile)
			for _, x := range tile.Surrounding() {
				if _, ok := added[x]; !ok {
					expand(x)
				}
			}
		}
	}()

	expand(g.GetTile(a))

	for _, x := range order {
		if validate(x) {
			return x
		}
	}

	return nil
}
