package mobile

import "ego/engine/physics/matrix"

type Mobile interface {
	GetMatrix() matrix.M
}
