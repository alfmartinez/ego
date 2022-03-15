package spritesheet

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if _, ok := s.(*spritesheet); !ok {
		t.Error("New should return *spritesheet")
	}
}
