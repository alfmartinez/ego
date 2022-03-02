package state

import (
	"ego/pkg/configuration"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"image"
	"testing"
)

type FakeMob interface {
	Idler
}
type fakeMob struct {
	motivator.NeedsCollection
	StateMachine
	movement.Positionnable
}

func (f *fakeMob) Name() string {
	return "fake"
}

func (f *fakeMob) Path() string {
	return "fake"
}

func (f *fakeMob) Size() uint {
	return 1
}

func CreateFakeMob(sm StateMachine) FakeMob {
	needs := motivator.CreateNeedsCollection()
	return &fakeMob{needs, sm, movement.Loc(image.Pt(0, 0))}
}

func TestCreateStateMachine(t *testing.T) {
	render := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	grid := terrain.CreateGrid(1, 1)
	sm := CreateStateMachine()
	obj := CreateFakeMob(sm)
	doing := sm.Doing()
	if doing != "nothing" {
		t.Errorf("Fresh StateMachine should be doing nothing, got %s instead", doing)
	}

	sm.Update(obj, grid)
	doing = sm.Doing()
	if doing != "preparing" {
		t.Errorf("StateMachine should change to idle, got %s", doing)
	}

	sm.Render(obj, render)
}
