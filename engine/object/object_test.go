package object

import (
	"ego/engine/context"
	"ego/engine/observer"
	"testing"

	"github.com/spf13/viper"
)

type fakeObject struct{}

func (f fakeObject) OnNotify(observer.Event) {}
func (f fakeObject) Update()                 {}

func TestCreateObject(t *testing.T) {
	context.CreateAndRegisterContext("test_object")
	cfg := viper.New()
	context.Set("cfg", cfg)

	RegisterObjectFactory("fake", func(key string) GameObject {
		return fakeObject{}
	})

	t.Run("Creates object base on type : 'fake' ", func(t *testing.T) {

		cfg.Set("mobs.fake.type", "fake")

		actual := CreateObject("fake")
		if _, ok := actual.(fakeObject); !ok {
			t.Errorf("Should return fakeObject, got %+v", actual)
		}
	})

	t.Run("Creates object base on type : missing ", func(t *testing.T) {

		cfg.Set("mobs.fake.type", "foo")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("should panic")
			}
		}()

		CreateObject("fake")

	})
}
