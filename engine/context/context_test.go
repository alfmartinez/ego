package context

import (
	"reflect"
	"testing"
)

func TestCreateContext(t *testing.T) {
	tests := []struct {
		name string
		want Context
	}{
		{
			"create context is simple",
			&context{make(map[string]interface{})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterContext(t *testing.T) {
	type args struct {
		name string
		ctx  Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"register stores context",
			args{
				"foo",
				&context{make(map[string]interface{})},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterContext(tt.args.name, tt.args.ctx)
			if got := GetContext(); !reflect.DeepEqual(got, tt.args.ctx) {
				t.Errorf("GetContext() = %v, want %v", got, tt.args.ctx)
			}

		})
	}
}

func TestCreateAndRegisterContext(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Create and Register",
			args{
				"foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAndRegisterContext(tt.args.name)
		})
	}
}

func Test_context_Set(t *testing.T) {
	type fields struct {
		values map[string]interface{}
	}
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"set & get",
			fields{
				values: make(map[string]interface{}),
			},
			args{
				"foo",
				"baz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &context{
				values: tt.fields.values,
			}
			c.Set(tt.args.name, tt.args.value)
			if got := c.Get(tt.args.name); !reflect.DeepEqual(got, tt.args.value) {
				t.Errorf("ctx.Get() = %v, want %v", got, tt.args.value)
			}
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"global set",
			args{
				"foo",
				"faz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAndRegisterContext("test")
			Set(tt.args.name, tt.args.value)
			if got := Get(tt.args.name); !reflect.DeepEqual(got, tt.args.value) {
				t.Errorf("context.Get() = %v, want %v", got, tt.args.value)
			}
		})
	}
}

func TestSwitchContext(t *testing.T) {
	c1 := CreateContext()
	c2 := CreateContext()
	c3 := CreateContext()
	c1.Set("foo", "beep")
	c2.Set("blah", "grrr")
	RegisterContext("foo", c1)
	RegisterContext("bar", c2)
	RegisterContext("baz", c3)
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Context
	}{
		{
			"C1",
			args{
				"foo",
			},
			c1,
		},
		{
			"C2",
			args{
				"bar",
			},
			c2,
		},
		{
			"C3",
			args{
				"baz",
			},
			c3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SwitchContext(tt.args.name)
			if got := GetContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("context.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
