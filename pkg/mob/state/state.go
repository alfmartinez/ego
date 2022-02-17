package state

import (
	"ego/pkg/terrain"
)

type State interface {
	Enter()
	Update(m *StateMachine, g terrain.Grid) State
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
			return &moveState{}
		},
	}

	return states[name]()
}
