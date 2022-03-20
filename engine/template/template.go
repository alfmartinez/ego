package template

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"reflect"
	"text/template"
)

func FromContext() *template.Template {
	return context.Get("template").(*template.Template)
}

var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

func InitializeTemplates() {
	strings := configuration.FromContext().GetStringMapString("templates")
	tmpl := template.New("Main")
	for name, content := range strings {
		_, err := tmpl.New(name).Funcs(fns).Parse(content)
		if err != nil {
			panic(err)
		}
	}
	context.Set("template", tmpl)
}
