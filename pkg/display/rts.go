package display

import "image"

type rts struct {
	buffer image.Image
}

func (d *rts) Render() image.Image {
	return d.buffer
}
