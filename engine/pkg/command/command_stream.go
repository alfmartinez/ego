package command

type CommandStream interface {
	After(Command)
	Before(Command)
	Execute() bool
	Abort()
}

func CreateCommandStream() CommandStream {
	return &commandStream{}
}

type commandStream struct {
	stream  []Command
	aborted bool
}

func (s *commandStream) Abort() {
	s.aborted = true
}

func (s *commandStream) After(c Command) {
	s.stream = append(s.stream, c)
}

func (s *commandStream) Before(c Command) {
	s.stream = append([]Command{c}, s.stream...)
}

func (s *commandStream) Execute() bool {
	if s.aborted {
		return true
	}
	if len(s.stream) > 0 {
		done := s.stream[0].Execute()
		if done {
			s.stream = s.stream[1:]
		}
	}
	return len(s.stream) == 0
}
