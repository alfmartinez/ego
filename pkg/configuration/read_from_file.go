package configuration

import (
	"embed"
	"errors"

	"gopkg.in/yaml.v2"
)

//go:embed data/*
var content embed.FS

func ReadConfigurationFromFile(filename string) (Configuration, error) {
	dat, err := content.ReadFile("data/" + filename)
	if err != nil {
		return Configuration{}, err
	}

	var config Configuration

	if string(dat) == "" {
		return Configuration{}, errors.New("launch.configuration.empty")
	}

	err = yaml.Unmarshal(dat, &config)

	if err != nil {
		return Configuration{}, err
	}

	if config.ModelVersion == "" {
		return Configuration{}, errors.New("launch.configuration.version.missing")
	}

	return config, nil
}
