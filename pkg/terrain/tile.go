package terrain

import (
	"ego/pkg/mob/movement"
	"ego/pkg/renderable"
	"image"
)

type Tile interface {
	movement.Positionnable
	renderable.Renderable
}

type tile struct {
	TileType
	position image.Point
}

func CreateTile(position image.Point, tileType TileType) Tile {
	return &tile{tileType, position}
}

func (t *tile) Position() image.Point {
	return t.position
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
