package gear

import (
	"ego/pkg/item"
	"testing"
)

func TestGear(t *testing.T) {

	t.Run("CreateGear should return *gear", func(t *testing.T) {
		g := CreateGear()
		if _, ok := g.(*gear); !ok {
			t.Errorf("expected *gear, got %T", g)
		}
	})

	t.Run("Gear can equip equippable", func(t *testing.T) {
		g := CreateGear()
		item := item.CreateItem(item.Equippable)
		g.Equip(item)
	})

	t.Run("Gear cannot equip item if already equipped", func(t *testing.T) {
		g := CreateGear()
		item1 := item.CreateItem(item.Equippable)
		item2 := item.CreateItem(item.Equippable)

		if !g.Equip(item1) {
			t.Error("Equip should return true")
		}

		if g.Equip(item2) {
			t.Error("Equip should return false if item already equipped")
		}

	})
}
