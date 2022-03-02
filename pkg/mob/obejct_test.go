package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/mob/state"
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
	if _, ok := actual.(state.StateMachine); !ok {
		t.Error("Should return StateMachine")
	}
}
