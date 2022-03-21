package instruction

type Instruction byte
type ItemAction byte
type GlobalAction byte
type PlaceAction byte

const (
	INST_NONE    Instruction = iota // Help
	INST_LITERAL                    // Literal value
	INST_GLOB                       // Execute Global Action
	INST_ITEM
)

const (
	ITEM_PICK ItemAction = iota
	ITEM_DROP
	ITEM_BREAK
	ITEM_USE
	ITEM_EXAMINE
	ITEM_COMBINE
)

const (
	GLOB_HELP GlobalAction = iota
	GLOB_INVENTORY
	GLOB_EXIT
)

const (
	PLACE_GO PlaceAction = iota
)
