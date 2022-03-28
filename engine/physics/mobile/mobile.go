package mobile

import (
	"github.com/alfmartinez/ego/engine/physics/matrix"
	"image"
)

type Mobile interface {
	GetMatrix() matrix.M
	SetMatrix(matrix.M)
	IsHit(image.Rectangle) bool
}

type Colliding interface {
	Hitbox() image.Rectangle
	IsSolid() bool
}
