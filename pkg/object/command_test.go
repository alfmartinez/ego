package object

import (
	"ego/pkg/configuration"
	"testing"
)

func TestCreateEvaluateCommand(t *testing.T) {
	t.Run("Need obvious", func(t *testing.T) {
		mob := CreateStateMob(configuration.Mob{
			Needs: []configuration.Need{
				{
					Type:  "health",
					Level: 1,
				},
			},
		})
		sm, _ := mob.(StateMob)
		cmd := CreateEvaluateCommand(sm)
		done := cmd.Execute()
		if !done {
			t.Error("Should be done if need is obvious")
		}
	})

	t.Run("Need not obvious", func(t *testing.T) {
		mob := CreateStateMob(configuration.Mob{
			Needs: []configuration.Need{
				{
					Type:  "health",
					Level: 100,
				},
			},
		})
		sm, _ := mob.(StateMob)
		cmd := CreateEvaluateCommand(sm)
		done := cmd.Execute()
		if done {
			t.Error("Should not be done if need is not obvious")
		}
	})

}
