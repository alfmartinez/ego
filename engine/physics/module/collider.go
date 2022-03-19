package module

import (
	"ego/engine/physics/mobile"
	"ego/engine/slices"
)

func init() {
	RegisterModuleFactory("collider", func() Module {
		return &collider{}
	})
}

type collider struct {
	mobiles []mobile.Mobile
	solidBs []mobile.Colliding
}

// Init implements Module
func (m *collider) Init() {
	m.mobiles = make([]mobile.Mobile, 0)
	m.solidBs = make([]mobile.Colliding, 0)
}

// Add implements Module
func (m *collider) Add(i any) {
	if o, ok := i.(mobile.Mobile); ok {
		m.mobiles = append(m.mobiles, o)
	}
	if o, ok := i.(mobile.Colliding); ok && o.IsSolid() {
		m.solidBs = append(m.solidBs, o)
	}
}

// Advance implements Module
func (m *collider) Advance(dt float64) {
	slices.Apply(m.mobiles, func(o mobile.Mobile) {
		slices.Apply(m.solidBs, func(c mobile.Colliding) {
			if o.(mobile.Colliding) != c && o.IsHit(c.Hitbox()) {
				matrix := o.GetMatrix()
				matrix.S.Y = 0
				matrix.A.Y = 0
				o.SetMatrix(matrix)
			}
		})
	})
}
