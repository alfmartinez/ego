package terrain

import (
	"ego/engine/movement"
	"ego/engine/observer"
	"ego/engine/physics"
	"ego/engine/renderer"
	"image"
)

type Tile interface {
	observer.Observer
	movement.Positionnable
	Surrounding() []Tile
	AddSurrounding([]Tile)
	Rect() image.Rectangle
	Resources
	Path() string
	Size() uint
}

type tile struct {
	TileType
	Resources
	rect        image.Rectangle
	size        int
	surrounding []Tile
}

func CreateTile(coord GridCoord, tileType TileType, tileSize int) Tile {
	surrounding := make([]Tile, 0)
	rect := image.Rect(0, 0, tileSize, tileSize)
	rect = rect.Add(image.Point(coord).Mul(tileSize))
	res := CreateResources()
	return &tile{tileType, res, rect, tileSize, surrounding}
}

func (t *tile) OnNotify(e observer.Event) {
	switch e.Type() {
	case observer.RENDER:
		r := renderer.FromContext()
		r.Render(t)
	case observer.PHYSICS:
		p := physics.FromContext()
		p.Add(t)
	}

}

func (t *tile) Size() uint {
	return uint(t.size)
}

func (t *tile) IsAt(pos movement.Positionnable) bool {
	return pos.Position().In(t.rect)
}

func (t *tile) Position() movement.Location {
	return movement.Location{
		X: float64(t.rect.Min.X),
		Y: float64(t.rect.Min.Y),
	}
}

func (t *tile) Surrounding() []Tile {
	return t.surrounding
}

func (t *tile) Rect() image.Rectangle {
	return t.rect
}

func (t *tile) AddSurrounding(surrounding []Tile) {
	t.surrounding = append(t.surrounding, surrounding...)
}

func (t *tile) Hitbox() image.Rectangle {
	return t.rect
}
