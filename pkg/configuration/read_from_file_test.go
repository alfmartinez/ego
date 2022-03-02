package configuration

import "testing"

func TestReadConfigurationFromFileMissing(t *testing.T) {
	_, err := ReadConfigurationFromFile("missing.yml")
	if err == nil {
		t.Error("Should not read missing file")
	}
}

func TestReadConfigurationFromEmptyFile(t *testing.T) {
	_, err := ReadConfigurationFromFile("empty.yml")
	if err == nil {
		t.Error("Should not read empty file")
	}
}

func TestReadConfiguration(t *testing.T) {
	_, err := ReadConfigurationFromFile("game.yml")
	if err != nil {
		t.Errorf("Should not receive error, got %+v", err)
	}
}
