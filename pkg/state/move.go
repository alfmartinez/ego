package state

import (
	"ego/pkg/movement"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("move", func(data []interface{}) State {
		args := data[0].(MoveArguments)
		return &moveState{
			destination: args.Destination,
			next:        args.Next,
		}
	})
}

type MoveArguments struct {
	Destination movement.Positionnable
	Next        State
}

type Mover interface {
	movement.Movement
}

type moveState struct {
	baseState
	destination movement.Positionnable
	next        State
}

func (s *moveState) Label() string {
	return "moving"
}

func (s *moveState) Update(sm interface{}, g terrain.Terrain) State {
	a := sm.(Mover)
	done := a.MoveTowards(s.destination)
	if done {
		return s.next
	}
	return nil
}
