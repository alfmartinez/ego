package movement

import (
	"ego/pkg/utils"
)

type Positionnable interface {
	Position() utils.Position
}

type Movement interface {
	Positionnable
	MoveTowards(utils.Position) bool
}

type movement struct {
	position utils.Position
}

func CreateMovement(position utils.Position) Movement {
	return &movement{position: position}
}

func (m *movement) Position() utils.Position {
	return m.position
}

func (m *movement) MoveTowards(destination utils.Position) bool {
	unit := m.position.UnitTowards(&destination)
	if unit.IsZero() {
		return true
	}

	newPosition := m.position.Add(unit)
	m.position = newPosition
	return false
}
