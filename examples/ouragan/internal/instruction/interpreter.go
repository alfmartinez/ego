package instruction

import (
	"fmt"
)

type Interpreter interface {
	Interpret([]Instruction, ApiClient)
}

type interpreter struct {
	Stack[int]
}

func (m *interpreter) Interpret(bytecode []Instruction, client ApiClient) {
	size := len(bytecode)
	for i := 0; i < size; i++ {
		instruction := bytecode[i]
		switch instruction {
		case INST_LITERAL:
			i++
			value := bytecode[i]
			m.Push(int(value))
		case INST_HELP:
			client.Global(GLOB_HELP)
		case INST_PICK:
			value := m.Pop()
			client.Item(ITEM_PICK, value)
		case INST_COMBINE:
			v2 := m.Pop()
			v1 := m.Pop()
			client.Item(ITEM_COMBINE, v1, v2)
		default:
			panic(fmt.Errorf("unknown instruction %v", instruction))

		}
	}
}
