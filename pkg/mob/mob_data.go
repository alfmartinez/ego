package mob

import "ego/pkg/utils"

type MobData struct {
	name     string
	position utils.Position
}

func CreateMobData(name string, position utils.Position) *MobData {
	return &MobData{name, position}
}

func (m MobData) Name() string {
	return m.name
}

func (m MobData) Position() utils.Position {
	return m.position
}
