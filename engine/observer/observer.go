package observer

type Observer interface {
	OnNotify(Event)
}
