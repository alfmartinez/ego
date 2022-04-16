package informer

import "io"

type Story interface {
	Start()
	SetWriter(io.Writer)
	Test()
}

type story struct {
	writer io.Writer
}

func (s *story) Start() {}
func (s *story) SetWriter(writer io.Writer) {
	s.writer = writer
}

func (s *story) Test() {

}
