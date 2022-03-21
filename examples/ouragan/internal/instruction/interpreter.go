package instruction

import (
	"fmt"
)

type Interpreter interface {
	Interpret([]Instruction, ApiClient)
}

type interpreter struct {
	Stack[byte]
}

func (m *interpreter) Interpret(bytecode ByteCode, client ApiClient) {
	size := len(bytecode)
	for i := 0; i < size; i++ {
		instruction := bytecode[i]
		switch instruction {
		case INST_LITERAL:
			i++
			value := bytecode[i]
			m.Push(value)
		case INST_GLOB:
			action := m.Pop()
			client.Global(GlobalAction(action))
		case INST_ITEM:
			action := m.Pop()
			switch byte(action) {
			case ITEM_PICK, ITEM_BREAK, ITEM_DROP, ITEM_EXAMINE, ITEM_USE:
				subject := m.Pop()
				client.Item(ItemAction(action), int(subject))
			case ITEM_COMBINE:
				object := m.Pop()
				subject := m.Pop()
				client.Item(ItemAction(action), int(subject), int(object))
			}
		default:
			panic(fmt.Errorf("unknown instruction %v", instruction))

		}
	}
}
