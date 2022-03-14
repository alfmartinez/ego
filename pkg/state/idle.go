package state

func init() {
	RegisterStateFactory(StateIdle, func(data []interface{}) State {
		return &idleState{}
	})
}

type idleState struct{}

func (s *idleState) Update(a Updatable) State {
	a.Frame(0, 0)
	return nil

}
