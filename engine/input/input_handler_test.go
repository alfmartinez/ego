package input

import (
	"reflect"
	"testing"
)

func TestCreateInputHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  InputHandler
		panic bool
	}{
		{
			"Existing key handler",
			args{
				"key",
			},
			CreateKeyHandler(),
			false,
		},
		{
			"Missing foo handler",
			args{
				"foo",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil && tt.panic {
					t.Error("should have panicked")
				}
				if r != nil && !tt.panic {
					t.Error("shoud not have panicked")
				}
			}()
			if got := CreateInputHandler(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInputHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
