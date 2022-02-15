package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func Launch(filename string) (*Game, error) {
	dat, err := os.ReadFile("configuration/" + filename)
	if err != nil {
		return nil, errors.New("launch.configuration.missing")
	}

	var configuration Configuration

	if string(dat) == "" {
		return nil, errors.New("launch.configuration.empty")
	}

	err = yaml.Unmarshal(dat, &configuration)

	if err != nil {
		return nil, err // errors.New("launch.configuration.parsing")
	}

	if configuration.ModelVersion == "" {
		return nil, errors.New("launch.configuration.version.missing")
	}

	game := generateGame(configuration)

	return game, nil
}
