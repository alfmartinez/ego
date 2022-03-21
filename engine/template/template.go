package template

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"reflect"
	"text/template"
)

func getLanguage() language.Tag {
	langKey := configuration.FromContext().GetString("lang")
	return language.Make(langKey)
}

func FromContext() *template.Template {
	return context.Get("template").(*template.Template)
}

func InitializeTemplates() {
	RegisterCatalog()

	var fns = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"translate": message.NewPrinter(getLanguage()).Sprintf,
	}

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
