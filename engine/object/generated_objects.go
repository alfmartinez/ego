// DO NOT EDIT
package object

import (
	"github.com/alfmartinez/ego/engine/component"
)

type Camera struct {
	component.Transform2D
	component.Camera
}

type Mob struct {
	component.Transform2D
	component.SpriteRenderer
	component.RigidBody2D
	component.Collider2D
}
