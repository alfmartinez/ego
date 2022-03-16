package object

import (
	"ego/pkg/command"
	"ego/pkg/data"
	"ego/pkg/memory"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/sequence"
	"ego/pkg/sprite"
	"ego/pkg/state"
	"fmt"
	"image"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func init() {
	RegisterObjectFactory("Mob", CreateStateMob)
}

type StateMob interface {
	GameObject
	state.StateMachine
	memory.Memory
	data.Data
	movement.Movement
	motivator.NeedsCollection
	sprite.Sprite
	command.CommandStream
}

type stateMob struct {
	state.StateMachine
	memory.Memory
	data.Data
	movement.Movement
	sprite.Sprite
	motivator.NeedsCollection
	command.CommandStream
}

func CreateStateMob(key string) GameObject {
	vPath := fmt.Sprintf("mobs.%s", key)
	sub := viper.Sub(vPath)
	if sub == nil {
		panic(fmt.Errorf("cant find mob configuration, %s", vPath))
	}
	log.Printf("%+v", sub.AllKeys())
	name := sub.GetString("name")
	position := image.Pt(sub.GetInt("position.x"), sub.GetInt("position.y"))

	mobData := data.CreateData(name)
	mvmnt := movement.CreateMovement(position)
	memo := memory.CreateMemory()
	sprt := sprite.CreateSprite(sub.GetString("sprite"))
	needs := motivator.CreateNeedsCollection()
	for needK, value := range sub.GetStringMapString("needs") {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		needs.AddNeed(motivator.CreateNeed(needK), intValue)
	}
	sm := state.CreateStateMachine()
	stream := command.CreateCommandStream()

	return &stateMob{sm, memo, mobData, mvmnt, sprt, needs, stream}
}

func (m *stateMob) Update() {
	if m.Execute() {
		m.After(sequence.CreateSequence(sequence.Evaluate)(m))
	}
	m.StateMachine.DoUpdate(m)
}
