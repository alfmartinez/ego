package instruction

import (
	"reflect"
	"testing"
)

type apiClient struct{}

var mockCall func(realm byte, action byte, data ...byte) int

func (c *apiClient) Call(realm byte, action byte, data ...byte) int {
	return mockCall(realm, action, data...)
}

func Test_interpreter_Interpret(t *testing.T) {
	type args struct {
		bytecode ByteCode
		client   ApiClient
	}
	tests := []struct {
		name        string
		args        args
		want        []int
		called      []byte
		returnValue int
	}{
		{
			"Empty",
			args{
				ByteCode{},
				&apiClient{},
			},
			[]int{},
			[]byte{},
			0,
		},
		{
			"Pop",
			args{
				ByteCode{INST_LITERAL, 34, INST_POP},
				&apiClient{},
			},
			[]int{},
			[]byte{},
			0,
		},
		{
			"Add",
			args{
				ByteCode{INST_LITERAL, 1, INST_LITERAL, 3, INST_ADD},
				&apiClient{},
			},
			[]int{4},
			[]byte{},
			0,
		},
		{
			"Sub",
			args{
				ByteCode{INST_LITERAL, 4, INST_LITERAL, 3, INST_SUB},
				&apiClient{},
			},
			[]int{1},
			[]byte{},
			0,
		},
		{
			"Sub neg",
			args{
				ByteCode{INST_LITERAL, 4, INST_LITERAL, 6, INST_SUB},
				&apiClient{},
			},
			[]int{-2},
			[]byte{},
			0,
		},
		{
			"Help",
			args{
				ByteCode{INST_LITERAL, GLOB_HELP, INST_LITERAL, REALM_GLOB, INST_CALL},
				&apiClient{},
			},
			[]int{0},
			[]byte{REALM_GLOB, GLOB_HELP},
			0,
		},
		{
			"goto",
			args{
				ByteCode{INST_LITERAL, 1, INST_LITERAL, 5, INST_LITERAL, 9, INST_GOTO, INST_LITERAL, 4, INST_LITERAL, 12},
				&apiClient{},
			},
			[]int{12, 5, 1},
			[]byte{},
			0,
		},
		{
			"if / true",
			args{
				ByteCode{
					INST_LITERAL, 56, // Previous value
					INST_LITERAL, 15, // Finally
					INST_LITERAL, 12, // True subroutine
					INST_LITERAL, 1, // Condition
					INST_IF, // THE IF !
					INST_LITERAL, 13,
					INST_RETURN,
					INST_LITERAL, 42,
					INST_RETURN,
					INST_LITERAL, 24,
				},
				&apiClient{},
			},
			[]int{24, 42, 56},
			[]byte{},
			0,
		},
		{
			"gosub/return",
			args{
				ByteCode{
					INST_LITERAL, 1, INST_LITERAL, 5, // 0
					INST_LITERAL, 12, INST_GOSUB, // 4
					INST_LITERAL, 4, // 7
					INST_LITERAL, 15, INST_GOTO, // 9
					INST_LITERAL, 52, INST_RETURN, // 12
					INST_LITERAL, 32, // 14
				},
				&apiClient{},
			},
			[]int{32, 4, 52, 5, 1},
			[]byte{},
			0,
		},
		{
			"Pick item 27",
			args{
				ByteCode{INST_LITERAL, 27, INST_LITERAL, ITEM_PICK, INST_LITERAL, REALM_ITEM, INST_CALL},
				&apiClient{},
			},
			[]int{0},
			[]byte{REALM_ITEM, ITEM_PICK, 27},
			0,
		},
		{
			"Combine item 27 on 12",
			args{
				ByteCode{INST_LITERAL, 27, INST_LITERAL, 12, INST_LITERAL, ITEM_COMBINE, INST_LITERAL, REALM_ITEM, INST_CALL},
				&apiClient{},
			},
			[]int{0},
			[]byte{REALM_ITEM, ITEM_COMBINE, 12, 27},
			0,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gotCalled := make([]byte, 0)
			mockCall = func(realm byte, action byte, data ...byte) int {
				gotCalled = append(gotCalled, realm, action)
				gotCalled = append(gotCalled, data...)
				return tt.returnValue
			}
			stack := CreateStack[int]()
			m := &interpreter{
				Stack: stack,
			}
			m.Interpret(tt.args.bytecode, tt.args.client)
			if got := stack.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expecting %v, got %v", tt.want, got)
			}
			if !reflect.DeepEqual(gotCalled, tt.called) {
				t.Errorf("Expecting calls %v, got %v", tt.called, gotCalled)
			}
		})
	}
	t.Run("Unknown instruction causes panic", func(t *testing.T) {
		stack := CreateStack[int]()
		m := &interpreter{
			Stack: stack,
		}
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should panic")
			}
		}()
		m.Interpret([]byte{255}, &apiClient{})
	})
}
