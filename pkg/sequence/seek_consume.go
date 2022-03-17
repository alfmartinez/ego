package sequence

import (
	"ego/pkg/command"
	"ego/pkg/context"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/state"
	"ego/pkg/terrain"
)

func init() {
	RegisterSequenceFactory(Seek, CreateSeekAndUseCommand)
}

type SeekerActor interface {
	SetState(state.StateType)
	movement.Movement
	Provide(motivator.Need, int, int)
}

func CreateSeekAndUseCommand(args ...interface{}) command.Command {
	actor := args[0].(SeekerActor)
	need := args[1].(motivator.Need)
	resource := terrain.GetResourcesProviding(need)
	var foundResource terrain.Resource
	var foundTile terrain.Tile

	commandStream := command.CreateCommandStream()
	seekCommand := func() func() bool {
		grid := context.GetContext().Get("terrain").(terrain.Terrain)
		return func() bool {
			actor.SetState(state.StateIdle)
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
			if foundTile == nil {
				commandStream.Abort()
				return true
			}
			return true
		}
	}()
	moveCommand := func() func() bool {
		return func() bool {
			actor.SetState(state.StateMove)
			return actor.MoveForward(foundTile)
		}
	}()
	consumeCommand := func() func() bool {
		var count int = 20
		return func() bool {
			actor.SetState(state.StateIdle)
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
