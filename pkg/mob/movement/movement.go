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
	return false
}
