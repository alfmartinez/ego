package physics

import (
	"ego/pkg/movement"
	"image"
)

type CollisionType int

const (
	COLLISION_NONE CollisionType = iota
	COLLISION_TOP
	COLLISION_BOTTOM
	COLLISION_LEFT
	COLLISION_RIGHT
)

type Body interface {
	Collider
}

type Collider interface {
	Hitbox() image.Rectangle
	IsIgnored() bool
}

type Mobile interface {
	IsHit(Collider) CollisionType
	Position() movement.Location
	SetPosition(movement.Location)
	Speed() movement.Speed
	SetSpeed(movement.Speed)
	Acceleration() movement.Acceleration
	SetAcceleration(movement.Acceleration)
	Land()
	Fall()
}
