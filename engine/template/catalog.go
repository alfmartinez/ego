package template

import (
	"ego/engine/configuration"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	Data map[string]string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	if value, ok := d.Data[key]; ok {
		return "\x02" + value, true
	}
	return "", false
}

func RegisterCatalog() {

	var dicts map[string]*dictionary
	cfg := configuration.FromContext()
	err := cfg.UnmarshalKey("l10n", &dicts)
	if err != nil {
		panic(err)
	}

	var dictMap = make(map[string]catalog.Dictionary)

	for key, value := range dicts {
		dictMap[key] = value
	}

	defaultOption := catalog.Fallback(language.Make("en"))

	cat, err := catalog.NewFromMap(dictMap, defaultOption)
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}
