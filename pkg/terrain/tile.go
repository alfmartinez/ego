package terrain

import (
	"ego/pkg/renderable"
	"ego/pkg/utils"
)

type Tile interface {
	renderable.Renderable
}

type tile struct {
	TileType
	position utils.Position
}

func CreateTile(position utils.Position, tileType TileType) Tile {
	return &tile{tileType, position}
}

func (t *tile) Position() utils.Position {
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
