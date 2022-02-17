package state

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"ego/pkg/utils"
	"log"
)

type State interface {
	Enter()
	Update(*StateMachine, terrain.Terrain) State
	Render(renderer.Renderer, renderer.Renderable)
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
			log.Printf("%+v", args)
			return &moveState{args.Position, args.Next}
		},
	}

	return states[name]()
}
