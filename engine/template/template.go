package template

import (
	"github.com/alfmartinez/ego/engine/configuration"
	"github.com/alfmartinez/ego/engine/context"
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

func InitializeTemplates(path string) {
	RegisterCatalog()

	var fns = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"translate": message.NewPrinter(getLanguage()).Sprintf,
	}

	tmpl, err := template.New("Main").Funcs(fns).ParseGlob(path)
	if err != nil {
		panic(err)
	}
	tmpl.Funcs(fns)

	context.Set("template", tmpl)
}
