package state

import "ego/pkg/input"

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
		return CreateState(StateMove)
	}

	return nil

}
