package object

import (
	"ego/pkg/configuration"
	"testing"
)

func TestNew(t *testing.T) {
	var config = configuration.Mob{
		Type: "Mob",
		Needs: []configuration.Need{{
			Type:     "building mocks",
			Priority: 0,
			Level:    100,
		}},
	}

	actual := CreateObject(config)
	if _, ok := actual.(StateMob); !ok {
		//t.Errorf("Should return StateMob, got %+v", actual)
	}
}
