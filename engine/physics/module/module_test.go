package module

import (
	"reflect"
	"testing"
)

func TestCreateModule(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  Module
		panic bool
	}{
		{
			"missing module",
			args{"missing"},
			nil,
			true,
		},
		{
			"gravity module",
			args{"gravity"},
			&gravity{},
			false,
		},
		{
			"collider module",
			args{"collider"},
			&collider{},
			false,
		},
		{
			"simulate module",
			args{"simulate"},
			&simulate{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tt.panic && r == nil {
					t.Error("should panic")
				}
				if !tt.panic && r != nil {
					t.Error("should panic")
				}
			}()

			if got := CreateModule(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateModule() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
