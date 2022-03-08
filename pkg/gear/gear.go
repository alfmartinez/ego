package gear

import "ego/pkg/item"

// Gear is the component in charge of carrying, wearing and using items for a character

type Gear interface {
	Equip(item.Item) bool
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
