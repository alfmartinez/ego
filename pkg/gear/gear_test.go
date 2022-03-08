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
		i := item.CreateItem(item.Equippable)
		g.Equip(i)
	})

	t.Run("Gear can retrieve item", func(t *testing.T) {
		g := CreateGear()
		i := item.CreateItem(item.Equippable)
		g.Equip(i)
		actual := g.Equipped()
		if actual != i {
			t.Errorf("should return equipped item, got %+v", actual)
		}
		actual = g.Equipped()
		if actual != i {
			t.Errorf("should return equipped item, got %+v", actual)
		}

	})

	t.Run("Gear can unequip item", func(t *testing.T) {
		g := CreateGear()
		i := item.CreateItem(item.Equippable)
		g.Equip(i)
		actual := g.Unequip()
		if actual != i {
			t.Errorf("should return equipped item, got %+v", actual)
		}
		actual = g.Equipped()
		if actual != nil {
			t.Errorf("should return equipped item, got %+v", actual)
		}

	})

	t.Run("Gear cannot equip item if already equipped", func(t *testing.T) {
		g := CreateGear()
		item1, item2 := item.CreateItem(item.Equippable), item.CreateItem(item.Equippable)

		if !g.Equip(item1) {
			t.Error("Equip should return true")
		}

		if g.Equip(item2) {
			t.Error("Equip should return false if item already equipped")
		}

	})
}
