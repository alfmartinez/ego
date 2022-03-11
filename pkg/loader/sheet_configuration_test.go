package loader

import "testing"

func TestSheetConfiguration(t *testing.T) {
	config := loadConfiguration()
	if len(config.Sheets) == 0 {
		t.Errorf("should return sheets configuration")
	}
}
