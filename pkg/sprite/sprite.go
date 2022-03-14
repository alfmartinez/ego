package sprite

import (
	"strconv"
	"strings"
)

type Sprite interface {
	Frame(x, y int)
	Path() string
	Size() uint
}

type Frame struct {
	X, Y int
}

func (f Frame) buildPathArray(a string) [3]string {
	b := strconv.Itoa(f.X)
	c := strconv.Itoa(f.Y)
	return [3]string{a, b, c}
}

func NewFrame(x, y int) Frame {
	return Frame{x, y}
}

type sprite struct {
	sheet string
	frame Frame
	size  uint
}

func CreateSprite(path string) Sprite {
	return &sprite{path, NewFrame(0, 0), 50}
}

func (s *sprite) Frame(x, y int) {
	s.frame = NewFrame(x, y)
}

func (s *sprite) Path() string {
	parts := s.frame.buildPathArray(s.sheet)
	return strings.Join(parts[:], ":")
}

func (s *sprite) Size() uint {
	return 512
}
