package command

type CommandStream interface {
	After(Command)
	Before(Command)
	Execute() bool
}

func CreateCommandStream() CommandStream {
	return &commandStream{}
}

type commandStream struct {
	stream []Command
}

func (s *commandStream) After(c Command) {
	s.stream = append(s.stream, c)
}

func (s *commandStream) Before(c Command) {
	s.stream = append([]Command{c}, s.stream...)
}

func (s *commandStream) Execute() bool {
	if len(s.stream) > 0 {
		done := s.stream[0].Execute()
		if done {
			s.stream = s.stream[1:]
		}
	}
	return len(s.stream) == 0
}
