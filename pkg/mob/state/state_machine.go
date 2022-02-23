package state

import (
	"ego/pkg/mob/data"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/motivator"
	"ego/pkg/mob/movement"
	"ego/pkg/renderer"
	"ego/pkg/sprite"
	"ego/pkg/terrain"
)

type StateMachine interface {
	memory.Memory
	data.Data
	movement.Movement
	motivator.NeedsCollection
	Update(terrain.Terrain)
	Render(renderer.Renderer)
	Doing() string
}

type stateMachine struct {
	memory.Memory
	data.Data
	movement.Movement
	sprite.Sprite
	motivator.NeedsCollection
	current State
	next    State
}

func CreateStateMachine(memory memory.Memory, data data.Data, movement movement.Movement, sprite sprite.Sprite, needs motivator.NeedsCollection) StateMachine {
	return &stateMachine{
		memory,
		data,
		movement,
		sprite,
		needs,
		nil,
		nil,
	}
}

func (m *stateMachine) Update(grid terrain.Terrain) {
	m.UpdateNeeds()
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

func (m *stateMachine) Render(r renderer.Renderer) {
	if m.current != nil {
		m.current.Render(r, m)
	}
}

func (m *stateMachine) Doing() string {
	return m.current.Label()
}
