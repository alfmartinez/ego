package text

type Story interface {
	Test()
}

func CreateStory(path string) Story {
	return &story{}
}

type story struct{}

func (s *story) Test() {

}
