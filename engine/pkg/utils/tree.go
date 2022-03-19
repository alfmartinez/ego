package utils

type Tree interface {
	AddChild(Tree)
	AddChildren([]Tree)
	Children() []Tree
}

func CreateTree() Tree {
	return &tree{}
}

type tree struct {
	children []Tree
}

func (t *tree) AddChild(a Tree) {
	t.children = append(t.children, a)
}

func (t *tree) AddChildren(a []Tree) {
	t.children = append(t.children, a...)
}

func (t *tree) Children() []Tree {
	return t.children
}
