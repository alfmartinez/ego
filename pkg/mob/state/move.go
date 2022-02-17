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

func (s moveState) Enter() {
	log.Print("Entering Idle State")
}

func (s moveState) Update(a *StateMachine, g terrain.Terrain) State {
	return nil
}

func (s moveState) Render(r renderer.Renderer, m renderer.Renderable) {
	r.Render(m)
}
