package object

import (
	"ego/pkg/command"
	"ego/pkg/motivator"
	"ego/pkg/state"
	"ego/pkg/terrain"
)

func CreateEvaluateCommand(actor StateMob) command.Command {
	f := func() func() bool {
		var stateChanged bool = false
		return func() bool {
			if !stateChanged {
				actor.SetState(state.StateIdle)
				stateChanged = true
			}
			need := actor.TopNeed()
			if need != motivator.None {
				actor.After(CreateSeekAndUseCommand(actor, actor.TopNeed()))
				return true
			}
			return false
		}
	}()
	return command.CreateCommand(f)
}

func CreateSeekCommand(actor StateMob, validate func(terrain.Tile) bool) command.Command {
	f := func() func() bool {
		var stateChanged bool = false
		grid := terrain.GetTerrain()
		return func() bool {
			if !stateChanged {
				actor.SetState(state.StateIdle)
				stateChanged = true
			}
			grid.FindClosest(actor, 1, validate)
			return true
		}
	}()
	return command.CreateCommand(f)
}

func CreateSeekAndUseCommand(actor StateMob, need motivator.Need) command.Command {
	resource := terrain.GetResourcesProviding(need)
	var foundResource terrain.Resource
	var foundTile terrain.Tile

	commandStream := command.CreateCommandStream()
	seekCommand := func() func() bool {
		var stateChanged bool = false
		grid := terrain.GetTerrain()
		return func() bool {
			if !stateChanged {
				actor.SetState(state.StateIdle)
				stateChanged = true
			}
			grid.FindClosest(actor, 1, func(t terrain.Tile) bool {
				for _, r := range resource {
					if t.HasResource(r) {
						foundResource = r
						foundTile = t
						return true
					}
				}
				return false
			})
			return true
		}
	}()
	moveCommand := func() func() bool {
		var stateChanged bool = false
		return func() bool {
			if foundTile == nil {
				commandStream.Abort()
				return true
			}
			if !stateChanged {
				actor.SetState(state.StateMove)
				stateChanged = true
			}
			return actor.MoveForward(foundTile)
		}
	}()
	consumeCommand := func() func() bool {
		var stateChanged bool = false
		var count int = 20
		return func() bool {
			if !stateChanged {
				actor.SetState(state.StateIdle)
				stateChanged = true
			}
			count--
			if count == 0 {
				foundTile.Consume(foundResource)
				actor.Provide(need, 1, 10)
				return true
			}
			return false
		}
	}()
	commandStream.After(command.CreateCommand(seekCommand))
	commandStream.After(command.CreateCommand(moveCommand))
	commandStream.After(command.CreateCommand(consumeCommand))
	return commandStream

}
