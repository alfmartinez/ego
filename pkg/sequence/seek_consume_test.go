package sequence

import (
	"ego/pkg/motivator"
	"ego/pkg/terrain"
	"testing"

	"github.com/spf13/viper"
)

func TestCreateSeekAndUseCommand(t *testing.T) {

	t.Run("Not found", func(t *testing.T) {
		terrain.RegisterTileType("plain", &fakeTile{})
		viper.Set("grid", terrain.GridData{
			Size: 100,
			Types: map[string]string{
				"a": "plain",
			},
			Content: "AAA\nAAA\nAAA",
		})
		terrain.CreateGrid(func(terrain.Tile) {})
		mob := createMockEvaluateActor(motivator.Health, 10)
		cmd := CreateSeekAndUseCommand(mob, motivator.Health)
		done := cmd.Execute()
		if done {
			t.Error("Should not be done")
		}
		done = cmd.Execute()
		if !done {
			t.Errorf("Sequence should Abort, %+v", cmd)
		}
	})
	t.Run("Found", func(t *testing.T) {
		terrain.RegisterTileType("plain", &fakeTile{})
		viper.Set("grid", terrain.GridData{
			Size: 100,
			Types: map[string]string{
				"a": "plain",
			},
			Content: "AAA\nAAA\nAAA",
		})
		grid := terrain.CreateGrid(func(terrain.Tile) {})
		grid.AddSource(1, 0, terrain.Medicine, 1)
		mob := createMockEvaluateActor(motivator.Health, 10)
		cmd := CreateSeekAndUseCommand(mob, motivator.Health)
		done := cmd.Execute()
		if done {
			t.Error("Should not be done")
		}
		done = cmd.Execute()
		if done {
			t.Error("Sequence should not be done")
		}
		var calledCount int = 0
		for !done {
			calledCount++
			done = cmd.Execute()
		}
		if calledCount != 120 {
			// 100 turns to move
			// 20 turns to consume
			t.Errorf("Moving and Consuming needs 120 turns, %d", calledCount)
		}
	})
}
