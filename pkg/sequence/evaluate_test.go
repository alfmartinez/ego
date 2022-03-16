package sequence

import (
	"ego/pkg/command"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/state"
	"ego/pkg/terrain"
	"image"
	"testing"

	"github.com/spf13/viper"
)

type mockEvaluator struct {
	command.CommandStream
	state.StateMachine
	motivator.NeedsCollection
	movement.Movement
}

func createMockEvaluateActor(need motivator.Need, value int) EvaluateActor {
	stream := command.CreateCommandStream()
	sm := state.CreateStateMachine()
	nc := motivator.CreateNeedsCollection()
	mvmt := movement.CreateMovement(image.Pt(0, 0))
	nc.AddNeed(need, value)
	return &mockEvaluator{stream, sm, nc, mvmt}
}

type fakeTile struct{}

func (f *fakeTile) Path() string { return "fake" }

func TestCreateEvaluateCommand(t *testing.T) {
	terrain.RegisterTileType("plain", &fakeTile{})
	viper.Set("grid", terrain.GridData{
		Size: 10,
		Types: map[string]string{
			"a": "plain",
		},
		Content: "AAA\nAAA\nAAA",
	})
	terrain.CreateGrid(func(terrain.Tile) {})

	t.Run("Need obvious", func(t *testing.T) {
		mob := createMockEvaluateActor(motivator.Health, 10)
		cmd := CreateEvaluateCommand(mob)
		done := cmd.Execute()
		if !done {
			t.Error("Should be done if need is obvious")
		}
	})

	t.Run("Need not obvious", func(t *testing.T) {
		mob := createMockEvaluateActor(motivator.Health, 100)
		cmd := CreateEvaluateCommand(mob)
		done := cmd.Execute()
		if done {
			t.Error("Should not be done if need is not obvious")
		}
	})

}
