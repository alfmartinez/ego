package state

import (
	"ego/pkg/terrain"
)

type StateMachine struct {
	current State
	next    State
}

func CreateStateMachine() StateMachine {
	return StateMachine{}
}

func (m *StateMachine) Update(grid terrain.Grid) {
	if m.next != nil {
		m.current = m.next
		m.next = nil
		m.current.Enter()
	}
	if m.current == nil {
		m.next = CreateState("idle")
	} else {
		m.next = m.current.Update(m, grid)
	}
}
