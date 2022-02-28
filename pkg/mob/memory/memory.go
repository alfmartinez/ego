package memory

import (
	"ego/pkg/mob/movement"
	"ego/pkg/utils"
)

type Memory interface{}

type memory struct {
	places    map[utils.Position]PlaceMemory
	interests []movement.Positionnable
}

func CreateMemory() Memory {
	places := make(map[utils.Position]PlaceMemory)
	var interests []movement.Positionnable
	return &memory{places, interests}
}
