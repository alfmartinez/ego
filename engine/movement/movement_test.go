package movement

import (
	"fmt"
	"image"
	"testing"
)

func TestCreateMovementReturnsMovement(t *testing.T) {
	i := CreateMovement(image.Pt(1, 10))

	if i.Position().X != 1 {
		t.Errorf("should be initialized with X value")
	}

	if i.Position().Y != 10 {
		t.Errorf("should be initialized with Y value")
	}
}

func TestMoveTowardsReturnsTrueIfDestinationReached(t *testing.T) {
	m := CreateMovement(image.Pt(10, 10))
	d := Loc(image.Pt(10, 10))

	reached := m.MoveForward(d)
	if !reached {
		t.Errorf("should have reached destination, is at %+v", m.Position())
	}
}

func TestMoveTowardsReturnsFalseIfDestinationNotReached(t *testing.T) {
	cases := []struct{ orig, destination, expected image.Point }{
		{image.Pt(10, 10), image.Pt(1, 10), image.Pt(9, 10)},
		{image.Pt(10, 10), image.Pt(10, 1), image.Pt(10, 9)},
		{image.Pt(10, 10), image.Pt(10, 20), image.Pt(10, 11)},
		{image.Pt(10, 10), image.Pt(20, 10), image.Pt(11, 10)},
		{image.Pt(10, 10), image.Pt(1, 1), image.Pt(9, 9)},
		{image.Pt(10, 10), image.Pt(20, 20), image.Pt(11, 11)},
		{image.Pt(10, 10), image.Pt(1, 20), image.Pt(9, 11)},
		{image.Pt(10, 10), image.Pt(20, 1), image.Pt(11, 9)},
	}

	for _, tt := range cases {
		testName := fmt.Sprintf("Move from %+s towards %+s", tt.orig, tt.destination)
		t.Run(testName, func(t *testing.T) {
			m := CreateMovement(tt.orig)
			d := Loc(tt.destination)

			reached := m.MoveForward(d)

			if reached {
				t.Errorf("should not have reached destination, is at %+v", m.Position())
			}
			if !m.Position().Eq(tt.expected) {
				t.Errorf("should be at %s, is at %s instead", m.Position(), tt.expected)
			}
		})
	}

}
