package movement

import (
	"ego/engine/physics/matrix"
	"time"
)

type Direction uint

const (
	ACCELERATION_UNIT = 50
)

const (
	MOVE_NONE Direction = iota
	MOVE_RIGHT
	MOVE_LEFT
	MOVE_UP
	MOVE_DOWN
)

type Positionnable interface {
	Position() Location
	IsAt(Positionnable) bool
}

type Movable interface {
	Position() Location
	SetMatrix(matrix.M)
	GetMatrix() matrix.M
}

type Movement interface {
	Positionnable
	Movable
	MoveDirection(Direction, time.Duration)
}

type movement struct {
	matrix matrix.M
}

func CreateMovement(coord Location) Movement {
	m := matrix.M{
		A: matrix.Vec2{X: 0, Y: 0},
		S: matrix.Vec2{X: 0, Y: 0},
		P: matrix.Vec2{X: coord.X, Y: coord.Y},
	}
	return &movement{m}
}

func (m *movement) Position() Location {
	return Location(m.matrix.P)
}

func (m *movement) IsAt(p Positionnable) bool {
	return m.Position().Point().Eq(p.Position().Point())
}

func (m *movement) GetMatrix() matrix.M {
	return m.matrix
}

func (m *movement) SetMatrix(o matrix.M) {
	m.matrix = o
}

func (m *movement) MoveDirection(direction Direction, dt time.Duration) {
	speed := matrix.Vec2{}
	switch direction {
	case MOVE_LEFT:
		speed.X = -10
	case MOVE_RIGHT:
		speed.X = 10
	}
	matrix := m.matrix
	matrix.S.Add(speed)
	m.matrix = matrix
}
