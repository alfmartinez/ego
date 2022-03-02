package game

import (
	"ego/pkg/configuration"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func init() {
	RegisterGameFactory("file", fromFile)
}

func fromFile() (Game, error) {
	filename := "game.yml"
	dat, err := os.ReadFile("assets/configuration/" + filename)
	if err != nil {
		return nil, err
	}

	var configuration configuration.Configuration

	if string(dat) == "" {
		return nil, errors.New("launch.configuration.empty")
	}

	err = yaml.Unmarshal(dat, &configuration)

	if err != nil {
		return nil, err
	}

	if configuration.ModelVersion == "" {
		return nil, errors.New("launch.configuration.version.missing")
	}

	game := generateGame(configuration)
	return game, nil
}
