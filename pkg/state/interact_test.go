package state

import "testing"

func TestInte(t *testing.T) {
	s := CreateState("interact")

	if s.Label() != "interacting" {
		t.Errorf("Interact State should interact")
	}
}
