package spritesheet

import (
	"image"
	"image/png"
	_ "image/png"
	"os"
)

func SaveImage(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return png.Encode(f, img)
}
