package terrain

import (
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"image"
	"testing"
)

func TestCreateGrid(t *testing.T) {
	g := CreateGrid(3, 3)

	if _, ok := g.(*grid); !ok {
		t.Errorf("Should return Grid, got %+v", g)
	}
	var cases = []struct {
		x, y     int
		expected int
	}{
		{0, 0, 3},
	}
	for _, r := range cases {
		pt := image.Pt(r.x, r.y)
		l := len(g.Tile(pt).Surrounding())
		if l != r.expected {
			t.Errorf("Expected %d surrounding %s, got %d", r.expected, pt, l)
		}
	}
}

func TestGridGetTile(t *testing.T) {
	g := CreateGrid(2, 2)
	var tile = g.GetTile(movement.Loc(image.Pt(0, 0)))
	if tile == nil {
		t.Error("Should return origin tile")
	}

	tile = g.GetTile(movement.Loc(image.Pt(100, 100)))
	if tile == nil {
		t.Error("Should return middle tile")
	}

	tile = g.GetTile(movement.Loc(image.Pt(200, 200)))
	if tile != nil {
		t.Error("Should return nil on outer")
	}
}

type fakeRenderer struct {
	Rendered []renderable.Renderable
}

func (f *fakeRenderer) Init()                {}
func (f *fakeRenderer) Start(exit chan bool) {}
func (f *fakeRenderer) IsAsync() bool        { return false }
func (f *fakeRenderer) Refresh()             {}
func (f *fakeRenderer) Render(r renderable.Renderable) {
	f.Rendered = append(f.Rendered, r)
}

func createFakeRenderer() renderer.Renderer {
	return &fakeRenderer{}
}

func TestRender(t *testing.T) {
	g := CreateGrid(1, 1)
	r := createFakeRenderer()

	g.Render(r)

	fr := r.(*fakeRenderer)
	rendered := fr.Rendered
	if len(rendered) != 1 {
		t.Errorf("Should have rendered 1 tile, got %d", len(rendered))
	}
}

func TestFindClosestSingleTrue(t *testing.T) {
	g := CreateGrid(1, 1)
	pos := movement.Loc(image.Pt(0, 0))
	tile := g.FindClosest(pos, 1, func(tile Tile) bool {
		return true
	})
	if tile == nil {
		t.Error("Should return single tile")
	}
}

func TestFindClosestMultipleFalse(t *testing.T) {
	g := CreateGrid(3, 3)
	pos := movement.Loc(image.Pt(0, 0))
	tile := g.FindClosest(pos, 1, func(tile Tile) bool {
		return false
	})[0]

	if tile != nil {
		t.Errorf("Should return nil, got %+v", tile)
	}
}

func TestFindClosestMultipleResource(t *testing.T) {

	g := CreateGrid(3, 3)
	pos := movement.Loc(image.Pt(0, 0))
	g.AddSource(2, 2, "health", 1)
	tile := g.FindClosest(pos, 1, func(tile Tile) bool {
		return tile.HasResource("health")
	})
	if tile == nil {
		t.Error("Should return tile with resource")
	}
}
