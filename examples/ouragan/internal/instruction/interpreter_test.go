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
		bytecode ByteCode
		client   ApiClient
	}
	type called struct {
		global []int
		item   []int
	}
	tests := []struct {
		name   string
		args   args
		want   []byte
		called called
	}{
		{
			"Empty",
			args{
				ByteCode{},
				&apiClient{},
			},
			[]byte{},
			called{},
		},
		{
			"Add",
			args{
				ByteCode{},
				&apiClient{},
			},
			[]byte{},
			called{},
		},
		{
			"Help",
			args{
				ByteCode{INST_LITERAL, GLOB_HELP, INST_GLOB},
				&apiClient{},
			},
			[]byte{},
			called{
				global: []int{int(GLOB_HELP)},
			},
		},
		{
			"Pick item 27",
			args{
				ByteCode{INST_LITERAL, 27, INST_LITERAL, ITEM_PICK, INST_ITEM},
				&apiClient{},
			},
			[]byte{},
			called{
				item: []int{int(ITEM_PICK), 27},
			},
		},
		{
			"Combine item 27 on 12",
			args{
				ByteCode{INST_LITERAL, 27, INST_LITERAL, 12, INST_LITERAL, ITEM_COMBINE, INST_ITEM},
				&apiClient{},
			},
			[]byte{},
			called{
				item: []int{int(ITEM_COMBINE), 27, 12},
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
			stack := CreateStack[byte]()
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
		stack := CreateStack[byte]()
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
