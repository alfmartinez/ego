package state

import (
	"ego/pkg/movement"
	"ego/pkg/terrain"
	"image"
	"testing"
)

type fakeMover struct {
	movement.Movement
}

func TestMoveState(t *testing.T) {
	s := CreateState("move", MoveArguments{
		Destination: movement.Loc(image.Pt(1, 1)),
		Next:        CreateState("idle"),
	})
	g := terrain.CreateGrid(2, 2)
	m := &fakeMover{movement.CreateMovement(image.Pt(0, 0))}

	if s.Label() != "moving" {
		t.Errorf("Label should be moving, got %+v", s.Label())
	}

	next := s.Update(m, g)
	if next != nil {
		t.Errorf("Update should return nil on move until done")
	}
	next = s.Update(m, g)
	if _, ok := next.(*idleState); !ok {
		t.Errorf("Should return next state on arrival")
	}

}
