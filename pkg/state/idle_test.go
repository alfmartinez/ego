package state

import (
	"ego/pkg/motivator"
	"ego/pkg/terrain"
	"reflect"
	"testing"
)

type fakeIdler struct {
	motivator.NeedsCollection
}

func createFakeIdler(need string) Idler {
	nc := motivator.CreateNeedsCollection()
	nc.AddNeed(motivator.CreateNeed(need, 0), 10)
	return &fakeIdler{nc}
}

func TestIdleState(t *testing.T) {
	s := CreateState("idle")
	g := terrain.CreateGrid(1, 1)

	cases := []struct {
		need              string
		expectedInterface string
	}{
		{"curiosity", "*state.exploreState"},
		{"social", "*state.interactState"},
		{"rest", "*state.restState"},
		{"health", "*state.healState"},
	}

	for _, x := range cases {
		sm := createFakeIdler(x.need)
		next := s.Update(sm, g)

		if reflect.TypeOf(next).String() != x.expectedInterface {
			t.Errorf("Expected state %+v, got %+v", x.expectedInterface, reflect.TypeOf(next).String())
		}

	}

	sm := createFakeIdler("other")
	next := s.Update(sm, g)
	if next != nil {
		t.Errorf("Other needs should return nil, got %+v", next)
	}

}
