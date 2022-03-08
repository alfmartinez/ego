package gear

import (
	"ego/pkg/item"
	"testing"
)

func TestSlot(t *testing.T) {

	t.Run("CreateGear should return *gear", func(t *testing.T) {
		g := CreateSlot(item.Equippable)
		if _, ok := g.(*slot); !ok {
			t.Errorf("expected *gear, got %T", g)
		}
	})

	t.Run("Gear can equip equippable", func(t *testing.T) {
		g := CreateSlot(item.Equippable)
		i := item.CreateItem(item.Equippable)
		g.Equip(i)
	})

	t.Run("Gear can tell if item is acceptable", func(t *testing.T) {
		g := CreateSlot(item.Equippable)
		t.Run("Acceptable", func(t *testing.T) {
			i := item.CreateItem(item.Equippable)
			ok := g.Accepts(i)
			if !ok {
				t.Error("item should be acceptable")
			}
		})
		t.Run("Not Acceptable", func(t *testing.T) {
			i := item.CreateItem(item.Weapon)
			ok := g.Accepts(i)
			if ok {
				t.Error("item should not be acceptable")
			}
		})
	})

	t.Run("Compound types", func(t *testing.T) {
		g := CreateSlot(item.Equippable | item.Consumable)
		t.Run("Compound", func(t *testing.T) {
			i := item.CreateItem(item.Equippable | item.Apparel)
			ok := g.Accepts(i)
			if !ok {
				t.Error("item should be acceptable")
			}
		})
	})

	t.Run("Gear can retrieve item", func(t *testing.T) {
		g := CreateSlot(item.Equippable)
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
		g := CreateSlot(item.Equippable)
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
		g := CreateSlot(item.Equippable)
		item1, item2 := item.CreateItem(item.Equippable), item.CreateItem(item.Equippable)

		if !g.Equip(item1) {
			t.Error("Equip should return true")
		}

		if g.Equip(item2) {
			t.Error("Equip should return false if item already equipped")
		}

	})
}
