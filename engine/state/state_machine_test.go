package state

import (
	"testing"
	"time"
)

type fakeUpdatable struct{}

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
		sm.SetStates(States{
			"default": func(dt time.Duration) string {
				return ""
			},
		})
		self := fakeUpdatable{}
		sm.DoUpdate(self, time.Second)
	})
}
