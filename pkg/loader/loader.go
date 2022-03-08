package loader

import (
	"embed"
	_ "embed"
	"image"
	_ "image/png"
)

//go:embed sprites/*
var content embed.FS

type Loader interface {
	Init()
	GetSprite(string, uint) image.Image
}

var loaderFactories = make(map[string]func() Loader)

func RegisterLoader(name string, f func() Loader) {
	loaderFactories[name] = f
}

func CreateSpriteLoader(name string) Loader {
	return loaderFactories[name]()
}

func loadSpriteSheet(path string) image.Image {
	f, err := content.Open("sprites/" + path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}
