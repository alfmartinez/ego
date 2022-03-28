package prompt

import (
	"bufio"
	"github.com/alfmartinez/ego/engine/input"
	"os"
)

func init() {
	input.RegisterInputHandler("prompt", func() input.InputHandler {
		scanner := bufio.NewScanner(os.Stdin)
		return &promptHandler{scanner}
	})
}

type TextHandler interface {
	GetText() string
}

type promptHandler struct {
	scanner *bufio.Scanner
}

func (h *promptHandler) Handle(input input.Event) {

}

func (h *promptHandler) GetText() string {
	h.scanner.Scan()
	return h.scanner.Text()
}
