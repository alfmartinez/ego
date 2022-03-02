package loader

import (
	"image"
	_ "image/png"
	"os"
)

type Loader interface {
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
