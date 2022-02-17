package movement

import (
	"ego/pkg/utils"
	"log"
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
	log.Printf("%+v", unit)
	if unit.IsZero() {
		return true
	}

	newPosition := m.position.Add(unit)
	log.Printf("%+v", newPosition)
	m.position = newPosition
	return false
}
