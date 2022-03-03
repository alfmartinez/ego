package terrain

import (
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"image"
)

type Tile interface {
	movement.Positionnable
	renderable.Renderable
	Surrounding() []Tile
	AddSurrounding([]Tile)
	Resources
}

type tile struct {
	TileType
	Resources
	rect        image.Rectangle
	size        int
	surrounding []Tile
}

func CreateTile(coord image.Point, tileType TileType, tileSize int) Tile {
	surrounding := make([]Tile, 0)
	rect := image.Rect(0, 0, tileSize, tileSize)
	rect = rect.Add(coord.Mul(tileSize))
	res := CreateResources()
	return &tile{tileType, res, rect, tileSize, surrounding}
}

func (t *tile) IsAt(pos movement.Positionnable) bool {
	return pos.Position().In(t.rect)
}

func (t *tile) Position() image.Point {
	return t.rect.Min
}

func (t *tile) Surrounding() []Tile {
	return t.surrounding
}

func (t *tile) AddSurrounding(surrounding []Tile) {
	t.surrounding = append(t.surrounding, surrounding...)
}

func (t *tile) Doing() string {
	return ""
}

func (t *tile) Name() string {
	return ""
}

func (t *tile) Size() uint {
	return uint(t.size)
}
