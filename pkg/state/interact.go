package state

func init() {
	RegisterStateFactory("interact", func(data []interface{}) State {
		return &interactState{}
	})
}

type interactState struct {
	baseState
}

func (s *interactState) Label() string {
	return "interacting"
}
