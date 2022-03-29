package text

type Story interface {
	Test()
}

func CreateStory(path string) Story {
	ast := ParseFile(path)
	semantix := CreateSemantix()
	return semantix.BuildStory(ast)
}

type story struct{}

func (s *story) Test() {

}
