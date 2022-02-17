package movement

import (
	"ego/pkg/utils"
)

type Movement struct {
	position utils.Position
}

func CreateMovement(position utils.Position) *Movement {
	return &Movement{position: position}
}

func (m *Movement) Position() utils.Position {
	return m.position
}

func (m *Movement) MoveTowards(destination utils.Position) bool {
	unit := m.position.UnitTowards(&destination)
	if unit.IsZero() {
		return true
	}

	newPosition := m.position.Add(unit)
	m.position = newPosition
	return false
}
