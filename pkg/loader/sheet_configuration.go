package loader

import (
	"embed"

	"gopkg.in/yaml.v2"
)

//go:embed config/sheets.yml
var configContent embed.FS

func loadConfiguration() SheetConfiguration {
	dat, err := configContent.ReadFile("config/sheets.yml")
	if err != nil {
		panic(err)
	}

	var config SheetConfiguration

	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		panic(err)
	}
	return config
}
