package state

import (
	"ego/pkg/mob/movement"
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
				Destination movement.Positionnable
				Next        string
			})
			return &moveState{args.Destination, args.Next}
		},
		"heal": func() State {
			return &healState{}
		},
		"rest": func() State {
			return &restState{}
		},
		"interact": func() State {
			return &interactState{}
		},
	}

	return states[name]()
}
