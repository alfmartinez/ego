package item

import "testing"

func TestItem(t *testing.T) {
	t.Run("CreateItem returns item with chosen type", func(t *testing.T) {
		i := CreateItem(Weapon)
		if i.Type() != Weapon {
			t.Errorf("Expected type weapon, got %v", i.Type())
		}
	})
}
