package sequence

import (
	"ego/pkg/command"
	"ego/pkg/motivator"
	"ego/pkg/state"
)

func init() {
	RegisterSequenceFactory(Evaluate, CreateEvaluateCommand)
}

type EvaluateActor interface {
	SetState(state.StateType)
	TopNeed() motivator.Need
	After(command.Command)
}

func CreateEvaluateCommand(args ...interface{}) command.Command {
	actor, _ := args[0].(EvaluateActor)
	f := func() func() bool {
		return func() bool {
			actor.SetState(state.StateIdle)
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
