package object

import (
	"testing"

	"github.com/spf13/viper"
)

type fakeObject struct{}

func (f fakeObject) Update() {}

func TestCreateObject(t *testing.T) {

	RegisterObjectFactory("fake", func(key string) GameObject {
		return fakeObject{}
	})

	t.Run("Creates object base on type : 'fake' ", func(t *testing.T) {
		viper.Set("mobs.fake.type", "fake")

		actual := CreateObject("fake")
		if _, ok := actual.(fakeObject); !ok {
			t.Errorf("Should return fakeObject, got %+v", actual)
		}
	})

	t.Run("Creates object base on type : 'Mob' ", func(t *testing.T) {
		viper.Set("mobs.fake", MobData{
			Needs: map[string]int{
				"health": 100,
			},
		})
		viper.Set("mobs.fake.type", "Mob")
		actual := CreateObject("fake")
		if _, ok := actual.(StateMob); !ok {
			t.Errorf("Should return StateMob, got %+v", actual)
		}
	})
}
