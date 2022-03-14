package movement

import (
	"image"
)

type Positionnable interface {
	Position() image.Point
	IsAt(Positionnable) bool
}

type Moveable interface {
	MoveForward(Positionnable) bool
}

type Movement interface {
	Positionnable
	Moveable
}

type movement struct {
	location
}

func CreateMovement(coord image.Point) Movement {
	position := location{coord}
	return &movement{position}
}

func (m *movement) MoveForward(destination Positionnable) bool {
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
