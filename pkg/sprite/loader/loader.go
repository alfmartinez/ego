package loader

import (
	"image"
	_ "image/png"
	"os"
)

type Loader interface {
	GetSprite(string, uint) image.Image
}

func CreateSpriteLoader(name string) Loader {
	loaders := map[string]func() Loader{
		"on_demand": func() Loader {
			folder := "assets/sprites/"
			sheets := make(map[string]image.Image)
			return &onDemandLoader{folder, sheets}
		},
	}

	return loaders[name]()
}

func loadSpriteSheet(path string) image.Image {
	f, err := os.Open(path)
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
