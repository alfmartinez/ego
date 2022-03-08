package gear

import "testing"

func TestGear(t *testing.T) {

	t.Run("CreateGear should return *gear", func(t *testing.T) {
		g := CreateGear()
		if _, ok := g.(*gear); !ok {
			t.Errorf("expected *gear, got %T", g)
		}
	})
}
