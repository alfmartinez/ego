package main

import (
	"fmt"
	"testing"
)

func TestLaunchShouldReturnFalseIfConfigurationFileNotFound(t *testing.T) {
	result, error := Launch("missing.yml")
	if result != nil {
		t.Errorf("Launch should return false if configuration file does not exist")
	}
	if error.Error() != "launch.configuration.missing" {
		t.Errorf("Launch should return error code for missing, got %s", error.Error())
	}

}

func TestLaunchShouldReturnFalseIfConfigurationFileIsEmpty(t *testing.T) {
	result, error := Launch("empty.yml")
	if result != nil {
		t.Errorf("Launch should return false if configuration file does not exist : %s", error.Error())
	}
	if error.Error() != "launch.configuration.empty" {
		t.Errorf("Launch should return error code for empty, got %s", error.Error())
	}
}

func TestLaunchShouldReturnFalseIfNoVersionIsGiving(t *testing.T) {
	result, error := Launch("noversion.yml")
	if result != nil {
		t.Errorf("Launch should return false if configuration file does not exist : %s", error.Error())
	}
	if error.Error() != "launch.configuration.version.missing" {
		t.Errorf("Launch should return error code for empty, got %s", error.Error())
	}
}

func TestLaunchShouldReturnTrueIfConfigurationFileFoundAndParsed(t *testing.T) {
	result, error := Launch("simple.yml")
	if result == nil {
		t.Errorf("Launch should return true if configuration file parsed : %s", error.Error())
	}
	if error != nil {
		t.Errorf("Launch should return no error code for missing, got %s", error.Error())
	}

}

func TestLaunchShouldReturnTrueIfConfigurationFileFoundAndParsedWithBasic(t *testing.T) {
	result, error := Launch("basic.yml")
	if result == nil {
		t.Errorf("Launch should return true if configuration file parsed : %s", error.Error())
	}
	if error != nil {
		t.Errorf("Launch should return no error code for missing, got %s", error.Error())
	}
	fmt.Printf("Game %+v", result)

}
