package informer

type (
	Index[T any] interface {
		Get(string) T
		Add(string, T)
	}

	index[T any] struct {
		items map[string]T
	}
)

func (i index[T]) Get(key string) T {
	return i.items[key]
}

func (i index[T]) Add(key string, o T) {
	i.items[key] = o
}

var objectIndex = index[Object]{make(map[string]Object)}
var aliasIndex = index[Object]{make(map[string]Object)}
