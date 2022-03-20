package state

import (
	"testing"
	"time"
)

type fakeState struct {
	StateType
}

func (f fakeState) Update(Updatable, time.Duration) State { return nil }
func TestCreateState(t *testing.T) {
	RegisterStateFactory(255, func(i []interface{}) State {
		return fakeState{255}
	})

	t.Run("CreateState uses registered factories", func(t *testing.T) {
		s := CreateState(255)
		if _, ok := s.(fakeState); !ok {
			t.Errorf("should return fakeState, got %+v", s)
		}
	})

}

func TestStateType(t *testing.T) {
	o := StateIdle
	if o.Type() != StateIdle {
		t.Error("Type should return object proper")
	}
}
