package event

type Broker interface {
	Register() (c, notify chan Event)
}

var brokerInstance Broker = NewBroker()

func GetBrokerInstance() Broker {
	return brokerInstance
}

func NewBroker() Broker {
	b := &broker{
		notify:    make(chan Event),
		observers: make([]chan Event, 0),
	}
	b.start()
	return b
}

type broker struct {
	notify    chan Event
	observers []chan Event
}

func (b *broker) start() {
	go func() {
		for e := range b.notify {
			b.notifyAll(e)
		}
	}()
}

// Notify implements Broker
func (b *broker) notifyAll(e Event) {
	for _, o := range b.observers {
		o <- e
	}
}

// Register implements Broker
func (b *broker) Register() (c, notify chan Event) {
	c = make(chan Event)
	b.observers = append(b.observers, c)
	return c, b.notify
}

func Register() (c, notify chan Event) {
	return brokerInstance.Register()
}
