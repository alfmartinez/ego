package object

import (
	"ego/pkg/command"
	"ego/pkg/context"
	"ego/pkg/data"
	"ego/pkg/memory"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/observer"
	"ego/pkg/physics"
	"ego/pkg/renderer"
	"ego/pkg/sprite"
	"ego/pkg/state"
	"fmt"
	"image"
	"time"

	"github.com/spf13/viper"
)

func init() {
	RegisterObjectFactory("Mob", CreateStateMob)
}

type StateMob interface {
	physics.Collider
	physics.Mobile
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
	Position movement.Location
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

	m := &stateMob{sm, memo, data, mvmnt, sprt, needs, stream}
	m.Fall()
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
	m.Execute()
	m.StateMachine.DoUpdate(m, dt)
}

func (m *stateMob) render() {
	r := context.GetContext().Get("renderer").(renderer.Renderer)
	r.Render(m)
}

func (m *stateMob) IsMobile() bool {
	return true
}

func (m *stateMob) Hitbox() image.Rectangle {
	return image.Rect(0, 0, int(m.Size()), int(m.Size())).Add(m.Position().Point())
}

func (m *stateMob) IsHit(collider physics.Collider) physics.CollisionType {
	intersect := m.Hitbox().Intersect(collider.Hitbox())
	if intersect != image.Rect(0, 0, 0, 0) {
		if intersect.Min.Y < m.Hitbox().Max.Y {
			return physics.COLLISION_BOTTOM
		}
		if intersect.Max.Y > m.Hitbox().Min.Y {
			return physics.COLLISION_TOP
		}
		if intersect.Min.X < m.Hitbox().Max.X {
			return physics.COLLISION_RIGHT
		}
		if intersect.Max.X > m.Hitbox().Min.X {
			return physics.COLLISION_LEFT
		}
	}
	return physics.COLLISION_NONE
}

func (m *stateMob) IsIgnored() bool {
	return false
}
