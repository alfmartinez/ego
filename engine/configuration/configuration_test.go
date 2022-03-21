package configuration

import (
	"ego/engine/context"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestFromContext(t *testing.T) {
	cfg := viper.New()
	context.CreateAndRegisterContext("test")
	context.Set("cfg", cfg)
	tests := []struct {
		name string
		want *viper.Viper
	}{
		{
			"Retrieve viper from context",
			cfg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateConfiguration(t *testing.T) {
	cfg := viper.New()
	cfg.BindEnv("lang")
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want ConfigurationFactory
	}{
		{
			"creates configuration factory",
			args{
				path: "foo",
			},
			&configuration{cfg, "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfiguration(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configuration_Get(t *testing.T) {
	type fields struct {
		viper *viper.Viper
		path  string
	}
	tests := []struct {
		name   string
		fields fields
		want   *viper.Viper
	}{
		{
			"Returns viper",
			fields{
				viper.New(),
				"foo",
			},
			viper.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				viper: tt.fields.viper,
				path:  tt.fields.path,
			}
			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configuration.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configuration_Init(t *testing.T) {
	type fields struct {
		viper *viper.Viper
		path  string
	}
	type want struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		panic  bool
		want   want
	}{
		{
			"wrong path should panic",
			fields{
				viper.New(),
				"foo",
			},
			true,
			want{"foo", "bar"},
		},
		{
			"existing path should not panic",
			fields{
				viper.New(),
				"testdata/simple/",
			},
			false,
			want{"foo", "bar"},
		},
		{
			"missing import should panic",
			fields{
				viper.New(),
				"testdata/missing_import/",
			},
			true,
			want{"foo", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				viper: tt.fields.viper,
				path:  tt.fields.path,
			}
			defer func() {
				r := recover()
				if r == nil && tt.panic {
					t.Error("should panic")
				}
				if r != nil && !tt.panic {
					t.Error("should not panic")
				}
			}()

			c.Init()
			if !tt.panic {
				if c.Get().GetString(tt.want.key) != tt.want.value {
					t.Error("wrong value")
				}
			}
		})
	}
}
