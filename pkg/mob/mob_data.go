package mob

import "ego/pkg/utils"

type MobData struct {
	name     string
	position utils.Position
}

func (m MobData) Name() string {
	return m.name
}

func (m MobData) Position() utils.Position {
	return m.position
}
