package state

import "testing"

var mockFrameMethod func(int, int)

type mockUpdatable struct{}

func (m *mockUpdatable) Frame(x, y int) { mockFrameMethod(x, y) }

func TestIdleState(t *testing.T) {
	var actualX, actualY = -1, -1
	s := CreateState(StateIdle)
	mockFrameMethod = func(x, y int) {
		actualX = x
		actualY = y
	}
	s.Update(&mockUpdatable{})
	if actualX != 0 || actualY != 0 {
		t.Error("Update should set frame to 0,0")
	}

}
