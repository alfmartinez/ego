package object

import (
	"ego/pkg/configuration"
	"testing"
)

type fakeObject struct{}

func (f fakeObject) Update() {}

func TestCreateObject(t *testing.T) {

	RegisterObjectFactory("fake", func(m configuration.Mob) GameObject {
		return fakeObject{}
	})

	t.Run("Creates object base on type : 'fake' ", func(t *testing.T) {
		var config = configuration.Mob{
			Type: "fake",
		}

		actual := CreateObject(config)
		if _, ok := actual.(fakeObject); !ok {
			t.Errorf("Should return fakeObject, got %+v", actual)
		}
	})

	t.Run("Creates object base on type : 'Mob' ", func(t *testing.T) {
		var config = configuration.Mob{
			Type: "Mob",
			Needs: []configuration.Need{{
				Type:  "health",
				Level: 100,
			}},
		}

		actual := CreateObject(config)
		if _, ok := actual.(StateMob); !ok {
			t.Errorf("Should return StateMob, got %+v", actual)
		}
	})
}
