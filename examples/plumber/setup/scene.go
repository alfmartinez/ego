package setup

import (
	"github.com/alfmartinez/ego/engine/component"
	"github.com/alfmartinez/ego/engine/object"
	eScene "github.com/alfmartinez/ego/engine/scene"
)

var defaultScene = eScene.Scene{
	Label: "default",
	Objects: []object.GameObject{
		{
			Label: "Camera",
			Comps: []component.Component{
				&component.Transform2D{},
				&component.Camera{},
			},
		},
		{
			Label: "Character",
			Comps: []component.Component{
				&component.Transform2D{},
				&component.SpriteRenderer{},
				&component.Collider2D{},
				&component.RigidBody2D{},
			},
		},
	},
}

func DefaultScene() eScene.Scene {
	return defaultScene
}
