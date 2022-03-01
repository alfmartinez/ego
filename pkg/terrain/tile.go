package terrain

import (
	"ego/pkg/mob/movement"
	"ego/pkg/renderable"
	"image"
)

type Tile interface {
	movement.Positionnable
	renderable.Renderable
	Surrounding() []Tile
	AddSurrounding([]Tile)
}

type tile struct {
	TileType
	movement.Location
	surrounding []Tile
}

func CreateTile(coord image.Point, tileType TileType) Tile {
	surrounding := make([]Tile, 0)
	return &tile{tileType, movement.Loc(coord), surrounding}
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

func (t *tile) Path() string {
	return "mario.png"
}

func (t *tile) Size() uint {
	return 100
}

func (t *tile) Multiplicator() int {
	return 100
}
