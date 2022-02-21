package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"ego/pkg/utils"
)

type State interface {
	Label() string
	Enter()
	Update(*StateMachine, terrain.Terrain) State
	Render(renderer.Renderer, renderable.Renderable)
}

func CreateState(name string, data ...interface{}) State {
	states := map[string]func() State{
		"idle": func() State {
			return &idleState{}
		},
		"explore": func() State {
			return &exploreState{}
		},
		"move": func() State {
			args := data[0].(struct {
				Position utils.Position
				Next     string
			})
			return &moveState{args.Position, args.Next}
		},
	}

	return states[name]()
}
