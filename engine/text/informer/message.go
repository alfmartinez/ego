package informer

type Action string

type (
	Message interface {
		IsMessage()
	}
	Event struct {
		Name string
	}
	Command struct {
		Cmd string
	}
	Try struct {
		Action   Action
		Argument string
	}
)

func (Event) IsMessage()   {}
func (Command) IsMessage() {}
func (Try) IsMessage()     {}
