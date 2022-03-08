package gear

import (
	"ego/pkg/item"
)

// Gear is the component in charge of carrying, wearing and using items for a character

type Gear interface {
	Equip(item.Item) bool
	Equipped() item.Item
	Unequip() item.Item
}

func CreateGear() Gear {
	return &gear{}
}

type gear struct {
	equipped item.Item
}

func (g *gear) Equip(item item.Item) bool {
	if g.equipped == nil {
		g.equipped = item
		return true
	}
	return false
}

func (g *gear) Equipped() item.Item {
	return g.equipped
}

func (g *gear) Unequip() item.Item {
	i := g.equipped
	g.equipped = nil
	return i
}
