package state

import (
	"ego/pkg/configuration"
	"ego/pkg/memory"
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"image"
	"testing"
)

func TestExploreLabelReturnsExploring(t *testing.T) {
	s := CreateState("explore")
	l := s.Label()

	if l != "exploring" {
		t.Errorf("Explore should have label exploring, returned %s instead", l)
	}
}

func TestExploreUpdate(t *testing.T) {
	var next State
	s := CreateState("explore")
	explorer := CreateFakeExplorer(0, 0)
	grid := terrain.CreateGrid(2, 2)

	next = s.Update(explorer, grid)
	if next != nil {
		t.Errorf("Should not return next state on first call, got %+v", next)
	}
	if !explorer.HasInterests() {
		t.Error("Explorer should have POI to explore")
	}

	destination := s.(*exploreState).exploring.Position()
	if destination != image.Pt(0, 0) {
		t.Errorf("Should have chosen (0,0), got %+v", destination)
	}

	var countExplore int
	for next == nil {
		next = s.Update(explorer, grid)
		countExplore++
	}

	if countExplore != 10 {
		t.Errorf("Explore should take 10, got %d", countExplore)
	}

	if _, ok := next.(*idleState); !ok {
		t.Errorf("Should return IdleState, got %+v", next)
	}

	s = CreateState("explore")
	next = nil
	for next == nil {
		next = s.Update(explorer, grid)
	}

	if _, ok := next.(*moveState); !ok {
		t.Errorf(" returned state %+v", next)
	}

}

func TestExploreRender(t *testing.T) {
	s := CreateState("explore")
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	m := CreateFakeRenderable()
	s.Render(r, m)
}

func CreateFakeRenderable() renderable.Renderable {
	return &fakeRenderable{}
}

type fakeRenderable struct{}

func (f *fakeRenderable) Doing() string         { return "fake" }
func (f *fakeRenderable) Name() string          { return "fake" }
func (f *fakeRenderable) Path() string          { return "fake" }
func (f *fakeRenderable) Position() image.Point { return image.Pt(0, 0) }
func (f *fakeRenderable) Size() uint            { return 1 }

type fakeExplorer struct {
	memory.Memory
	movement.Positionnable
}

func CreateFakeExplorer(x, y int) Explorer {
	memo := memory.CreateMemory()
	return &fakeExplorer{memo, movement.Loc(image.Pt(x, y))}
}
