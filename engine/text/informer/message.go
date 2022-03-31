package informer

type Action string

type Message struct {
	Story  Story
	Phase  Phase
	Action Action
}
