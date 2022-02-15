package main

type Template struct {
	config TemplateConfig
}

func newTemplate(config TemplateConfig) *Template {
	template := &Template{config}
	return template
}

func (template *Template) apply(mob *Mob) {
	for _, x := range template.config.Behaviours {
		mob.addBehaviour(x)
	}
}
