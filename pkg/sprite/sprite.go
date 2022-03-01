package sprite

type Sprite interface {
	Path() string
	Size() uint
	Multiplicator() int
}

type sprite struct {
	path string
}

func CreateSprite(path string) Sprite {
	return &sprite{path}
}

func (s *sprite) Path() string {
	return s.path
}

func (s *sprite) Size() uint {
	return 100
}

func (s *sprite) Multiplicator() int {
	return 1
}
