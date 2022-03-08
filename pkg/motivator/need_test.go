package motivator

import "testing"

func TestNeed(t *testing.T) {

	var needs = map[string]Need{
		"health":     Health,
		"food":       Food,
		"rest":       Rest,
		"learn":      Learn,
		"outdoor":    Outdoor,
		"indoor":     Indoor,
		"recreation": Recreation,
		"beauty":     Beauty,
	}
	for label, expected := range needs {
		t.Run(label, func(t *testing.T) {
			need := CreateNeed(label)
			if need != expected {
				t.Errorf("Expected %v, got %v", expected, need)
			}
		})
	}

}
