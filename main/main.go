package main

import (
	"ego/configuration"
	"ego/game"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func Launch(filename string) (*game.Game, error) {
	dat, err := os.ReadFile("configuration/" + filename)
	if err != nil {
		return nil, errors.New("launch.configuration.missing")
	}

	var configuration configuration.Configuration

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

	game := game.GenerateGame(configuration)

	return game, nil
}
