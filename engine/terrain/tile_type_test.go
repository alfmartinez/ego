package terrain

import "testing"

func TestCreateTileType(t *testing.T) {
	t.Run("Unknown type should panic", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("Should have panicked")
			}
		}()
		CreateTileType("foo")
	})

}
