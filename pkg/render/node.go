package render

import (
	"ego/pkg/display"
	"ego/pkg/utils"
	"image"
)

type RenderNode interface {
	utils.Attributes
	utils.Tree
	Apply(func(RenderNode))
	SetValue(RenderableObject)
	Value() RenderableObject
	Display() display.Displayable
	CreateFolder(string) RenderNode
	CreateObject(RenderableObject)
}

type RenderableObject interface {
	Path() string
	Size() uint
	Position() image.Point
}

func CreateRenderNode() RenderNode {
	return &renderNode{}
}

type renderNode struct {
	utils.Tree
	utils.Attributes
	value RenderableObject
}

func (n *renderNode) SetValue(value RenderableObject) {
	n.value = value
}

func (n *renderNode) Value() RenderableObject {
	return n.value
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
	folder.SetAttribute("name", name)
	folder.SetAttribute("type", "folder")
	n.AddChild(folder)
	return folder
}

func (n *renderNode) CreateObject(o RenderableObject) {
	node := CreateRenderNode()
	node.SetAttribute("type", "node")
	node.SetValue(o)
	n.AddChild(node)
}
