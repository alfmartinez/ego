package mob

import (
	"ego/behaviour"
	"ego/configuration"
)

type Template struct {
	config configuration.Template
}

func NewTemplate(config configuration.Template) *Template {
	template := &Template{config}
	return template
}

func (template *Template) Apply(mob *Mob) {
	for _, x := range template.config.Behaviours {
		behaviour := behaviour.New(x)
		mob.addBehaviour(behaviour)
	}
}
