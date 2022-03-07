package game

import (
	"ego/pkg/renderer"
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
	s.root.Apply(func(node SceneNode) {
		node.Update()
	})
}

func (s *scene) Render() renderer.RenderTree {
	renderTree := renderer.CreateRenderTree()
	return renderTree
}
