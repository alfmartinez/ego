package main

import "testing"

func TestLaunchShouldReturnFalseIfConfigurationFileNotFound(t *testing.T) {
	result, error := Launch("missing.yml")
	if result != false {
		t.Errorf("Launch should return false if configuration file does not exist")
	}
	if error.Error() != "launch.configuration.missing" {
		t.Errorf("Launch should return error code for missing, got %s", error.Error())
	}

}

func TestLaunchShouldReturnFalseIfConfigurationFileIsEmpty(t *testing.T) {
	result, error := Launch("empty.yml")
	if result != false {
		t.Errorf("Launch should return false if configuration file does not exist : %s", error.Error())
	}
	if error.Error() != "launch.configuration.modelVersion.missing" {
		t.Errorf("Launch should return error code for empty, got %s", error.Error())
	}
}

func TestLaunchShouldReturnTrueIfConfigurationFileFoundAndParsed(t *testing.T) {
	result, error := Launch("simple.yml")
	if result != true {
		t.Errorf("Launch should return true if configuration file parsed : %s", error.Error())
	}
	if error != nil {
		t.Errorf("Launch should return no error code for missing, got %s", error.Error())
	}

}
