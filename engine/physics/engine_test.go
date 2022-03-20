package physics

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/physics/module"
	"reflect"
	"testing"
	"time"
)

func TestCreateEngine(t *testing.T) {

	t.Run("CreatePhysicsEngine", func(t *testing.T) {
		t.Run("no modules", func(t *testing.T) {
			initConfig()
			configuration.FromContext().Set("physics.modules", []string{})
			CreatePhysicsEngine()
		})
		t.Run("unknown module", func(t *testing.T) {
			initConfig()

			configuration.FromContext().Set("physics.modules", []string{"foo"})

			defer func() {
				r := recover()
				if r == nil {
					t.Error("should panic")
				}
			}()

			CreatePhysicsEngine()
		})
	})

}

func initConfig() {
	context.CreateAndRegisterContext("test")
	cfg := configuration.CreateConfiguration("")
	context.Set("cfg", cfg.Get())
}

func TestFromContext(t *testing.T) {
	tests := []struct {
		name string
		want Engine
	}{
		{
			"blank engine",
			&phyiscsEngine{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context.CreateAndRegisterContext("test_phy")
			context.Set("physics", tt.want)
			if got := FromContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fakeModule struct{}

var mockInit func()
var mockAdd func(any)
var mockAdvance func(float64)

func (f *fakeModule) Init()              { mockInit() }
func (f *fakeModule) Add(a any)          { mockAdd(a) }
func (f *fakeModule) Advance(dt float64) { mockAdvance(dt) }

func Test_phyiscsEngine_Init(t *testing.T) {

	type fields struct {
		modules []module.Module
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"single module",
			fields{
				modules: []module.Module{
					&fakeModule{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" Init", func(t *testing.T) {
			e := &phyiscsEngine{
				modules: tt.fields.modules,
			}
			var called bool
			mockInit = func() {
				called = true
			}
			e.Init()
			if !called {
				t.Error("Init should be called")
			}
		})
		t.Run(tt.name+" Add", func(t *testing.T) {
			e := &phyiscsEngine{
				modules: tt.fields.modules,
			}

			var arg any
			mockAdd = func(a any) {
				arg = a
			}
			e.Add("foo")
			if arg != "foo" {
				t.Error("Add should be called")
			}
		})
		t.Run(tt.name+" Advance", func(t *testing.T) {
			e := &phyiscsEngine{
				modules: tt.fields.modules,
			}
			var arg float64
			mockAdvance = func(a float64) {
				if a != 1 {
					t.Errorf("expected %f, got %f", 1.0, a)
				}
				arg = a
			}
			e.Advance(time.Second)
			if arg != 1 {
				t.Errorf("Advance should be called, got %v", arg)
			}

		})
	}
}
