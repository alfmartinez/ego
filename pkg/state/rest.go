package state

func init() {
	RegisterStateFactory("rest", func(data []interface{}) State {
		return &restState{}
	})
}

type restState struct {
	baseState
}

func (s *restState) Label() string {
	return "resting"
}
