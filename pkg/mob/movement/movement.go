package movement

import (
	"image"
)

type Positionnable interface {
	Position() image.Point
}

type Movement interface {
	Positionnable
	MoveTowards(Positionnable) bool
}

type movement struct {
	position image.Point
}

func CreateMovement(position image.Point) Movement {
	return &movement{position: position}
}

func (m *movement) Position() image.Point {
	return m.position
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
