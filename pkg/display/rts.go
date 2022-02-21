package display

import (
	"ego/pkg/renderable"
	"image"
)

type rts struct {
	buffer image.Image
}

func (d *rts) Render() image.Image {
	return d.buffer
}

func (d *rts) AddObject(s renderable.Renderable) {

}
