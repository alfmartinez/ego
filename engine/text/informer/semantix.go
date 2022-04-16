package informer

type Semantix interface {
	GetStory() Story
}

func CreateRuleSemantix(debug bool) Semantix {
	return &semantix{
		debug: debug,
	}
}

type semantix struct {
	debug bool
	tests []string
}

func (s *semantix) GetStory() Story {
	return &story{}
}
