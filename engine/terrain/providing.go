package terrain

import (
	"ego/engine/motivator"
	"fmt"
)

func GetResourcesProviding(need motivator.Need) []Resource {
	switch need {
	case motivator.Health:
		return []Resource{Medicine}
	case motivator.Food:
		return []Resource{Food, Water}
	default:
		panic(fmt.Errorf("unknown need %d", need))
	}
}
