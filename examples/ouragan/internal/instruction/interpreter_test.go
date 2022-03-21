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
		bytecode []byte
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
				[]byte{},
				&apiClient{},
			},
			[]int{},
			called{},
		},
		{
			"Help",
			args{
				[]byte{byte(INST_LITERAL), byte(GLOB_HELP), byte(INST_GLOB)},
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
				[]byte{byte(INST_LITERAL), 27, byte(INST_LITERAL), byte(ITEM_PICK), byte(INST_ITEM)},
				&apiClient{},
			},
			[]int{},
			called{
				item: []int{int(ITEM_PICK), 27},
			},
		},
		{
			"Combine item 27 on 12",
			args{
				[]byte{byte(INST_LITERAL), 27, byte(INST_LITERAL), 12, byte(INST_LITERAL), byte(ITEM_COMBINE), byte(INST_ITEM)},
				&apiClient{},
			},
			[]int{},
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
