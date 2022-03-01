package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type State interface {
	Label() string
	Enter()
	Update(StateMachine, terrain.Terrain) State
	Render(renderer.Renderer, renderable.Renderable)
}

var states = make(map[string]func([]interface{}) State)

func RegisterStateFactory(name string, factory func([]interface{}) State) {
	states[name] = factory
}

func CreateState(name string, data ...interface{}) State {
	return states[name](data)
}
