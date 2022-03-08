package object

import (
	"ego/pkg/configuration"
	"testing"
)

func TestStateMob(t *testing.T) {
	var config = configuration.Mob{
		Type: "Mob",
		Needs: []configuration.Need{{
			Type:  "health",
			Level: 100,
		}},
	}
	t.Run("CreateStateMob creates Mob from configuration", func(t *testing.T) {
		o := CreateStateMob(config)
		if _, ok := o.(*stateMob); !ok {
			t.Errorf("Mob should be StateMob, got %+v", o)
		}
	})
	t.Run("StateMob delegates Update to StateMachine", func(t *testing.T) {
		o := CreateStateMob(config)
		o.Update()
		// Checking delegation to do
	})
}
