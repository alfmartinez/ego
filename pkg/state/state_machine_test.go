package state

import "testing"

type fakeUpdatable struct {
	Need string
}

func (f fakeUpdatable) TopNeed() string {
	return f.Need
}

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
		self.Need = "curiosity"
		sm.Update(self)
	})
}
