package state

func init() {
	RegisterStateFactory("rest", func(data []interface{}) State {
		return &restState{}
	})
}

type restState struct{}

func (s *restState) Update(sm Updatable) State {
	return nil
}
