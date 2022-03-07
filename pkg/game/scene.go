package game

import (
	"ego/pkg/render"
)

type Scene interface {
	Update()
	Render() render.RenderTree
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

func (s *scene) Render() render.RenderTree {
	renderTree := render.CreateRenderTree()
	folder := renderTree.Root().CreateFolder("render")
	s.root.Apply(func(node SceneNode) {
		o := node.Value()
		if o != nil {
			folder.CreateObject(o)
		}
	})
	return renderTree
}
