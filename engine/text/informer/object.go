package informer

type (
	Object interface {
		Kind
	}

	object struct {
		Kind
	}
)

func CreateObject(kind string) Object {
	return &object{
		kinds[kind],
	}
}
