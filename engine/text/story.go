package text

import (
	"fmt"
	"io"
)

type Story interface {
	Test()
	Start() chan *Command
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
	current  string
	tests    []string
	testMode bool
}

func (s *story) SetWriter(w io.Writer) {
	s.writer = w
}

func (s *story) Start() chan *Command {
	cmdChan := make(chan *Command)
	go func() {
		s.Render()
		for msg := range cmdChan {
			if s.testMode {
				s.println("> " + msg.String())
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
		command := ParseCommand(cmd)
		cmdChan <- command
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
	room := s.rooms[s.current]
	s.println(room.Name())
	s.println(room.Description())
	s.println("")
}

func (s *story) Update(cmd *Command) {
	room := s.rooms[s.current]
	if result := room.Execute(cmd); result != "" {
		s.current = result
	}
}
