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
	"article": func(x string) string {
		var vowelsLike = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true, 'h': true}
		var article = "a "
		if _, ok := vowelsLike[x[0]]; ok {
			article = "an "
		}
		return article + x
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
