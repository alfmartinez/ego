package object

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/layer"
	"github.com/alfmartinez/ego/engine/movement"
	"github.com/alfmartinez/ego/engine/object"
	"github.com/alfmartinez/ego/engine/observer"
	"github.com/alfmartinez/ego/engine/physics"
	"github.com/alfmartinez/ego/engine/renderer"
	"github.com/alfmartinez/ego/engine/sprite"
	"github.com/alfmartinez/ego/engine/state"
	"image"
	"time"

	"github.com/spf13/viper"
)

func Register() {
	object.RegisterObjectFactory("Mob", CreateStateMob)
}

type StateMob interface {
	object.GameObject
	state.StateMachine
	movement.Movement
	sprite.Sprite
	layer.Layered
}

type stateMob struct {
	state.StateMachine
	movement.Movement
	sprite.Sprite
	layer.Layered
}

type MobData struct {
	Position movement.Location
	States   string
	Sprite   struct {
		Path string
		Size uint
	}
}

func CreateStateMob(key string) object.GameObject {
	var mobData MobData
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	vPath := fmt.Sprintf("mobs.%s", key)
	err := viper.UnmarshalKey(vPath, &mobData)
	if err != nil {
		panic(err)
	}
	position := mobData.Position
	mvmnt := movement.CreateMovement(position)
	sprt := sprite.CreateSprite(mobData.Sprite.Path, mobData.Sprite.Size)
	sm := state.CreateStateMachine()

	l := layer.CreateLayered()
	l.SetLayer(layer.FOREGROUND)

	m := &stateMob{sm, mvmnt, sprt, l}
	states := state.CreateStates(mobData.States, m)
	sm.SetStates(states)
	return m
}

func (m *stateMob) OnNotify(e observer.Event) {
	switch e.Type() {
	case observer.UPDATE:
		dt := e.Data()[0].(time.Duration)
		m.update(dt)
	case observer.RENDER:
		m.render()
	case observer.PHYSICS:
		p := physics.FromContext()
		p.Add(m)
	}
}

func (m *stateMob) update(dt time.Duration) {
	m.StateMachine.DoUpdate(m, dt)
}

func (m *stateMob) render() {
	r := renderer.FromContext()
	r.Render(m)
}

func (m *stateMob) Hitbox() image.Rectangle {
	rect := image.Rect(0, 0, int(m.Size()), int(m.Size())).Inset(int(m.Size()) / 4)
	pos := image.Point{int(m.Position().X), int(m.Position().Y)}
	return rect.Add(pos)
}

func (m *stateMob) IsHit(rect image.Rectangle) bool {
	return m.Hitbox().Intersect(rect) != image.Rect(0, 0, 0, 0)
}

func (m *stateMob) IsSolid() bool {
	return true
}
