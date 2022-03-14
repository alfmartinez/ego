package state

func init() {
	RegisterStateFactory(StateMove, func(data []interface{}) State {
		return &moveState{0}
	})
}

type moveState struct {
	frame int
}

func (s *moveState) Update(a Updatable) State {
	a.Frame(s.frame, 0)
	s.frame = (s.frame + 1) % 20
	return nil

}
