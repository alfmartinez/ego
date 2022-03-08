package state

import "testing"

type fakeState struct{}

func (f fakeState) Update(s Updatable) State { return nil }
func TestCreateState(t *testing.T) {
	RegisterStateFactory("fake", func(i []interface{}) State {
		return fakeState{}
	})

	t.Run("CreateState uses registered factories", func(t *testing.T) {
		s := CreateState("fake")
		if _, ok := s.(fakeState); !ok {
			t.Errorf("should return fakeState, got %+v", s)
		}
	})
}
