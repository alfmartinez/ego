package instruction

type Instruction byte
type ItemAction byte
type GlobalAction byte
type PlaceAction byte

type ByteCode []byte

const (
	INST_NONE    byte = iota // Help
	INST_LITERAL             // Literal value
	INST_GLOB                // Execute Global Action
	INST_ITEM                // Execute Item Action
	INST_ADD                 // Add
	INST_SUB                 // SUB
	INST_GOTO
)

const (
	ITEM_PICK byte = iota
	ITEM_DROP
	ITEM_BREAK
	ITEM_USE
	ITEM_EXAMINE
	ITEM_COMBINE
)

const (
	GLOB_HELP byte = iota
	GLOB_INVENTORY
	GLOB_EXIT
)

const (
	PLACE_GO PlaceAction = iota
)
