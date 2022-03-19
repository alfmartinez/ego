package object

import (
	"ego/engine/observer"
	"testing"

	"github.com/spf13/viper"
)

func TestStateMob(t *testing.T) {
	viper.Set("mobs.fake", MobData{})

	t.Run("CreateStateMob creates Mob from configuration", func(t *testing.T) {
		o := CreateStateMob("fake")
		if _, ok := o.(*stateMob); !ok {
			t.Errorf("Mob should be StateMob, got %+v", o)
		}
	})
	t.Run("StateMob delegates Update to StateMachine", func(t *testing.T) {
		o := CreateStateMob("fake")
		o.OnNotify(observer.CreateEvent(observer.UPDATE))
		// Checking delegation to do
	})
}
