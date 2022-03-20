package template

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"text/template"
)

func FromContext() *template.Template {
	return context.Get("template").(*template.Template)
}

func InitializeTemplates() {
	strings := configuration.FromContext().GetStringMapString("templates")
	tmpl := template.New("Main")
	for name, content := range strings {
		_, err := tmpl.New(name).Parse(content)
		if err != nil {
			panic(err)
		}
	}
	context.Set("template", tmpl)
}
