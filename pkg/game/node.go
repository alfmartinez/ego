package game

import (
	"ego/pkg/object"
	"ego/pkg/utils"
)

type SceneNode interface {
	utils.Attributes
	utils.Tree
	Value() object.GameObject
	SetValue(object.GameObject)
	Update()
	CreateFolder(name string) SceneNode
	AddObject(object.GameObject) SceneNode
}

func CreateSceneNode() SceneNode {
	return &sceneNode{
		Tree:       utils.CreateTree(),
		Attributes: utils.CreateAttributes(),
	}
}

type sceneNode struct {
	utils.Tree
	utils.Attributes
	value object.GameObject
}

func (s *sceneNode) Value() object.GameObject {
	return s.value
}

func (s *sceneNode) SetValue(o object.GameObject) {
	s.value = o
}

func (s *sceneNode) Update() {
	if s.value != nil {
		s.value.Update()
	}
}

func (s *sceneNode) CreateFolder(name string) SceneNode {
	folder := CreateSceneNode()
	folder.SetAttribute("name", name)
	folder.SetAttribute("type", "folder")
	s.AddChild(folder)
	return folder
}

func (s *sceneNode) AddObject(o object.GameObject) SceneNode {
	node := CreateSceneNode()
	node.SetAttribute("type", "object")
	node.SetValue(o)
	s.AddChild(node)
	return node
}
