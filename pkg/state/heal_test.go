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

func TestCreateHealState(t *testing.T) {
	s := CreateState("heal")
	label := s.Label()
	if label != "healing" {
		t.Errorf("Label was %s", label)
	}
}

func TestHealRender(t *testing.T) {
	s := CreateState("heal")
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	m := CreateFakeRenderable()
	s.Render(r, m)
}

func TestHealUpdateMedicineOnCurrentTile(t *testing.T) {
	patient := CreateFakePatient(0, 0)
	s := CreateState("heal")
	grid := terrain.CreateGrid(1, 1)
	grid.AddSource(0, 0, "health", 1)

	var next State = nil

	for next == nil {
		next = s.Update(patient, grid)
	}

	if !patient.IsAt(movement.Loc(image.Pt(0, 0))) {
		t.Error("Patient should not move")
	}
	tile := grid.Tile(image.Pt(0, 0))
	if tile.HasResource("health") {
		t.Error("Resource should be exhausted")
	}

}

func TestHealUpdateMedicineOnAnotherTile(t *testing.T) {
	patient := CreateFakePatient(0, 0)
	s := CreateState("heal")
	grid := terrain.CreateGrid(2, 2)
	grid.AddSource(1, 1, "health", 1)

	var next State = nil

	for next == nil {
		next = s.Update(patient, grid)
	}

	if _, ok := next.(*moveState); !ok {
		t.Errorf("Patient should move, got %+v", next)
	}
	tile := grid.Tile(image.Pt(0, 0))
	if tile.HasResource("health") {
		t.Error("Resource should be exhausted")
	}

}

func TestHealUpdateMedicineOnYetAnotherTile(t *testing.T) {
	patient := CreateFakePatient(0, 0)
	s := CreateState("heal")
	grid := terrain.CreateGrid(3, 3)
	grid.AddSource(2, 2, "health", 1)

	var next State = nil

	for next == nil {
		next = s.Update(patient, grid)
	}

	if _, ok := next.(*moveState); !ok {
		t.Errorf("Patient should move, got %+v", next)
	}
	tile := grid.Tile(image.Pt(0, 0))
	if tile.HasResource("health") {
		t.Error("Resource should be exhausted")
	}

}

func CreateFakePatient(x, y int) Patient {
	needs := motivator.CreateNeedsCollection()
	needs.AddNeed(motivator.CreateNeed("health", 0), 30)
	return &fakePatient{
		movement.Loc(image.Pt(x, y)),
		needs,
	}
}

type fakePatient struct {
	movement.Positionnable
	motivator.NeedsCollection
}
