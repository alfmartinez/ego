package physics

import (
	"ego/pkg/configuration"
	"ego/pkg/context"
	"testing"
)

func TestCreateEngine(t *testing.T) {
	context.CreateAndRegisterContext("test")
	cfg := configuration.CreateConfiguration()
	context.Set("cfg", cfg.Get())
	t.Run("CreatePhysicsEngine", func(t *testing.T) {
		context.Set("physics.modules", []string{"fake"})
		engine := CreatePhysicsEngine()
		engine.Init()
	})

}
