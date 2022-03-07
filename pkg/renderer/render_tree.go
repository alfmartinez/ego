package renderer

import "ego/pkg/utils"

type RenderTree interface {
	utils.Tree
	utils.Attributes
}

func CreateRenderTree() RenderTree {
	return &renderTree{
		Tree:       utils.CreateTree(),
		Attributes: utils.CreateAttributes(),
	}
}

type renderTree struct {
	utils.Tree
	utils.Attributes
}
