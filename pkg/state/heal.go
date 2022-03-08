package state

import (
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("heal", func(data []interface{}) State {
		return &healState{}
	})
}

type Patient interface {
	movement.Positionnable
	Provide(motivator.Need, int, int)
}

type healState struct {
	target terrain.Tile
}

func (s *healState) Update(a Updatable) State {
	return nil
}
