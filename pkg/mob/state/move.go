package state

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"ego/pkg/utils"
	"log"
)

type moveState struct {
	position utils.Position
	next     string
}

func (s moveState) Label() string {
	return "move"
}

func (s moveState) Enter() {
	log.Print("Entering Moving State")
}

func (s moveState) Update(a *StateMachine, g terrain.Terrain) State {
	done := a.MoveTowards(s.position)
	if done {
		return CreateState(s.next)
	}
	return nil
}

func (s moveState) Render(r renderer.Renderer, m renderer.Renderable) {
	r.Render(m)
}
