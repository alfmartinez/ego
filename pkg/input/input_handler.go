package input

type InputHandler interface {
	Handle(Event)
	IsPressed(Key) bool
	AllReleased() bool
}

func CreateInputHandler() InputHandler {
	return &inputHandler{
		keyStatus: make(map[Key]bool),
	}
}

type inputHandler struct {
	keyStatus map[Key]bool
}

func (h *inputHandler) Handle(e Event) {
	switch e.Action {
	case PRESSED:
		h.keyStatus[e.Key] = true
	case RELEASED:
		delete(h.keyStatus, e.Key)
	}
}

func (h *inputHandler) IsPressed(k Key) bool {
	if _, ok := h.keyStatus[k]; !ok {
		return false
	}
	return true
}

func (h *inputHandler) AllReleased() bool {
	for _, key := range h.keyStatus {
		if key {
			return true
		}
	}
	return true
}
