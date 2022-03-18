package movement

import (
	"image"
)

type Direction uint

const (
	MOVE_NONE Direction = iota
	MOVE_RIGHT
	MOVE_LEFT
	MOVE_UP
	MOVE_DOWN
)

type Positionnable interface {
	Position() image.Point
	IsAt(Positionnable) bool
}

type Moveable interface {
	MoveForward(Positionnable) bool
	MoveDirection(Direction)
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

func (m *movement) MoveDirection(d Direction) {
	var dp image.Point
	switch d {
	case MOVE_DOWN:
		dp = image.Pt(0, 1)
	case MOVE_UP:
		dp = image.Pt(0, -1)
	case MOVE_LEFT:
		dp = image.Pt(-1, 0)
	case MOVE_RIGHT:
		dp = image.Pt(1, 0)
	}

	m.position = m.position.Add(dp)
}
