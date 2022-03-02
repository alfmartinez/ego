package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type StateMachine interface {
	Render(interface{}, renderer.Renderer)
	Update(interface{}, terrain.Terrain)
	Doing() string
}

type stateMachine struct {
	current State
	next    State
}

func CreateStateMachine() StateMachine {
	return &stateMachine{}
}

func (m *stateMachine) Update(self interface{}, grid terrain.Terrain) {
	if m.current == nil {
		m.next = CreateState("idle")
	}
	if m.next != nil {
		m.current = m.next
		m.next = nil
		m.current.Enter()
	}
	m.next = m.current.Update(self, grid)
}

func (m *stateMachine) Render(self interface{}, r renderer.Renderer) {
	if m.current != nil {
		s := self.(renderable.Renderable)
		m.current.Render(r, s)
	}
}

func (m *stateMachine) Doing() string {
	var doing string = "nothing"
	if m.current != nil {
		doing = m.current.Label()
	}
	return doing
}
