package terrain

import (
	"ego/pkg/movement"
	"image"
)

type Terrain interface {
	Tile(image.Point) Tile
	GetTile(movement.Positionnable) Tile
	FindClosest(movement.Positionnable, int, func(Tile) bool) []Tile
	AddSource(int, int, Resource, uint)
}

var terrainSingleton Terrain = nil

func GetTerrain() Terrain {
	return terrainSingleton
}
