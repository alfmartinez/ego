package spritesheet

import (
	"image"
	"image/png"
	_ "image/png"
	"os"
)

func SaveImage(path string, img image.Image) error {
	pwd, _ := os.Getwd()
	f, err := os.Create(pwd + path)
	if err != nil {
		panic(err)
	}
	return png.Encode(f, img)
}
