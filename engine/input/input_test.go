package input

import (
	"ego/engine/context"
	"reflect"
	"testing"
)

func TestFromContext(t *testing.T) {
	tests := []struct {
		name string
		want InputHandler
	}{
		{
			"retrieve input handler from context",
			CreateInputHandler("key"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context.CreateAndRegisterContext("test_input")
			context.Set("input", tt.want)
			if got := FromContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
