package render

import (
	"ego/pkg/display"
	"ego/pkg/utils"
)

type RenderNode interface {
	utils.Tree
	Apply(func(RenderNode))
	Display() display.Displayable
	CreateFolder(string) RenderNode
	CreateObject(interface{})
}

func CreateRenderNode() RenderNode {
	return &renderNode{
		Tree:  utils.CreateTree(),
		value: nil,
	}
}

func CreateObjectNode(v display.Displayable) RenderNode {
	return &renderNode{
		Tree:  utils.CreateTree(),
		value: v,
	}
}

type renderNode struct {
	utils.Tree
	value display.Displayable
}

func (n *renderNode) Display() display.Displayable {
	return n.value
}

func (n *renderNode) Apply(f func(RenderNode)) {
	f(n)
	for _, x := range n.Children() {
		node := x.(RenderNode)
		node.Apply(f)
	}
}

func (n *renderNode) CreateFolder(name string) RenderNode {
	folder := CreateRenderNode()
	n.AddChild(folder)
	return folder
}

func (n *renderNode) CreateObject(o interface{}) {
	node := CreateObjectNode(ConvertObjectToDisplayable(o))
	n.AddChild(node)
}
