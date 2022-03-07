package renderer

import "ego/pkg/utils"

type RenderTree interface {
	Apply(func(RenderTree))
}

func CreateRenderTree() RenderTree {
	return &renderTree{utils.CreateTree(), utils.CreateAttributes()}
}

type renderTree struct {
	utils.Tree
	utils.Attributes
}

func (t *renderTree) Apply(f func(RenderTree)) {
	f(t)
	for _, x := range t.Children() {
		c := x.(RenderTree)
		c.Apply(f)
	}
}
