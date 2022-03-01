package movement

import (
	"image"
)

type Positionnable interface {
	Position() image.Point
	IsAt(Positionnable) bool
}

type Moveable interface {
	MoveTowards(Positionnable) bool
}

type Movement interface {
	Positionnable
	Moveable
}

type Location struct {
	position image.Point
}

func Loc(pt image.Point) Location {
	return Location{pt}
}

func (m *Location) Position() image.Point {
	return m.position
}

func (m *Location) IsAt(pos Positionnable) bool {
	return m.position.Eq(pos.Position())
}

type movement struct {
	Location
}

func CreateMovement(coord image.Point) Movement {
	position := Location{coord}
	return &movement{position}
}

func (m *movement) MoveTowards(destination Positionnable) bool {
	v := destination.Position().Sub(m.Position())
	dp := image.Pt(0, 0)
	switch {
	case v.X > 0:
		dp.X = 1
	case v.X < 0:
		dp.X = -1
	}
	switch {
	case v.Y > 0:
		dp.Y = 1
	case v.Y < 0:
		dp.Y = -1
	}
	m.position = m.position.Add(dp)

	return dp.Eq(image.Point{})
}
