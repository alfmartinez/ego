package object

import (
	"ego/pkg/command"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/terrain"
)

func CreateEvaluateCommand(m StateMob) command.Command {
	f := func() func() bool {
		actor := m
		return func() bool {
			need := m.TopNeed()
			if need != motivator.None {
				actor.After(CreateSeekAndUseCommand(m, m.TopNeed()))
				return true
			}
			return false
		}
	}()
	return command.CreateCommand(f)
}

func CreateSeekCommand(m StateMob, validate func(terrain.Tile) bool, callback func([]terrain.Tile)) command.Command {
	f := func() func() bool {
		actor := m
		grid := terrain.GetTerrain()
		return func() bool {
			tiles := grid.FindClosest(actor, 1, validate)
			callback(tiles)
			return true
		}
	}()
	return command.CreateCommand(f)
}

func CreateMoveCommand(m StateMob, dest movement.Positionnable) command.Command {
	f := func() func() bool {
		actor := m
		destination := dest
		return func() bool {
			return actor.MoveTowards(destination)
		}
	}()
	return command.CreateCommand(f)
}

func CreateConsumeCommand(m StateMob, tile terrain.Tile, need motivator.Need, r terrain.Resource) command.Command {
	f := func() func() bool {
		return func() bool {
			tile.Consume(r)
			m.Provide(need, 1, 10)
			return true
		}
	}()
	return command.CreateCommand(f)
}

func CreateSeekAndUseCommand(m StateMob, need motivator.Need) command.Command {
	resource := terrain.GetResourcesProviding(need)
	var foundResource terrain.Resource
	return CreateSeekCommand(m, func(t terrain.Tile) bool {
		for _, r := range resource {
			if t.HasResource(r) {
				foundResource = r
				return true
			}
		}
		return false
	}, func(tiles []terrain.Tile) {
		if len(tiles) > 0 {
			tile := tiles[0]
			m.After(CreateMoveCommand(m, tile))
			m.After(CreateConsumeCommand(m, tile, need, foundResource))
		}
	})
}
