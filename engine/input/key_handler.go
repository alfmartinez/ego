package input

func init() {
	RegisterInputHandler("key", func() InputHandler {
		return CreateKeyHandler()
	})
}

type KeyHandler interface {
	IsPressed(Key) bool
	AllReleased() bool
}

func CreateKeyHandler() InputHandler {
	return &keyHandler{
		keyStatus: make(map[Key]bool),
	}
}

type keyHandler struct {
	keyStatus map[Key]bool
}

func (h *keyHandler) Handle(e Event) {
	switch e.Action {
	case PRESSED:
		h.keyStatus[e.Key] = true
	case RELEASED:
		delete(h.keyStatus, e.Key)
	}
}

func (h *keyHandler) IsPressed(k Key) bool {
	if _, ok := h.keyStatus[k]; !ok {
		return false
	}
	return true
}

func (h *keyHandler) AllReleased() bool {
	for _, key := range h.keyStatus {
		if key {
			return false
		}
	}
	return true
}
