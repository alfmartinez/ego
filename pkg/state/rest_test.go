package state

import "testing"

func TestRest(t *testing.T) {
	s := CreateState("rest")

	if s.Label() != "resting" {
		t.Errorf("Rest state should rest")
	}
}
