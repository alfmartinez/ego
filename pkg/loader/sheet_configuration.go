package loader

import (
	"embed"

	"gopkg.in/yaml.v2"
)

//go:embed config/sheets.yml
var configContent embed.FS

func loadConfiguration() SheetConfiguration {
	dat, _ := configContent.ReadFile("config/sheets.yml")

	var config SheetConfiguration
	yaml.Unmarshal(dat, &config)

	return config
}
