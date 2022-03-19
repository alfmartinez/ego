package memory

import (
	"ego/engine/movement"
	"image"
	"testing"
)

func TestCreateMemory(t *testing.T) {
	m := CreateMemory()
	if _, ok := m.(*memory); !ok {
		t.Error("CreateMemory should &memory")
	}
}

func TestMemoryCanRecordInterests(t *testing.T) {
	m := CreateMemory()
	if m.HasInterests() {
		t.Error("Newly created memory should have no interests")
	}

	var interest movement.Positionnable

	interest = m.GetNextInterest()
	if interest != nil {
		t.Error("With no interests should return nil")
	}

	origin := movement.Loc(image.Pt(0, 0))
	otherPoint := movement.Loc(image.Pt(1, 1))
	interests := []movement.Positionnable{
		origin,
		otherPoint,
	}
	m.AddInterests(interests)
	if !m.HasInterests() {
		t.Error("Memory should have interests")
	}
	interest = m.GetNextInterest()
	if !interest.IsAt(origin) {
		t.Error("Interest Location should be second")
	}

	interest = m.GetNextInterest()
	if !interest.IsAt(otherPoint) {
		t.Error("Interest should be first")
	}
}

func TestMemoryCanRecordExplorationOfPlaces(t *testing.T) {
	m := CreateMemory()
	coord := movement.Loc(image.Pt(0, 0))
	otherCoord := movement.Loc(image.Pt(1, 1))
	if m.HasExplored(coord) {
		t.Error("Newly create memory cannot know if place is explored")
	}

	var count int
	for !m.HasExplored(coord) {
		m.Explore(coord)
		count++
	}
	if count != 10 {
		t.Errorf("Place should be explored after 10 steps, counted %d", count)
	}
	if m.HasExplored(otherCoord) {
		t.Error("Explore should not affect other place")
	}

}
