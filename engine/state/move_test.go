package state

import "testing"

func TestMoveState(t *testing.T) {
	s := CreateState(StateMove)
	mockFrameMethod = func(x, y int) {
		if x != 0 || y != 0 {
			t.Error("Should be 0,0")
		}
	}
	s.Update(&mockUpdatable{})
	mockFrameMethod = func(x, y int) {
		if x != 1 || y != 0 {
			t.Error("Should be 1,0")
		}
	}
	s.Update(&mockUpdatable{})
}
