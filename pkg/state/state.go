package state

import (
	"ego/pkg/terrain"
)

type Actor interface {
}

type State interface {
	Enter()
	Update(m Actor, g terrain.Grid) State
}

func CreateState(name string) State {
	states := map[string]func() State{
		"idle": func() State {
			return &idleState{}
		},
		"explore": func() State {
			return &exploreState{}
		},
	}

	return states[name]()
}
