package loader

import (
	"image"
)

type onDemandLoader struct {
	folder string
	sheets map[string]image.Image
}

func (l *onDemandLoader) GetSprite(name string) image.Image {
	src, ok := l.sheets[name]
	if !ok {
		src = loadSpriteSheet(l.folder + name)
		l.sheets[name] = src
	}
	return src
}
