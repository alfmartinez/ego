package text

import (
	"fmt"
	"io"
)

type Story interface {
	Test()
	Start() chan string
	SetWriter(io.Writer)
}

func CreateStory(path string) Story {
	ast := ParseFile(path)
	semantix := CreateSemantix()
	return semantix.BuildStory(ast)
}

type storyState struct {
	titleDisplayed bool
}

type story struct {
	state    storyState
	title    string
	writer   io.Writer
	rooms    map[string]Room
	startAt  string
	tests    []string
	testMode bool
}

func (s *story) SetWriter(w io.Writer) {
	s.writer = w
}

func (s *story) Start() chan string {
	cmdChan := make(chan string)
	go func() {
		s.Render()
		for msg := range cmdChan {
			if s.testMode {
				s.println("> " + msg)
				s.println("")
			}
			s.Update(msg)
			s.Render()
		}
	}()
	return cmdChan

}

func (s *story) Test() {
	s.testMode = true
	cmdChan := s.Start()
	for _, cmd := range s.tests {
		cmdChan <- cmd
	}
	close(cmdChan)
	s.testMode = false
}

func (s *story) println(value string) {
	fmt.Fprintln(s.writer, value)
}

func (s *story) Render() {
	if !s.state.titleDisplayed {
		s.println(s.title)
		s.println("")
		s.state.titleDisplayed = true
	}
	room := s.rooms[s.startAt]
	s.println(room.Name())
	s.println(room.Description())
	s.println("")
}

func (s *story) Update(cmd string) {

}
