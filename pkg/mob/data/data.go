package data

import "ego/pkg/utils"

type Data struct {
	name     string
	position utils.Position
}

func CreateMobData(name string, position utils.Position) *Data {
	return &Data{name, position}
}

func (m Data) Name() string {
	return m.name
}

func (m Data) Position() utils.Position {
	return m.position
}

func (m Data) MoveTowards(destination utils.Position) bool {
	return false
}
