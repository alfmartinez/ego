package instruction

type Instruction byte

const (
	INST_HELP      Instruction = iota // Help
	INST_INVENTORY                    // Inventory
	INST_EXIT                         // Exit game
	INST_LOOK                         // Look at target
	INST_GO                           // Go somewhere
	INST_PICK                         // Pick item
	INST_DROP                         // Drop item
	INST_USE                          // Use item
	INST_COMBINE                      // Use item with another item
	INST_EXAMINE                      // Examine item
	INST_BREAK                        // Break item
	INST_LITERAL                      // Literal value
)
