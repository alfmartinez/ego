package render

type RenderTree interface {
	Root() RenderNode
	Apply(func(RenderNode))
}

type renderTree struct {
	root RenderNode
}

func CreateRenderTree() RenderTree {
	return &renderTree{
		root: CreateRenderNode(),
	}
}

func (t *renderTree) Root() RenderNode {
	return t.root
}

func (t *renderTree) Apply(f func(n RenderNode)) {
	t.root.Apply(f)
}
