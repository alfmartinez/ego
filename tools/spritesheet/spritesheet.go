package spritesheet

import "log"

type Spritesheet interface {
	Load(path string)
	Export(path string)
}

func New() Spritesheet {
	return &spritesheet{}
}

type spritesheet struct {
}

func (s *spritesheet) Load(path string) {
	log.Printf("In  : %s", path)
}

func (s *spritesheet) Export(path string) {
	log.Printf("Out : %s", path)
}
