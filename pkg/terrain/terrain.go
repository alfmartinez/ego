package terrain

import "ego/pkg/utils"

type Terrain interface {
	GetTile(utils.Position)
}
