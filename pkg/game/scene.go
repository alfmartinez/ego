package game

import (
	"ego/pkg/renderer"
	"ego/pkg/utils"
)

type Scene interface {
	Update()
	Render() renderer.RenderTree
	Root() SceneNode
}

type scene struct {
	root SceneNode
}

func CreateScene() Scene {
	return &scene{
		root: CreateSceneNode(),
	}
}

func (s *scene) Root() SceneNode {
	return s.root
}

func (s *scene) Update() {
	s.root.Apply(func(n utils.Tree) {
		node := n.(SceneNode)
		node.Update()
	})
}

func (s *scene) Render() renderer.RenderTree {
	renderTree := renderer.CreateRenderTree()
	return renderTree
}
