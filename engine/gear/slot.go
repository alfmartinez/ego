package gear

import "ego/engine/item"

// Gear is the component in charge of carrying, wearing and using items for a character

type Slot interface {
	Accepts(item.Item) bool
	Equip(item.Item) bool
	Equipped() item.Item
	Unequip() item.Item
}

func CreateSlot(i item.ItemType) Slot {
	return &slot{
		accepts: i,
	}
}

type slot struct {
	accepts  item.ItemType
	equipped item.Item
}

func (g *slot) Accepts(item item.Item) bool {
	return item.Type()&g.accepts != 0
}

func (g *slot) Equip(i item.Item) bool {
	if g.equipped == nil && g.Accepts(i) {
		g.equipped = i
		return true
	}
	return false
}

func (g *slot) Equipped() item.Item {
	return g.equipped
}

func (g *slot) Unequip() item.Item {
	i := g.equipped
	g.equipped = nil
	return i
}
