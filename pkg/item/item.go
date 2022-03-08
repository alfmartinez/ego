package item

type ItemType int

const (
	Equippable ItemType = 1 << iota
	Apparel
	Weapon
	Consumable
)

type Item interface {
}

func CreateItem(itemType ItemType) Item {
	return &item{itemType: itemType}
}

type item struct {
	itemType ItemType
}
