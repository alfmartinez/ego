package physics

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"testing"
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
