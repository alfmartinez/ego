package terrain

import (
	"ego/pkg/context"
	"ego/pkg/movement"
	"image"
	"strings"

	"github.com/spf13/viper"
)

type GridCoord image.Point

func (c GridCoord) Add(p image.Point) GridCoord {
	return GridCoord{
		X: c.X + p.X,
		Y: c.Y + p.Y,
	}
}

type grid struct {
	tiles map[GridCoord]Tile
	size  int
}

type GridData struct {
	Size    int
	Content string
	Types   map[string]string
}

func CreateGrid(register func(Tile)) Terrain {
	var gridData GridData
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	err := viper.UnmarshalKey("grid", &gridData)
	if err != nil {
		panic(err)
	}

	tiles := make(map[GridCoord]Tile)

	lines := strings.Split(gridData.Content, "\n")
	for y, line := range lines {
		for x, vType := range line {
			typeName := strings.ToLower(string(vType))
			tileType := CreateTileType(gridData.Types[typeName])
			position := GridCoord(image.Pt(x, y))
			tiles[position] = CreateTile(position, tileType, gridData.Size)
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

	myGrid := &grid{tiles, gridData.Size}

	return myGrid

}

func (g *grid) Tiles() map[GridCoord]Tile {
	return g.tiles
}

func (g *grid) Tile(pt GridCoord) Tile {
	return g.tiles[pt]
}

func (g *grid) GetTile(pos movement.Positionnable) Tile {
	coords := g.LocationToGridCoord(pos.Position())
	return g.tiles[coords]
}

func (g *grid) AddSource(x int, y int, kind Resource, quantity uint) {
	tile := g.tiles[GridCoord{x, y}]
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

func (g *grid) LocationToGridCoord(l movement.Location) GridCoord {
	return GridCoord{
		X: int(l.X / float64(g.size)),
		Y: int(l.Y / float64(g.size)),
	}
}
