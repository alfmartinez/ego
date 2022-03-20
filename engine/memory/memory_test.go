package memory

import (
	"ego/engine/movement"
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

	interest, found := m.GetNextInterest()
	if found {
		t.Error("With no interests should return false on found")
	}

	origin := movement.Location{X: 0, Y: 0}
	otherPoint := movement.Location{X: 1, Y: 1}
	interests := []movement.Location{
		origin,
		otherPoint,
	}
	m.AddInterests(interests)
	if !m.HasInterests() {
		t.Error("Memory should have interests")
	}
	interest, found = m.GetNextInterest()
	if interest != origin {
		t.Error("Interest Location should be second")
	}

	interest, found = m.GetNextInterest()
	if interest != otherPoint {
		t.Error("Interest should be first")
	}
}

func TestMemoryCanRecordExplorationOfPlaces(t *testing.T) {
	m := CreateMemory()
	coord := movement.Location{}
	otherCoord := movement.Location{X: 1, Y: 1}
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
