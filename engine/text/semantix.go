package text

type Semantix interface {
	BuildStory(*Grammar) Story
}

func CreateSemantix() Semantix {
	return &semantix{}
}

type semantix struct {
}

func (s *semantix) BuildStory(ast *Grammar) Story {
	return &story{}
}
