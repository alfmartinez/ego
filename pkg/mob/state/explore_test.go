package state

import "testing"

func TestExploreLabelReturnsExploring(t *testing.T) {
	s := CreateState("explore")
	l := s.Label()

	if l != "exploring" {
		t.Errorf("Explore should have label exploring, returned %s instead", l)
	}
}

// If machine has no existing interests, add surrounding tiles which have not been explored
// If machine has existing interests, move to closest one
func TestExploreUpdateExploresExistingInterests(t *testing.T) {
	CreateState("explore")

}
