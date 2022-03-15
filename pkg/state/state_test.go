package state

import "testing"

type fakeState struct {
	StateType
}

func (f fakeState) Update(s Updatable) State { return nil }
func TestCreateState(t *testing.T) {
	RegisterStateFactory(255, func(i []interface{}) State {
		return fakeState{}
	})

	t.Run("CreateState uses registered factories", func(t *testing.T) {
		s := CreateState(255)
		if _, ok := s.(fakeState); !ok {
			t.Errorf("should return fakeState, got %+v", s)
		}
	})
}
