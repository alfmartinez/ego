package object

import (
	"testing"

	"github.com/spf13/viper"
)

func TestStateMob(t *testing.T) {
	viper.Set("mobs.fake", MobData{
		Needs: map[string]int{
			"health": 100,
		}})

	t.Run("CreateStateMob creates Mob from configuration", func(t *testing.T) {
		o := CreateStateMob("fake")
		if _, ok := o.(*stateMob); !ok {
			t.Errorf("Mob should be StateMob, got %+v", o)
		}
	})
	t.Run("StateMob delegates Update to StateMachine", func(t *testing.T) {
		o := CreateStateMob("fake")
		o.Update()
		// Checking delegation to do
	})
}
