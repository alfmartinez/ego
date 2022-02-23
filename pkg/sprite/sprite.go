package sprite

type Sprite interface {
	Path() string
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
