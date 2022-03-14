package state

import (
	"ego/pkg/motivator"
	"testing"
)

type fakeUpdatable struct {
	N motivator.Need
}

func (f fakeUpdatable) TopNeed() motivator.Need {
	return f.N
}

func (f fakeUpdatable) Frame(x, y int) {}

func TestStateMachine(t *testing.T) {
	t.Run("CreateStateMachine", func(t *testing.T) {
		sm := CreateStateMachine()
		if _, ok := sm.(*stateMachine); !ok {
			t.Errorf("should return stateMachine, got %+v", sm)
		}
	})

	t.Run("Update with no current state", func(t *testing.T) {
		sm := CreateStateMachine()
		self := fakeUpdatable{}
		self.N = motivator.Learn
		sm.DoUpdate(self)
	})

	t.Run("SetState overrides current state", func(t *testing.T) {
		sm := CreateStateMachine()
		sm.SetState(StateMove)
	})
}
