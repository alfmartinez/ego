package game

import (
	"ego/pkg/configuration"
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func CreateGame(filename string) (*Game, error) {
	dat, err := os.ReadFile("assets/configuration/" + filename)
	if err != nil {
		return nil, err
	}

	log.Print("Read configuration file")

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
	log.Print("Game created")
	return game, nil
}
