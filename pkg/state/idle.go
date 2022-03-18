package state

import (
	"ego/pkg/input"
	"ego/pkg/movement"
)

func init() {
	RegisterStateFactory(StateIdle, func(data []interface{}) State {
		return &idleState{StateIdle}
	})
}

type idleState struct {
	StateType
}

func (s *idleState) Update(a Updatable) State {
	a.Frame(0, 0)

	inputHandler := input.FromContext()
	if inputHandler.IsPressed(input.JUMP) {
		return CreateState(StateMove, movement.MOVE_NONE)
	}
	if inputHandler.IsPressed(input.RIGHT) {
		return CreateState(StateMove, movement.MOVE_RIGHT)
	}

	return nil

}
