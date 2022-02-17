package state

import (
	"ego/pkg/mob/data"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/movement"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type StateMachine struct {
	*memory.Memory
	*data.Data
	*movement.Movement
	current State
	next    State
}

func CreateStateMachine(memory *memory.Memory, data *data.Data, movement *movement.Movement) *StateMachine {
	return &StateMachine{
		memory,
		data,
		movement,
		nil,
		nil,
	}
}

func (m *StateMachine) Update(grid terrain.Terrain) {
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

func (m *StateMachine) Render(r renderer.Renderer) {
	if m.current != nil {
		m.current.Render(r, m)
	}
}

func (m *StateMachine) Doing() string {
	return m.current.Label()
}
