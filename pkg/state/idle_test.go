package state

import (
	"ego/pkg/motivator"
	"fmt"
	"testing"
)

func TestIdleState(t *testing.T) {
	s := CreateState("idle")
	sm := fakeUpdatable{}

	cases := []struct {
		name     string
		need     motivator.Need
		expected string
	}{
		{
			"none",
			motivator.None,
			"<nil>",
		}, {
			"learn",
			motivator.Learn,
			"*state.exploreState",
		}, {
			"recreation",
			motivator.Recreation,
			"*state.interactState",
		}, {
			"rest",
			motivator.Rest,
			"*state.restState",
		}, {
			"health",
			motivator.Health,
			"*state.healState",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sm.N = tt.need
			n := s.Update(sm)
			typeOf := fmt.Sprintf("%T", n)
			if typeOf != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, typeOf)
			}
		})
	}
}
