package loader

import (
	"image"
	"strconv"

	"github.com/nfnt/resize"
)

func init() {
	RegisterLoader("on_demand", func() Loader {
		sheets := make(map[string]image.Image)
		return &onDemandLoader{sheets}
	})
}

type onDemandLoader struct {
	sheets map[string]image.Image
}

func (l *onDemandLoader) Init() {}

func (l *onDemandLoader) GetSprite(name string, size uint) image.Image {
	key := name + "_" + strconv.Itoa(int(size))
	src, ok := l.sheets[key]
	if !ok {
		src, ok := l.sheets[name]
		if !ok {
			src = loadSpriteSheet(name)
			l.sheets[name] = src
		}
		resized := resize.Resize(size, 0, src, resize.Lanczos2)
		l.sheets[key] = resized
		return resized
	}
	return src
}
