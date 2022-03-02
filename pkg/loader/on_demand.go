package loader

import (
	"image"
	"strconv"

	"github.com/nfnt/resize"
)

func init() {
	RegisterLoader("on_demand", func() Loader {
		folder := "assets/sprites/"
		sheets := make(map[string]image.Image)
		return &onDemandLoader{folder, sheets}
	})
}

type onDemandLoader struct {
	folder string
	sheets map[string]image.Image
}

func (l *onDemandLoader) GetSprite(name string, size uint) image.Image {
	key := name + "_" + strconv.Itoa(int(size))
	src, ok := l.sheets[key]
	if !ok {
		src, ok := l.sheets[name]
		if !ok {
			src = loadSpriteSheet(l.folder + name)
			l.sheets[name] = src
		}
		resized := resize.Resize(size, 0, src, resize.Lanczos2)
		l.sheets[key] = resized
		return resized
	}
	return src
}
