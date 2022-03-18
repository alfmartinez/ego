package state

func init() {
	RegisterStateFactory(StateMove, func(data []interface{}) State {
		return &moveState{StateMove, 0}
	})
}

type moveState struct {
	StateType
	frame int
}

func (s *moveState) Update(a Updatable) State {
	a.Frame(s.frame, 0)
	s.frame = (s.frame + 1) % 20
	if s.frame == 0 {
		return CreateState(StateIdle)
	}
	return nil

}
