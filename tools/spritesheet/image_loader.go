package spritesheet

import (
	"image"
	_ "image/png"
	"log"
	"os"
)

func LoadImage(path string) image.Image {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Printf("Path : %s", path)
		panic(err)
	}
	return img
}
