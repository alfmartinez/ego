package state

import (
	"ego/pkg/mob/memory"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type StateMachine struct {
	memory  *memory.Memory
	current State
	next    State
}

func CreateStateMachine(memory *memory.Memory) *StateMachine {
	return &StateMachine{memory: memory}
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

func (m *StateMachine) Memory() *memory.Memory {
	return m.memory
}
