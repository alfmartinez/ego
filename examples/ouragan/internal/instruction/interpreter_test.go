package instruction

import (
	"reflect"
	"testing"
)

type apiClient struct{}

var mockGlobal func(GlobalAction)
var mockItem func(ItemAction, ...int)

func (c *apiClient) Global(a GlobalAction) {
	mockGlobal(a)
}
func (c *apiClient) Item(a ItemAction, data ...int) {
	mockItem(a, data...)
}

func Test_interpreter_Interpret(t *testing.T) {
	type args struct {
		bytecode []Instruction
		client   ApiClient
	}
	type called struct {
		global []int
		item   []int
	}
	tests := []struct {
		name   string
		args   args
		want   []int
		called called
	}{
		{
			"Empty",
			args{
				[]Instruction{},
				&apiClient{},
			},
			[]int{},
			called{},
		},
		{
			"Help",
			args{
				[]Instruction{INST_HELP},
				&apiClient{},
			},
			[]int{},
			called{
				global: []int{int(GLOB_HELP)},
			},
		},
		{
			"Pick item 27",
			args{
				[]Instruction{INST_LITERAL, 27, INST_PICK},
				&apiClient{},
			},
			[]int{},
			called{
				item: []int{int(ITEM_PICK), 27},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gotCalled := called{}
			mockGlobal = func(a GlobalAction) {
				gotCalled.global = append(gotCalled.global, int(a))
			}
			mockItem = func(a ItemAction, data ...int) {
				gotCalled.item = append(gotCalled.item, int(a))
				gotCalled.item = append(gotCalled.item, data...)
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
}
