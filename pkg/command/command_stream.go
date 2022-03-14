package command

type CommandStream interface {
}

func CreateCommandStream() CommandStream {
	return &commandStream{}
}

type commandStream struct {
}
