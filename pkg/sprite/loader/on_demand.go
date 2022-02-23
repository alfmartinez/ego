package loader

import (
	"image"
	"log"
)

type onDemandLoader struct {
	folder string
	sheets map[string]image.Image
}

func (l *onDemandLoader) GetSprite(name string) image.Image {
	log.Printf("Loader %+v", l.sheets)
	src, ok := l.sheets[name]
	if !ok {
		src = loadSpriteSheet(l.folder + name)
		l.sheets[name] = src
	}
	return src
}
