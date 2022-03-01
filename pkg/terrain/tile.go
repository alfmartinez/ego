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
	position    image.Point
	surrounding []Tile
}

func CreateTile(position image.Point, tileType TileType) Tile {
	surrounding := make([]Tile, 0)
	return &tile{tileType, position, surrounding}
}

func (t *tile) Position() image.Point {
	return t.position
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
	return 10
}

func (t *tile) Multiplicator() int {
	return 10
}
