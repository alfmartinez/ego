package loader

import (
	"image"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func init() {
	RegisterLoader("pre_load", func() Loader {
		imgs := make(map[index]image.Image)
		return &preLoad{imgs}
	})
}

type SheetConfiguration struct {
	Sheets []struct {
		Label string
		Path  string
		Rect  struct {
			X int
			Y int
		}
		Sizes []uint
	}
}

type index struct {
	label string
	x, y  int
	size  uint
}

type CropableImage interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

type preLoad struct {
	imgs map[index]image.Image
}

func (l *preLoad) Init() {

	config := loadConfiguration()
	for _, sheet := range config.Sheets {
		label := sheet.Label
		img := loadSpriteSheet("sheets/" + sheet.Path)
		columns, lines := divideRect(img.Bounds(), image.Rect(0, 0, sheet.Rect.X, sheet.Rect.Y))
		rect := image.Rect(0, 0, sheet.Rect.X, sheet.Rect.Y)
		dX := image.Pt(rect.Dx(), 0)
		dY := image.Pt(0, rect.Dy())
		var subImage image.Image
		line := rect
		for j := 0; j < lines; j++ {
			col := line
			for i := 0; i < columns; i++ {
				subImage = img.(CropableImage).SubImage(col)
				for _, size := range sheet.Sizes {
					l.imgs[index{label, i, j, size}] = resize.Resize(uint(size), 0, subImage, resize.Lanczos2)
				}
				col = col.Add(dX)
			}
			line = rect.Add(dY)
		}
	}
}

func (l *preLoad) GetSprite(name string, size uint) image.Image {
	parts := strings.Split(name, ":")
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	idx := index{parts[0], x, y, size}
	return l.imgs[idx]
}

func divideRect(src image.Rectangle, tile image.Rectangle) (int, int) {
	return src.Dx() / tile.Dx(), src.Dy() / tile.Dy()
}
