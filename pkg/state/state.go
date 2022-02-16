package state

import (
	"ego/pkg/terrain"
	"ego/pkg/utils"
)

type Actor interface {
	Position() utils.Position
	HasFullyExplored(utils.Position) bool
	FindTileToExplore(g terrain.Grid) *terrain.Tile
	ExecuteCommand(command string, options ...interface{})
}

type State interface {
	Enter()
	Update(m Actor, g terrain.Grid) State
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
