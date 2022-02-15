package main

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func Launch(filename string) (bool, error) {
	dat, err := os.ReadFile("configuration/" + filename)
	if err != nil {
		return false, errors.New("launch.configuration.missing")
	}

	var configuration Configuration

	if string(dat) == "" {
		return false, errors.New("launch.configuration.empty")
	}

	err = yaml.Unmarshal(dat, &configuration)

	if err != nil {
		return false, err // errors.New("launch.configuration.parsing")
	}

	if configuration.ModelVersion == "" {
		return false, errors.New("launch.configuration.version.missing")
	}

	log.Printf("Parsed file version %v", configuration.ModelVersion)

	return true, nil
}
