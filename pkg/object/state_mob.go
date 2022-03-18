package object

import (
	"ego/pkg/command"
	"ego/pkg/context"
	"ego/pkg/data"
	"ego/pkg/memory"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/observer"
	"ego/pkg/renderer"
	"ego/pkg/sprite"
	"ego/pkg/state"
	"fmt"
	"image"

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

type MobData struct {
	Name     string
	Position image.Point
	Sprite   struct {
		Path string
		Size uint
	}
	Needs map[string]int
}

func CreateStateMob(key string) GameObject {
	var mobData MobData
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	vPath := fmt.Sprintf("mobs.%s", key)
	err := viper.UnmarshalKey(vPath, &mobData)
	if err != nil {
		panic(err)
	}
	name := mobData.Name
	position := mobData.Position
	data := data.CreateData(name)
	mvmnt := movement.CreateMovement(position)
	memo := memory.CreateMemory()
	sprt := sprite.CreateSprite(mobData.Sprite.Path, mobData.Sprite.Size)
	needs := motivator.CreateNeedsCollection()
	for needK, value := range mobData.Needs {
		needs.AddNeed(motivator.CreateNeed(needK), value)
	}
	sm := state.CreateStateMachine()
	stream := command.CreateCommandStream()

	return &stateMob{sm, memo, data, mvmnt, sprt, needs, stream}
}

func (m *stateMob) OnNotify(e observer.Event) {
	switch e.Type() {
	case observer.UPDATE:
		m.update()
	case observer.RENDER:
		m.render()
	}
}

func (m *stateMob) update() {
	m.Execute()
	m.StateMachine.DoUpdate(m)
}

func (m *stateMob) render() {
	r := context.GetContext().Get("renderer").(renderer.Renderer)
	r.Render(m)
}
