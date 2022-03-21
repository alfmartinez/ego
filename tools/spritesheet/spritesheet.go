package spritesheet

import (
	"image"
	"image/draw"
	"os"
	"path/filepath"
)

type Spritesheet interface {
	Load(path string)
	Export(path string)
}

func New() Spritesheet {
	return &spritesheet{}
}

type spritesheet struct {
	dest image.Image
}

func (s *spritesheet) Load(path string) {
	pwd, _ := os.Getwd()
	var files []string
	filepath.Walk(pwd+"/"+path+"/", func(p string, info os.FileInfo, err error) error {
		if info == nil {
			panic(err)
		}
		if !info.IsDir() {
			files = append(files, p)
		}
		return nil
	})

	var images []image.Image
	var rect = image.Rect(0, 0, 0, 0)
	for _, x := range files {
		img := LoadImage(x)
		rect = rect.Union(img.Bounds())
		images = append(images, img)
	}
	n := len(images)
	dest := image.NewNRGBA(image.Rect(0, 0, rect.Max.X*n, rect.Max.Y))
	dx := image.Pt(rect.Dx(), 0)
	for _, sprite := range images {
		draw.Draw(dest, rect, sprite, image.Pt(0, 0), draw.Src)
		rect = rect.Add(dx)
	}
	s.dest = dest
}

func (s *spritesheet) Export(path string) {
	err := SaveImage(path, s.dest)
	if err != nil {
		panic(err)
	}
}
