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

func (m *interpreter) Interpret(bytecode ByteCode, client ApiClient) {
	size := len(bytecode)
	for i := 0; i < size; i++ {
		instruction := bytecode[i]
		switch instruction {
		case INST_GOTO:
			i = int(m.Pop())
		case INST_LITERAL:
			i++
			value := bytecode[i]
			m.Push(int(value))
		case INST_ADD:
			a, b := m.Pop(), m.Pop()
			m.Push(a + b)
		case INST_SUB:
			a, b := m.Pop(), m.Pop()
			m.Push(b - a)
		case INST_CALL:
			realm := byte(m.Pop())
			action := byte(m.Pop())
			var data []byte
			switch realm {
			case REALM_ITEM:
				switch action {
				case ITEM_COMBINE:
					data = append(data, byte(m.Pop()))
					fallthrough
				case ITEM_PICK, ITEM_BREAK, ITEM_DROP, ITEM_EXAMINE, ITEM_USE:
					data = append(data, byte(m.Pop()))
				}
			}
			value := client.Call(realm, action, data...)
			m.Push(value)
		default:
			panic(fmt.Errorf("unknown instruction %v", instruction))
		}
	}
}
