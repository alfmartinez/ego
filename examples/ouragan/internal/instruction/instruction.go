package instruction

type Instruction byte
type Realm byte
type ItemAction byte
type GlobalAction byte
type PlaceAction byte

type ByteCode []byte

const (
	INST_NONE    byte = iota // Help
	INST_LITERAL             // Literal value
	INST_CALL                // Execute Action
	INST_ADD                 // Add
	INST_SUB                 // SUB
	INST_GOTO
	INST_GOSUB
	INST_RETURN
	INST_POP
	INST_IF
)

const (
	REALM_NONE byte = iota
	REALM_GLOB
	REALM_ITEM
	REALM_CHARACTER
	REALM_PLACE
)

const (
	ITEM_NONE byte = iota
	ITEM_PICK
	ITEM_DROP
	ITEM_BREAK
	ITEM_USE
	ITEM_EXAMINE
	ITEM_COMBINE
)

const (
	GLOB_NONE byte = iota
	GLOB_HELP
	GLOB_INVENTORY
	GLOB_EXIT
)

const (
	PLACE_NONE byte = iota
	PLACE_GO
)
