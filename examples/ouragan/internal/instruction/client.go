package instruction

type ItemAction int
type GlobalAction int
type PlaceAction int

const (
	ITEM_PICK ItemAction = iota
	ITEM_DROP
	ITEM_BREAK
	ITEM_USE
	ITEM_EXAMINE
)

const (
	GLOB_HELP GlobalAction = iota
	GLOB_INVENTORY
)

const (
	PLACE_GO PlaceAction = iota
)

type ApiClient interface {
	Global(GlobalAction)
	Item(ItemAction, ...int)
}
