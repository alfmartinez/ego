package state

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
	return nil

}
