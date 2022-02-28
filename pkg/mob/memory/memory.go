package memory

import (
	"ego/pkg/mob/movement"
	"image"
)

type Memory interface{}

type memory struct {
	places    map[image.Point]PlaceMemory
	interests []movement.Positionnable
}

func CreateMemory() Memory {
	places := make(map[image.Point]PlaceMemory)
	var interests []movement.Positionnable
	return &memory{places, interests}
}
