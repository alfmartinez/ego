package movement

import (
	"image"
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
	Acceleration() Acceleration
	SetAcceleration(Acceleration)
	Position() Location
	SetPosition(Location)
	Speed() Speed
	SetSpeed(Speed)
	MoveForward(Positionnable) bool
	MoveDirection(Direction, time.Duration)
	Land()
	Fall()
}

type Movement interface {
	Positionnable
	Movable
}

type movement struct {
	location     Location
	speed        Speed
	acceleration Acceleration
	onGround     bool
}

func CreateMovement(coord Location) Movement {
	position := Location(coord)
	return &movement{
		location:     position,
		speed:        Speed{0, 0},
		acceleration: Acceleration{0, 0},
	}
}

func (m *movement) Land() {
	m.acceleration.Y = 0
}

func (m *movement) Fall() {
	m.acceleration.Y = 10
}

func (m *movement) SetPosition(l Location) {
	m.location = l
}

func (m *movement) Position() Location {
	return m.location
}

func (m *movement) SetSpeed(s Speed) {
	m.speed = s
}

func (m *movement) Speed() Speed {
	return m.speed
}

func (m *movement) SetAcceleration(a Acceleration) {
	m.acceleration = a
}

func (m *movement) Acceleration() Acceleration {
	return m.acceleration
}

func (m *movement) MoveForward(destination Positionnable) bool {
	v := destination.Position().Sub(m.location)
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
	m.location = m.location.Add(dp)

	return dp.Eq(image.Point{})
}

func (m *movement) MoveDirection(d Direction, dt time.Duration) {
	var a Acceleration
	switch d {
	case MOVE_NONE:
		a = Acceleration{0, 0}
	case MOVE_DOWN:
		a = Acceleration{0, ACCELERATION_UNIT}
	case MOVE_UP:
		a = Acceleration{0, -ACCELERATION_UNIT}
	case MOVE_LEFT:
		a = Acceleration{-ACCELERATION_UNIT, 0}
	case MOVE_RIGHT:
		a = Acceleration{ACCELERATION_UNIT, 0}
	}
	m.acceleration = a
}

func (m *movement) IsAt(p Positionnable) bool {
	mPt := m.location.Point()
	return mPt.Eq(p.Position().Point())
}
