package state

func init() {
	RegisterStateFactory("interact", func(data []interface{}) State {
		return &interactState{}
	})
}

type interactState struct{}

func (s *interactState) Update(sm Updatable) State {
	return nil
}
